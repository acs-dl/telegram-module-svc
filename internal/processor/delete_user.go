package processor

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateDeleteUser(msg data.ModulePayload) error {
	return validation.Errors{
		"username": validation.Validate(msg.Username, validation.Required),
		"phone":    validation.Validate(msg.Phone, validation.Required),
	}.Filter()
}

func (p *processor) handleDeleteUserAction(msg data.ModulePayload) error {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateDeleteUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to validate fields")
	}

	user, err := p.telegramClient.GetUserFromApi(msg.Username, msg.Phone)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get user from api")
	}

	dbUser, err := p.usersQ.FilterByTelegramIds(user.TelegramId).Get()
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get user")
	}
	if dbUser == nil {
		p.log.Errorf("no such user in module for message action with id `%s`", msg.RequestId)
		return errors.Errorf("no such user in module")
	}

	permissions, err := p.permissionsQ.FilterByTelegramIds(user.TelegramId).Select()
	if err != nil {
		p.log.WithError(err).Errorf("failed to select permissions by telegram id `%d` for message action with id `%s`", user.TelegramId, msg.RequestId)
		return errors.Wrap(err, "failed to select permissions")
	}

	for _, permission := range permissions {
		chatUser, err := p.telegramClient.GetChatUserFromApi(msg.Username, msg.Phone, permission.Link)
		if err != nil {
			p.log.WithError(err).Errorf("failed to get chat user from API for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "some error while checking user from api")
		}

		if chatUser != nil {
			err = p.telegramClient.DeleteFromChatFromApi(msg.Username, msg.Phone, permission.Link)
			if err != nil {
				p.log.WithError(err).Errorf("failed to remove user from API for message action with id `%s`", msg.RequestId)
				return errors.Wrap(err, "some error while removing user from api")
			}
		}

		if err = p.permissionsQ.Delete(permission.TelegramId, permission.Link); err != nil {
			p.log.WithError(err).Errorf("failed to delete permission from db for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "failed to delete permission")
		}
	}

	err = p.usersQ.Delete(user.TelegramId)
	if err != nil {
		p.log.WithError(err).Errorf("failed to delete user by telegram id `%d` for message action with id `%s`", user.TelegramId, msg.RequestId)
		return errors.Wrap(err, "failed to delete user")
	}

	if dbUser.Id == nil {
		err = p.sendDeleteUser(msg.RequestId, *dbUser)
		if err != nil {
			p.log.WithError(err).Errorf("failed to publish delete user for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "failed to publish delete user")
		}
	}

	p.resetFilters()
	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return nil
}
