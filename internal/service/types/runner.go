package types

import (
	"context"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
)

type Runner = func(context context.Context, config config.Config)
