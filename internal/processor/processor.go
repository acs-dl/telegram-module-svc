package processor

import (
	"context"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/acs-dl/telegram-module-svc/internal/config"
	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/data/manager"
	"github.com/acs-dl/telegram-module-svc/internal/data/postgres"
	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	"github.com/acs-dl/telegram-module-svc/internal/sender"
	"github.com/acs-dl/telegram-module-svc/internal/tg_client"
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
	HandleAddUserAction(msg data.ModulePayload) (string, error)
	HandleUpdateUserAction(msg data.ModulePayload) (string, error)
	HandleRemoveUserAction(msg data.ModulePayload) (string, error)
	HandleDeleteUserAction(msg data.ModulePayload) (string, error)
	HandleVerifyUserAction(msg data.ModulePayload) (string, error)
	SendDeleteUser(uuid string, user data.User) error
}

type processor struct {
	log             *logan.Entry
	telegramClient  tg_client.TelegramClient
	permissionsQ    data.Permissions
	usersQ          data.Users
	chatsQ          data.Chats
	managerQ        *manager.Manager
	sender          *sender.Sender
	pqueues         *pqueue.PQueues
	unverifiedTopic string
	identityTopic   string
}

func NewProcessorAsInterface(cfg config.Config, ctx context.Context) interface{} {
	return interface{}(&processor{
		log:             cfg.Log().WithField("service", ServiceName),
		telegramClient:  tg_client.TelegramClientInstance(ctx),
		permissionsQ:    postgres.NewPermissionsQ(cfg.DB()),
		usersQ:          postgres.NewUsersQ(cfg.DB()),
		chatsQ:          postgres.NewChatsQ(cfg.DB()),
		managerQ:        manager.NewManager(cfg.DB()),
		sender:          sender.SenderInstance(ctx),
		pqueues:         pqueue.PQueuesInstance(ctx),
		unverifiedTopic: cfg.Amqp().Unverified,
		identityTopic:   cfg.Amqp().Identity,
	})
}

func ProcessorInstance(ctx context.Context) Processor {
	return ctx.Value(ServiceName).(Processor)
}

func CtxProcessorInstance(entry interface{}, ctx context.Context) context.Context {
	return context.WithValue(ctx, ServiceName, entry)
}
