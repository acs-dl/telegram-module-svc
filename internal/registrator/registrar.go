package registrator

import (
	"context"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
)

func Run(ctx context.Context, cfg config.Config) {
	NewRegistrar(cfg).Run(ctx)
}