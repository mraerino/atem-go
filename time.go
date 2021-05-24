package atem

import (
	"context"
	"time"

	"github.com/mraerino/atem-go/models"
	"github.com/mraerino/atem-go/packet"
	"github.com/mraerino/atem-go/packet/cmds"
	"github.com/pkg/errors"
)

var (
	ErrChannelFull   = errors.New("Channel for time requests is full")
	ErrChannelClosed = errors.New("Channel was closed")
)

// Timecode requests the current system timecode from the switcher
//
// The context can be used to provide a timeout
func (c *Client) Timecode(ctx context.Context) (*models.Timecode, error) {
	msg := c.makeHeader(packet.FlagACK)
	msg.Commands = append(msg.Commands, cmds.TiRqCmd{})
	if err := c.send(msg); err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}

	ch := make(chan models.Timecode)

	select {
	case c.timeRequests <- ch:
	default:
		return nil, ErrChannelFull
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	select {
	case time, open := <-ch:
		if !open {
			return nil, ErrChannelClosed
		}
		return &time, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
