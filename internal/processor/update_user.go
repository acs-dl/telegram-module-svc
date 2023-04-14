package processor

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/helpers"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateUpdateUser(msg data.ModulePayload) error {
	phoneValidationCase := validation.When(msg.Username == nil, validation.Required.Error("phone is required if username is not set"))
	usernameValidationCase := validation.When(msg.Phone == nil, validation.Required.Error("username is required if phone is not set"))

	return validation.Errors{
		"link":         validation.Validate(msg.Link, validation.Required),
		"username":     validation.Validate(msg.Username, usernameValidationCase),
		"phone":        validation.Validate(msg.Phone, phoneValidationCase),
		"access_level": validation.Validate(msg.AccessLevel, validation.Required),
	}.Filter()
}

func (p *processor) handleUpdateUserAction(msg data.ModulePayload) error {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateUpdateUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to validate fields")
	}

	user, err := helpers.GetUser(p.pqueues.SuperPQueue, any(p.telegramClient.GetUserFromApi), []any{any(msg.Username), any(msg.Phone)}, pqueue.NormalPriority)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from api for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get user from api")
	}

	if user == nil {
		p.log.Errorf("no user was found for message action with id `%s`", msg.RequestId)
		return errors.New("no user was found")
	}

	chat, err := helpers.GetChat(p.pqueues.SuperPQueue, any(p.telegramClient.GetChatFromApi), []any{any(msg.Link)}, pqueue.NormalPriority)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get chat from api for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get chat from api")
	}

	if chat == nil {
		p.log.Errorf("no chat `%s` was found for message action with id `%s`", msg.Link, msg.RequestId)
		return errors.New("no chat was found")
	}

	user, err = helpers.GetUser(p.pqueues.SuperPQueue, any(p.telegramClient.GetChatUserFromApi), []any{any(*user), any(*chat)}, pqueue.NormalPriority)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from api for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get user from api")
	}

	_, err = p.getUserFromDbByTelegramId(user.TelegramId)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from db for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get user from")
	}

	err = helpers.GetRequestError(p.pqueues.SuperPQueue, any(p.telegramClient.UpdateUserInChatFromApi), []any{any(*user), any(*chat)}, pqueue.NormalPriority)
	if err != nil {
		p.log.WithError(err).Errorf("failed to update user from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to update user from api")
	}

	if err = p.permissionsQ.UpdateAccessLevel(data.Permission{
		TelegramId:  user.TelegramId,
		AccessLevel: msg.AccessLevel,
		Link:        msg.Link,
	}); err != nil {
		p.log.WithError(err).Errorf("failed to update user in db for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to update user in db")
	}

	p.resetFilters()
	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return nil
}
