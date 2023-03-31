package pgdb

import (
	"context"
	"sync"
	"time"

	"github.com/lib/pq"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

//EventsRouterOpts - config for new EventsRouter
type EventsRouterOpts struct {
	Log      *logan.Entry
	Channel  string
	Listener *pq.Listener
}

//EventsRouter - listens to events of postgres channel and sends notification to listeners
type EventsRouter interface {
	RunNewListener(ctx context.Context, opts EventsListenerOpts)
}

//NewEventsRouter - creates new EventsRouter and starts corresponding routines
func NewEventsRouter(ctx context.Context, opts EventsRouterOpts) EventsRouter {
	r := &router{
		EventsRouterOpts: opts,
		lock:             sync.Mutex{},
		listeners:        nil,
	}

	go running.WithBackOff(ctx, r.Log, "run", r.run, time.Second*10, time.Second*10, time.Minute)
	return r
}

type router struct {
	EventsRouterOpts
	lock      sync.Mutex
	listeners []*eventListener
}

//RunNewListener - subscribes new listener to postgres events and runs blocking function to process them
func (r *router) RunNewListener(ctx context.Context, opts EventsListenerOpts) {
	notify := make(chan struct{}, 1) // buffer of 1 is needed to ensure we do not miss event
	notify <- struct{}{}             // trigger processing of events that were created before start of the listener
	l := &eventListener{
		EventsListenerOpts: opts,
		notify:             notify,
	}
	r.subscribe(l)
	// creation of new goroutine should be handle by caller
	l.run(ctx)
}

func (r *router) run(ctx context.Context) error {
	err := r.Listener.UnlistenAll()
	if err != nil {
		return errors.Wrap(err, "failed to unlisten")
	}

	err = r.Listener.Listen(r.Channel)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to channel")
	}

	for {
		r.Log.Debug("waiting for signal")
		var notification *pq.Notification
		var ok bool
		select {
		case <-ctx.Done():
			return ctx.Err()

		case notification, ok = <-r.Listener.NotificationChannel():
			if !ok {
				panic(errors.New("notification channel closed"))
			}

			if notification == nil {
				r.Log.Debug("reconnected")
			} else {
				r.Log.Debug("received notification about db event")
			}
		}

		r.Log.Debug("starting to notify subscribers")
		r.notifyAll(notification)
		r.Log.Debug("finished iteration")
	}
}

func (r *router) subscribe(l *eventListener) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.listeners = append(r.listeners, l)
}

func (r *router) notifyAll(notification *pq.Notification) {
	r.lock.Lock()
	defer r.lock.Unlock()
	for _, sub := range r.listeners {
		// notification is nil on reconnect, so need to notify all listeners in case we've missed something
		if notification != nil && sub.Payload != "" && sub.Payload != notification.Extra {
			continue
		}

		select {
		case sub.notify <- struct{}{}:
		default:

		}
	}
}
