package processor

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/helpers"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateRemoveUser(msg data.ModulePayload) error {
	phoneValidationCase := validation.When(msg.Username == nil, validation.Required.Error("phone is required if username is not set"))
	usernameValidationCase := validation.When(msg.Phone == nil, validation.Required.Error("username is required if phone is not set"))

	return validation.Errors{
		"link":     validation.Validate(msg.Link, validation.Required),
		"username": validation.Validate(msg.Username, usernameValidationCase),
		"phone":    validation.Validate(msg.Phone, phoneValidationCase),
	}.Filter()
}

func (p *processor) handleRemoveUserAction(msg data.ModulePayload) error {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateRemoveUser(msg)
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

	dbUser, err := p.getUserFromDbByTelegramId(user.TelegramId)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from db for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get user from")
	}

	err = helpers.GetRequestError(p.pqueues.SuperPQueue, any(p.telegramClient.DeleteFromChatFromApi), []any{any(*user), any(*chat)}, pqueue.NormalPriority)
	if err != nil {
		p.log.WithError(err).Errorf("failed to remove user from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "some error while removing user from api")
	}

	err = p.managerQ.Transaction(func() error {
		err = p.permissionsQ.FilterByTelegramIds(user.TelegramId).FilterByLinks(msg.Link).Delete()
		if err != nil {
			p.log.WithError(err).Errorf("failed to delete permission by telegram id `%d` for message action with id `%s`", user.TelegramId, msg.RequestId)
			return errors.Wrap(err, "failed to delete permission")
		}

		permissionsAmount, err := p.permissionsQ.Count().FilterByTelegramIds(user.TelegramId).GetTotalCount()
		if err != nil {
			p.log.WithError(err).Errorf("failed to get permissions amount by telegram id `%d` for message action with id `%s`", user.TelegramId, msg.RequestId)
			return errors.Wrap(err, "failed to get permissions amount")
		}

		if permissionsAmount == 0 {
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
		}

		return nil
	})
	if err != nil {
		p.log.WithError(err).Errorf("failed to make remove user transaction for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to make remove user transaction")
	}

	p.resetFilters()
	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return nil
}
