package service

import (
	"context"
	"sync"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/processor"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/receiver"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/registrator"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/sender"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/handlers"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/tg_client"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/worker"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
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
	{"receiver", receiver.NewReceiverAsInterface, receiver.RunReceiverAsInterface, receiver.CtxReceiverInstance},
	{"worker", worker.NewWorkerAsInterface, worker.RunWorkerAsInterface, worker.CtxWorkerInstance},
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
	go pqueues.SuperPQueue.ProcessQueue(cfg.RateLimit().RequestsAmount, cfg.RateLimit().TimeLimit, stopProcessQueue)
	go pqueues.UsualPQueue.ProcessQueue(cfg.RateLimit().RequestsAmount, cfg.RateLimit().TimeLimit, stopProcessQueue)
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
