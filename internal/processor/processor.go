package processor

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data/manager"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data/postgres"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/tg"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	serviceName = data.ModuleName + "-processor"

	//add needed actions for module
	GetUsersAction   = "get_users"
	AddUserAction    = "add_user"
	UpdateUserAction = "update_user"
	RemoveUserAction = "remove_user"
	VerifyUserAction = "verify_user"
	DeleteUserAction = "delete_user"
)

type Processor interface {
	HandleNewMessage(msg data.ModulePayload) error
}

type processor struct {
	log            *logan.Entry
	telegramClient tg.TelegramClient
	permissionsQ   data.Permissions
	usersQ         data.Users
	managerQ       *manager.Manager
}

var handleActions = map[string]func(proc *processor, msg data.ModulePayload) error{
	GetUsersAction:   (*processor).handleGetUsersAction,
	AddUserAction:    (*processor).handleAddUserAction,
	UpdateUserAction: (*processor).handleUpdateUserAction,
	RemoveUserAction: (*processor).handleRemoveUserAction,
	VerifyUserAction: (*processor).handleVerifyUserAction,
	DeleteUserAction: (*processor).handleDeleteUserAction,
}

func NewProcessor(cfg config.Config) Processor {
	return &processor{
		log:            cfg.Log().WithField("service", serviceName),
		telegramClient: tg.NewTg(cfg.Telegram(), cfg.Log()),
		permissionsQ:   postgres.NewPermissionsQ(cfg.DB()),
		usersQ:         postgres.NewUsersQ(cfg.DB()),
		managerQ:       manager.NewManager(cfg.DB()),
	}
}

func (p *processor) HandleNewMessage(msg data.ModulePayload) error {
	p.log.Infof("handling message with id `%s`", msg.RequestId)

	err := validation.Errors{
		"action": validation.Validate(msg.Action, validation.Required),
	}.Filter()
	if err != nil {
		p.log.WithError(err).Errorf("no such action to handle for message with id `%s`", msg.RequestId)
		return errors.Wrap(err, fmt.Sprintf("no such action `%s` to handle for message with id `%s`", msg.Action, msg.RequestId))
	}

	requestHandler := handleActions[msg.Action]
	if err = requestHandler(p, msg); err != nil {
		p.log.WithError(err).Errorf("failed to handle message with id `%s`", msg.RequestId)
		return err
	}

	p.log.Infof("finish handling message with id `%s`", msg.RequestId)
	return nil
}
