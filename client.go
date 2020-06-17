package atem

import (
	"bytes"
	"context"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mraerino/atem-go/models"
	"github.com/mraerino/atem-go/packet"
	"github.com/mraerino/atem-go/packet/cmds"
	"github.com/netlify/netlify-commons/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Client struct {
	host string
	port int

	conn *net.UDPConn
	log  logrus.FieldLogger

	started   util.AtomicBool
	connected util.AtomicBool

	sessionID uint32 // atomic value (actually uint16)
	seqNum    uint32 // atomic counter (actually uint16)

	state      SwitcherState
	stateMutex sync.RWMutex
}

func NewClient(log logrus.FieldLogger, host string) *Client {
	return NewClientWithPort(log, host, 9910)
}

func NewClientWithPort(log logrus.FieldLogger, host string, port int) *Client {
	return &Client{
		host:      host,
		port:      port,
		log:       log,
		started:   util.NewAtomicBool(false),
		connected: util.NewAtomicBool(false),
	}
}

func (c *Client) State() SwitcherState {
	c.stateMutex.RLock()
	state := c.state
	c.stateMutex.RUnlock()
	return state
}

var ErrAlreadyStarted = errors.New("Already started")

func (c *Client) Start(ctx context.Context) error {
	if c.started.Set(true) {
		return ErrAlreadyStarted
	}

	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(c.host, strconv.Itoa(c.port)))
	if err != nil {
		return errors.Wrap(err, "Failed to resolve UDP address")
	}

	c.conn, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		return errors.Wrap(err, "Failed to open UDP socket")
	}

	c.log.Debug("Created socket. Starting message processing task...")
	connected := c.startHandleMessages(ctx)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	select {
	case err, open := <-connected:
		if !open {
			return nil
		}
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

var InitPayload = []byte{1, 0, 0, 0, 0, 0, 0, 0}

func (c *Client) sendInit() error {
	initMSG := c.makeHeader(packet.FlagInit)
	initMSG.Length = 20 // custom well-known length
	buf := new(bytes.Buffer)
	initMSG.Serialize(buf)
	buf.Write(InitPayload)

	_, err := c.conn.Write(buf.Bytes())
	return err
}

func (c *Client) send(msg *packet.Message) error {
	buf := new(bytes.Buffer)
	msg.Serialize(buf)

	_, err := c.conn.Write(buf.Bytes())
	return err
}

func (c *Client) resetSession() {
	sessionID := rand.Uint32()
	atomic.StoreUint32(&c.sessionID, sessionID)
	atomic.StoreUint32(&c.seqNum, 0)
	c.connected.Set(false)

	// init state
	c.state = NewSwitcherState()
}

func (c *Client) startHandleMessages(ctx context.Context) chan error {
	errCh := make(chan error)

	go func() {
		var innerCtx context.Context
		var cancel context.CancelFunc
		for ctx.Err() == nil {
			// cancel previous session
			if cancel != nil {
				cancel()
			}
			innerCtx, cancel = context.WithCancel(ctx)

			c.resetSession()
			var connected bool

			c.log.Debug("Sending init packet")
			err := c.sendInit()
			if err != nil {
				errCh <- errors.Wrap(err, "Failed to send init message")
				break
			}

			// ack packets in background
			ackQueue := make(chan uint16, 20) // seq nums to ack
			go c.ackWorker(innerCtx, ackQueue)

			// read loop
			var rerr error
			for {
				raw := make([]byte, 1500) // safe size for expected MTU
				// todo: read deadline
				n, err := c.conn.Read(raw)
				if err != nil {
					rerr = err
					break
				}
				c.log.WithField("len", n).Debug("Got packet")

				buf := bytes.NewBuffer(raw[:n])
				msg, err := packet.Deserialize(c.log, buf)
				if err != nil {
					rerr = err
					break
				}
				log := c.log.WithField("seq", msg.SeqNum)
				log.WithFields(msg.Flags.Debug()).Debug("Read packet")

				oldSession := atomic.SwapUint32(&c.sessionID, uint32(msg.SessionID))
				if oldSession != uint32(msg.SessionID) {
					log.WithField("session", msg.SessionID).Info("Using new session id from switcher")
				}

				if msg.Flags.Has(packet.FlagRetrans) {
					// todo: handle double sends
					log.Warn("Retransmission detected")
				}

				if msg.Flags.Has(packet.FlagInit) {
					// always ack init packets
					msg.Flags |= packet.FlagNeedACK
				}

				if msg.Flags.Has(packet.FlagNeedACK) {
					log.Debug("queueing packet for ack")
					ackQueue <- msg.SeqNum
				}

				gotInit, err := c.processCommands(log, msg)
				if err != nil {
					log.WithError(err).Warn("Error while processing commands")
				}

				if gotInit && !connected {
					connected = true
					c.connected.Set(true)
					errCh <- nil // avoid risk of double-closing
				}
			}
			if rerr != nil {
				c.log.WithError(rerr).Warn("Failed to read")
				if !connected {
					errCh <- errors.Wrap(rerr, "Failed to read message")
				}
				break
			}
		}
		if cancel != nil {
			cancel()
		}
		_ = c.conn.Close()
		c.started.Set(false)
	}()

	return errCh
}

func (c *Client) processCommands(log logrus.FieldLogger, msg packet.Message) (init bool, err error) {
	c.stateMutex.Lock()
	defer c.stateMutex.Unlock()
	for _, cmd := range msg.Commands {
		switch t := cmd.(type) {
		case *cmds.UnknownCommand:
			log.WithField("slug", t.Slug()).Debug("Got unknown command")
		case *cmds.IncmCmd:
			init = true
		case *cmds.VerCmd:
			c.state.Version.Major = int(t.Major)
			c.state.Version.Minor = int(t.Minor)
			log.WithFields(logrus.Fields{"major": t.Major, "minor": t.Minor}).Debug("Got version")
		case *cmds.PinCmd:
			c.state.Config.ProductName = t.Value()
		case *cmds.WarnCmd:
			c.state.Warning = t.Value()
		case *cmds.TopCmd:
			c.state.Config.Topology = models.Topology(*t)
		case *cmds.MecCmd:
			c.state.Config.MixEffect[int(t.ME)] = int(t.Keyers)
		case *cmds.MplCmd:
			c.state.Config.MediaPlayer = models.MediaPlayerConfig(*t)
		case *cmds.MvcCmd:
			c.state.Config.MultiViews = int(*t)
		case *cmds.SscCmd:
			c.state.Config.SuperSources = int(*t)
		case *cmds.TlcCmd:
			c.state.Config.TallyChannels = int(*t)
		case *cmds.MacCmd:
			c.state.Config.MacroBanks = int(*t)
		case *cmds.PowrCmd:
			c.state.Power = models.PowerStatus(*t)
		case *cmds.VidmCmd:
			c.state.VideoMode = models.VideoMode(*t)
		case *cmds.InprCmd:
			c.state.Inputs[t.SourceIndex] = models.InputProperties(*t)
		case *cmds.PrgiCmd:
			c.state.Program[int(t.Bus)] = t.Source
		case *cmds.PrviCmd:
			c.state.Preview[int(t.Bus)] = t.Source
		case *cmds.AuxsCmd:
			c.state.Aux[int(t.Bus)] = t.Source
		case *cmds.MpceCmd:
			mp, ok := c.state.MediaPlayer[int(t.PlayerIndex)]
			if !ok {
				mp = &models.MediaPlayer{}
				c.state.MediaPlayer[int(t.PlayerIndex)] = mp
			}
			mp.Type = t.Type
			mp.StillIndex = t.StillIndex
			mp.ClipIndex = t.ClipIndex
		case *cmds.MpfeCmd:
			c.state.MediaFiles[int(t.Index)] = models.MediaStillFrame{
				Used:     t.Used,
				Hash:     t.Hash,
				Filename: t.Filename,
			}
		default:
			log.Debug("Unhandled packet type")
		}
	}
	return
}

const ackDebounceInterval = time.Millisecond * 20
const ackMaxDebounce = time.Millisecond * 100

// since the protocol seems to allow just acking the seq num
// of the packet that was received last, the actual sending
// of the ack is debounced to improve performance
func (c *Client) ackWorker(ctx context.Context, queue <-chan uint16) {
	timer := time.NewTimer(ackDebounceInterval)
	timer.Stop() // start with a stopped timer

	lastSend := time.Now()
	var latestNum uint16
	for {
		log := c.log.WithField("num", latestNum)
		select {
		case <-ctx.Done():
			timer.Stop()
			log.WithError(ctx.Err()).Debug("ACK worker stopped")
			return
		case seqNum, open := <-queue:
			if !open {
				log.Debug("ACK work channel closed")
				return
			}
			if seqNum > latestNum {
				latestNum = seqNum
			}
			timer.Stop()
			if time.Since(lastSend) > ackMaxDebounce {
				timer.Reset(0)
			} else {
				timer.Reset(ackDebounceInterval)
			}
		case <-timer.C:
			log.Debug("Ack timer fired. Sending ack...")
			// defer timer fired, actually send ack
			msg := c.makeHeader(packet.FlagACK)
			msg.AckID = latestNum
			err := c.send(msg)
			if err != nil {
				log.WithError(err).Warn("Failed sending ack")
			}
			lastSend = time.Now()
		}
	}
}

func (c *Client) makeHeader(flags ...packet.Flags) *packet.Message {
	seqNum := atomic.AddUint32(&c.seqNum, 1) - 1 // allow seq num to start at 0
	return &packet.Message{
		Flags:     packet.FlagsFrom(flags...),
		SessionID: uint16(atomic.LoadUint32(&c.sessionID)),
		SeqNum:    uint16(seqNum),
	}
}
