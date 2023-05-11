package registrator

import (
	"context"

	"github.com/acs-dl/telegram-module-svc/internal/config"
)

func Run(ctx context.Context, cfg config.Config) {
	NewRegistrar(cfg).Run(ctx)
}

func RunRegistrarAsInterface(structure interface{}, ctx context.Context) {
	(structure.(Registrar)).Run(ctx)
}
