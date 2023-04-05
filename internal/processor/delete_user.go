package processor

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
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

	item, err := p.addFunctionInPqueue(any(p.telegramClient.GetUserFromApi), []any{any(msg.Username), any(msg.Phone)}, 10)
	if err != nil {
		p.log.WithError(err).Errorf("failed to add function in pqueue for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to add function in pqueue")
	}

	err = item.Response.Error
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get user from api")
	}
	user, err := p.convertUserFromInterfaceAndCheck(item.Response.Value)
	if err != nil {
		p.log.WithError(err).Errorf("something wrong with user for message action with id `%s`", msg.RequestId)
		return errors.Errorf("something wrong with user from api")
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
		item, err = p.addFunctionInPqueue(any(p.telegramClient.GetChatUserFromApi), []any{any(msg.Username), any(msg.Phone), any(permission.Link)}, 10)
		if err != nil {
			p.log.WithError(err).Errorf("failed to add function in pqueue for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "failed to add function in pqueue")
		}

		err = item.Response.Error
		if err != nil {
			p.log.WithError(err).Errorf("failed to get chat user from API for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "some error while checking user from api")
		}
		chatUser, ok := item.Response.Value.(*data.User)
		if !ok {
			p.log.WithError(err).Errorf("failed to convert interface to user for message action with id `%s`", msg.RequestId)
			return errors.Errorf("wrong response type while getting users from api")
		}

		if chatUser != nil {
			item, err = p.addFunctionInPqueue(any(p.telegramClient.DeleteFromChatFromApi), []any{any(msg.Username), any(msg.Phone), any(permission.Link)}, 10)
			if err != nil {
				p.log.WithError(err).Errorf("failed to add function in pqueue for message action with id `%s`", msg.RequestId)
				return errors.Wrap(err, "failed to add function in pqueue")
			}

			err = item.Response.Error
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
