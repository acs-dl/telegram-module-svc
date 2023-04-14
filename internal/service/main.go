package service

import (
	"context"
	"sync"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/receiver"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/registrator"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/sender"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/handlers"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/tg"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/worker"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/types"
)

var availableServices = map[string]types.Runner{
	"api":       api.Run,
	"sender":    sender.Run,
	"receiver":  receiver.Run,
	"worker":    worker.Run,
	"registrar": registrator.Run,
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
	ctx = handlers.CtxPQueues(&pqueues, ctx)

	// create new tg sessions better from this point
	tgClient := tg.NewTg(cfg.Telegram(), cfg.Log())
	ctx = handlers.CtxTGClient(&tgClient, ctx)

	for serviceName, service := range availableServices {
		wg.Add(1)

		go func(name string, runner types.Runner) {
			defer wg.Done()

			runner(ctx, cfg)

		}(serviceName, service)

		logger.WithField("service", serviceName).Info("Service started")
	}

	wg.Wait()

}
