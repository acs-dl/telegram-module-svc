package pgdb

import (
	"context"
	"time"

	"gitlab.com/distributed_lab/running"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

//Handler - handles notification of new event received form postgres or ticks
type Handler interface {
	Handle(ctx context.Context) error
}

//EventsListenerOpts - defines config for events listener
type EventsListenerOpts struct {
	Log        *logan.Entry
	Payload    string // skip notifications from the channel if payload does not match. Empty string is treated as no filter is present
	TickPeriod time.Duration
	Handler    Handler
}

type eventListener struct {
	EventsListenerOpts
	notify chan struct{}
}

func (l *eventListener) run(ctx context.Context) {
	l.Log.Info("starting")
	if l.TickPeriod == 0 {
		l.TickPeriod = 30 * time.Second
	}
	ticker := time.NewTicker(l.TickPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			l.Log.Debug("ticked")
		case _, ok := <-l.notify:
			if !ok {
				panic(errors.New("expected notify channel to always be open"))
			}
		}

		l.resetTicker(ticker)

		running.UntilSuccess(ctx, l.Log, "run_once", func(ctx context.Context) (bool, error) {
			err := l.Handler.Handle(ctx)
			return err == nil, err
		}, 10*time.Second, time.Second*30)
	}
}

func (l *eventListener) resetTicker(ticker *time.Ticker) {
	ticker.Reset(l.TickPeriod)
	select {
	case <-ticker.C:
	default:
	}
}
