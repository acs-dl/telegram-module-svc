package tg_client

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gotd/td/telegram"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type TelegramClient interface {
	GetUserFromApi(client *telegram.Client, username, phone *string) (*data.User, error)
	GetChatUsersFromApi(chat Chat) ([]data.User, error)
	GetChatUserFromApi(user data.User, chat Chat) (*data.User, error)
	SearchByFromApi(username, phone *string, amount int) ([]data.User, error)

	GetChatFromApi(title string) (*Chat, error)
	AddUserInChatFromApi(user data.User, chat Chat) error
	UpdateUserInChatFromApi(user data.User, chat Chat) error
	DeleteFromChatFromApi(user data.User, chat Chat) error

	SendMessageFromApi(info data.MessageInfo) error
	GenerateChatLinkFromApi(chat Chat) (string, error)

	GetTg() *tgInfo
	GetSuperClient() *telegram.Client
	GetUsualClient() *telegram.Client
}

type tgInfo struct {
	superUserClient *telegram.Client
	userClient      *telegram.Client
	log             *logan.Entry
	cfg             config.Config
	ctx             context.Context
}

type Chat struct {
	id         int64
	accessHash *int64
}

func NewTgAsInterface(cfg config.Config, ctx context.Context) interface{} {
	log := cfg.Log()
	currentDir, err := os.Getwd()
	if err != nil {
		log.WithError(err).Errorf("failed to get current directory path")
		panic(errors.Wrap(err, "failed to get current directory path"))
	}

	superSessionFile := filepath.Join(currentDir, "super_session.json")
	superClient, err := login(ctx, cfg.Telegram().SuperUser, superSessionFile)
	if err != nil {
		log.WithError(err).Errorf("failed to authenticate super user")
		panic(errors.Wrap(err, "failed to authenticate super user"))
	}

	usualSessionFile := filepath.Join(currentDir, "usual_session.json")
	usualClient, err := login(ctx, cfg.Telegram().User, usualSessionFile)
	if err != nil {
		log.WithError(err).Errorf("failed to authenticate usual user")
		panic(errors.Wrap(err, "failed to authenticate usual user"))
	}

	return interface{}(&tgInfo{
		superUserClient: superClient,
		userClient:      usualClient,
		log:             log,
		cfg:             cfg,
		ctx:             ctx,
	})
}

func TelegramClientInstance(ctx context.Context) TelegramClient {
	return ctx.Value("telegram").(TelegramClient)
}

func CtxTelegramClientInstance(entry interface{}, ctx context.Context) context.Context {
	return context.WithValue(ctx, "telegram", entry)
}
