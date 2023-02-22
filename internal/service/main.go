package service

import (
	"context"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/receiver"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/sender"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/registrator"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/worker"
	"sync"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/types"
)

var availableServices = map[string]types.Runner{
	"api":      api.Run,
	"sender":   sender.Run,
	"receiver": receiver.Run,
	"worker":   worker.Run,
}

func Run(cfg config.Config) {
	// module registration before starting all services
	regCfg := cfg.Registrator()
	if err := registrator.RegisterModule(data.ModuleName, regCfg); err != nil {
		panic(err)
	}

	logger := cfg.Log().WithField("service", "main")
	ctx := context.Background()
	wg := new(sync.WaitGroup)

	logger.Info("Starting all available services...")

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
