package registrator

import (
	"context"
	"time"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
)

const serviceName = data.ModuleName + "-registrar"

type Registrar interface {
	Run(ctx context.Context)
	UnregisterModule() error
}

type registrar struct {
	logger      *logan.Entry
	config      config.RegistratorConfig
	runnerDelay time.Duration
}

func NewRegistrar(cfg config.Config) Registrar {
	return &registrar{
		logger:      cfg.Log().WithField("runner", serviceName),
		config:      cfg.Registrator(),
		runnerDelay: cfg.Runners().Registrar,
	}
}

func (r *registrar) Run(ctx context.Context) {
	running.WithBackOff(
		ctx,
		r.logger,
		serviceName,
		r.registerModule,
		r.runnerDelay,
		r.runnerDelay,
		r.runnerDelay,
	)
}
