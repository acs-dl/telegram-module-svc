package registrator

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"time"

	"github.com/acs-dl/telegram-module-svc/internal/config"
	"github.com/acs-dl/telegram-module-svc/internal/data"
	"gitlab.com/distributed_lab/running"
)

const ServiceName = data.ModuleName + "-registrar"

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
		logger:      cfg.Log().WithField("runner", ServiceName),
		config:      cfg.Registrator(),
		runnerDelay: cfg.Runners().Registrar,
	}
}

func NewRegistrarAsInterface(cfg config.Config, _ context.Context) interface{} {
	return interface{}(&registrar{
		logger:      cfg.Log().WithField("runner", ServiceName),
		config:      cfg.Registrator(),
		runnerDelay: cfg.Runners().Registrar,
	})
}

func (r *registrar) Run(ctx context.Context) {
	running.WithBackOff(
		ctx,
		r.logger,
		ServiceName,
		r.registerModule,
		r.runnerDelay,
		r.runnerDelay,
		r.runnerDelay,
	)
}
