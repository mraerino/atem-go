package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/mraerino/atem-go"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Level = logrus.InfoLevel
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	client := atem.NewClient(log, "172.22.26.50")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	err := client.Start(ctx)
	if err != nil {
		log.WithError(err).Fatal("Failed to start")
	}
	log.Info("Connected")

	ticker := time.Tick(time.Second / 2)
	for {
		select {
		case <-ctx.Done():
			log.WithError(ctx.Err()).Fatal("exiting")
		case <-ticker:
		}

		ctx, cancel := context.WithTimeout(ctx, time.Second*2)
		tc, err := client.Timecode(ctx)
		if err != nil {
			log.WithError(err).Warn("failed to get time")
		} else {
			fmt.Printf("Time: %#v\n", tc)
		}
		cancel()
		// stateSer, _ := json.MarshalIndent(client.State(), "", "  ")
		// fmt.Printf("State:\n%s\n", string(stateSer))
	}
}
