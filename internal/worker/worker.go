package worker

import (
	"context"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
)

func Run(ctx context.Context, cfg config.Config) {
	NewWorker(cfg).Run(ctx)
}
