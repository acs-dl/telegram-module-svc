package processor

import (
	"context"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data/manager"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data/postgres"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/sender"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/tg_client"
	"gitlab.com/distributed_lab/logan/v3"
)

const (
	ServiceName = data.ModuleName + "-processor"

	//add needed actions for module
	SetUsersAction    = "set_users"
	DeleteUsersAction = "delete_users"

	UpdateTelegramAction = "update_telegram"
	RemoveTelegramAction = "update_telegram"
)

type Processor interface {
	HandleGetUsersAction(msg data.ModulePayload) error
	HandleAddUserAction(msg data.ModulePayload) error
	HandleUpdateUserAction(msg data.ModulePayload) error
	HandleRemoveUserAction(msg data.ModulePayload) error
	HandleDeleteUserAction(msg data.ModulePayload) error
	HandleVerifyUserAction(msg data.ModulePayload) error
	SendDeleteUser(uuid string, user data.User) error
}

type processor struct {
	log            *logan.Entry
	telegramClient tg_client.TelegramClient
	permissionsQ   data.Permissions
	usersQ         data.Users
	managerQ       *manager.Manager
	sender         *sender.Sender
	pqueues        *pqueue.PQueues
}

func NewProcessorAsInterface(cfg config.Config, ctx context.Context) interface{} {
	return interface{}(&processor{
		log:            cfg.Log().WithField("service", ServiceName),
		telegramClient: tg_client.TelegramClientInstance(ctx),
		permissionsQ:   postgres.NewPermissionsQ(cfg.DB()),
		usersQ:         postgres.NewUsersQ(cfg.DB()),
		managerQ:       manager.NewManager(cfg.DB()),
		sender:         sender.SenderInstance(ctx),
		pqueues:        pqueue.PQueuesInstance(ctx),
	})
}

func ProcessorInstance(ctx context.Context) Processor {
	return ctx.Value(ServiceName).(Processor)
}

func CtxProcessorInstance(entry interface{}, ctx context.Context) context.Context {
	return context.WithValue(ctx, ServiceName, entry)
}
