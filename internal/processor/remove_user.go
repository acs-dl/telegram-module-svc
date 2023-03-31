package processor

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
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

	user, err := p.telegramClient.GetChatUserFromApi(msg.Username, msg.Phone, msg.Link)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "some error while getting user from api")
	}
	if user == nil {
		p.log.Errorf("user is not in chat for message action with id `%s`", msg.RequestId)
		return errors.Errorf("user is not in chat")
	}

	dbUser, err := p.getUserFromDbByTelegramId(user.TelegramId)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from db for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get user from")
	}

	err = p.telegramClient.DeleteFromChatFromApi(msg.Username, msg.Phone, msg.Link)
	if err != nil {
		p.log.WithError(err).Errorf("failed to remove user from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to remove user from api")
	}

	err = p.managerQ.Transaction(func() error {
		err = p.permissionsQ.Delete(user.TelegramId, msg.Link)
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
			err = p.usersQ.Delete(user.TelegramId)
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

func (p *processor) sendDeleteInUnverifiedOrUpdateInIdentity(requestId string, user data.User) error {
	if user.Id == nil {
		err := p.SendDeleteUser(requestId, user)
		if err != nil {
			return errors.Wrap(err, "failed to publish delete telegram user in telegram-module")
		}
	} else {
		err := p.sendUpdateUserTelegram(requestId, data.ModulePayload{
			RequestId: requestId,
			UserId:    fmt.Sprintf("%d", *user.Id),
			Action:    RemoveTelegramAction,
		})
		if err != nil {
			return errors.Wrap(err, "failed to publish update telegram user in identity-svc")
		}
	}

	return nil
}
