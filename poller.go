package golangtbotapi

import (
	"time"

	"github.com/udev-21/golang-tbot-api/method"
	"github.com/udev-21/golang-tbot-api/types"
)

type Poller interface {
	Poll(b *BotAPI, dest chan types.Update, stop chan struct{})
}

type internalLongPoller struct {
	updates method.GetUpdates
}

func NewLongPoller() Poller {
	return &internalLongPoller{}
}

// Poll does long polling.
func (p *internalLongPoller) Poll(b *BotAPI, dest chan types.Update, stop chan struct{}) {
	for {
		select {
		case <-stop:
			return
		default:
		}
		updates, err := b.GetUpdates(&p.updates)

		if err != nil {
			// b.debug(err)
			continue
		}

		for _, update := range updates {
			p.updates.Offset = update.UpdateID
			dest <- update
		}

		time.Sleep(10 * time.Millisecond)

		if len(updates) > 0 {
			p.updates.Offset++
		} else {
			time.Sleep(10 * time.Millisecond)
		}

	}
}
