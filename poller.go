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
	updateCnf method.GetUpdates
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

		updates, err := b.GetUpdates(&p.updateCnf)
		if len(updates) == 0 {
			time.Sleep(10 * time.Millisecond)
		}

		if err != nil {
			writeLog(LogLevelError, b.logger, "can't get updates")
			continue
		}

		if len(updates) > 0 {
			for _, update := range updates {
				p.updateCnf.Offset = update.UpdateID
				dest <- update
			}
			p.updateCnf.Offset++
		}
	}
}
