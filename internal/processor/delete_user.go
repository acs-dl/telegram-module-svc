package processor

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/helpers"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateDeleteUser(msg data.ModulePayload) error {
	phoneValidationCase := validation.When(msg.Username == nil, validation.Required.Error("phone is required if username is not set"))
	usernameValidationCase := validation.When(msg.Phone == nil, validation.Required.Error("username is required if phone is not set"))

	return validation.Errors{
		"username": validation.Validate(msg.Username, usernameValidationCase),
		"phone":    validation.Validate(msg.Phone, phoneValidationCase),
	}.Filter()
}

func (p *processor) handleDeleteUserAction(msg data.ModulePayload) error {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateDeleteUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to validate fields")
	}

	user, err := helpers.GetUser(p.pqueues.UsualPQueue, any(p.telegramClient.GetUserFromApi), []any{any(msg.Username), any(msg.Phone)}, pqueue.NormalPriority)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from api for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get user from api")
	}

	if user == nil {
		p.log.Errorf("no user was found for message action with id `%s`", msg.RequestId)
		return errors.New("no user was found")
	}

	dbUser, err := p.getUserFromDbByTelegramId(user.TelegramId)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from db for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get user from")
	}

	permissions, err := p.permissionsQ.FilterByTelegramIds(user.TelegramId).Select()
	if err != nil {
		p.log.WithError(err).Errorf("failed to select permissions by telegram id `%d` for message action with id `%s`", user.TelegramId, msg.RequestId)
		return errors.Wrap(err, "failed to select permissions")
	}

	for _, permission := range permissions {
		chat, err := helpers.GetChat(p.pqueues.SuperPQueue, any(p.telegramClient.GetChatFromApi), []any{any(permission.Link)}, pqueue.NormalPriority)
		if err != nil {
			p.log.WithError(err).Errorf("failed to get chat from api for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "failed to get chat from api")
		}

		if chat == nil {
			p.log.Errorf("no chat `%s` was found for message action with id `%s`", permission.Link, msg.RequestId)
			return errors.New("no chat was found")
		}

		chatUser, err := helpers.GetUser(p.pqueues.SuperPQueue, any(p.telegramClient.GetChatUserFromApi), []any{any(*user), any(*chat)}, pqueue.NormalPriority)
		if err != nil {
			p.log.WithError(err).Errorf("failed to get user from api for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "failed to get user from api")
		}

		if chatUser != nil {
			err = helpers.GetRequestError(p.pqueues.SuperPQueue, any(p.telegramClient.DeleteFromChatFromApi), []any{any(*user), any(*chat)}, pqueue.NormalPriority)
			if err != nil {
				p.log.WithError(err).Errorf("failed to remove user from API for message action with id `%s`", msg.RequestId)
				return errors.Wrap(err, "some error while removing user from api")
			}
		}

		if err = p.permissionsQ.FilterByTelegramIds(permission.TelegramId).FilterByLinks(permission.Link).Delete(); err != nil {
			p.log.WithError(err).Errorf("failed to delete permission from db for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "failed to delete permission")
		}
	}

	err = p.usersQ.FilterByTelegramIds(user.TelegramId).Delete()
	if err != nil {
		p.log.WithError(err).Errorf("failed to delete user by telegram id `%d` for message action with id `%s`", user.TelegramId, msg.RequestId)
		return errors.Wrap(err, "failed to delete user")
	}

	err = p.sendDeleteInUnverifiedOrUpdateInIdentity(msg.RequestId, *dbUser)
	if err != nil {
		p.log.WithError(err).Errorf("failed to send delete unverified or update identity for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to send delete unverified or update identity")
	}

	p.resetFilters()
	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return nil
}
