package service

import (
	"context"
	"sync"

	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	"github.com/acs-dl/telegram-module-svc/internal/processor"
	"github.com/acs-dl/telegram-module-svc/internal/receiver"
	"github.com/acs-dl/telegram-module-svc/internal/registrator"
	"github.com/acs-dl/telegram-module-svc/internal/sender"
	"github.com/acs-dl/telegram-module-svc/internal/service/api"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/handlers"
	"github.com/acs-dl/telegram-module-svc/internal/tg_client"
	"github.com/acs-dl/telegram-module-svc/internal/worker"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/acs-dl/telegram-module-svc/internal/config"
)

type svc struct {
	Name    string
	New     func(config.Config, context.Context) interface{}
	Run     func(interface{}, context.Context)
	Context func(interface{}, context.Context) context.Context
}

var services = []svc{
	{"telegram", tg_client.NewTgAsInterface, nil, tg_client.CtxTelegramClientInstance},
	{"sender", sender.NewSenderAsInterface, sender.RunSenderAsInterface, sender.CtxSenderInstance},
	{"processor", processor.NewProcessorAsInterface, nil, processor.CtxProcessorInstance},
	{"worker", worker.NewWorkerAsInterface, worker.RunWorkerAsInterface, worker.CtxWorkerInstance},
	{"receiver", receiver.NewReceiverAsInterface, receiver.RunReceiverAsInterface, receiver.CtxReceiverInstance},
	{"registrar", registrator.NewRegistrarAsInterface, registrator.RunRegistrarAsInterface, nil},
	{"api", api.NewRouterAsInterface, api.RunRouterAsInterface, nil},
}

func Run(cfg config.Config) {
	logger := cfg.Log().WithField("service", "main")
	ctx := context.Background()
	wg := new(sync.WaitGroup)

	logger.Info("Starting all available services...")

	stopProcessQueue := make(chan struct{})
	pqueues := pqueue.NewPQueues()
	go pqueues.SuperUserPQueue.ProcessQueue(cfg.RateLimit().RequestsAmount, cfg.RateLimit().TimeLimit, stopProcessQueue)
	go pqueues.UserPQueue.ProcessQueue(cfg.RateLimit().RequestsAmount, cfg.RateLimit().TimeLimit, stopProcessQueue)
	ctx = pqueue.CtxPQueues(&pqueues, ctx)
	ctx = handlers.CtxConfig(cfg, ctx)

	for _, mySvc := range services {
		wg.Add(1)

		instance := mySvc.New(cfg, ctx)
		if instance == nil {
			logger.WithField("service", mySvc.Name).Warn("Service instance not created")
			panic(errors.Errorf("`%s` instance not created", mySvc.Name))
		}

		if mySvc.Context != nil {
			ctx = mySvc.Context(instance, ctx)
		}

		if mySvc.Run != nil {
			wg.Add(1)
			go func(structure interface{}, runner func(interface{}, context.Context)) {
				defer wg.Done()

				runner(structure, ctx)

			}(instance, mySvc.Run)
		}
		logger.WithField("service", mySvc.Name).Info("Service started")
	}

	wg.Wait()
}
