package processor

import (
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

func (p *processor) HandleRemoveUserAction(msg data.ModulePayload) (string, error) {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateRemoveUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to validate fields")
	}

	user, err := p.checkUserExistence(msg.Username, msg.Phone)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to get user")
	}

	err = p.deleteRemotePermission(msg.Link, *user)
	if err != nil {
		p.log.WithError(err).Errorf("failed to remove permission from API for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "some error while removing permission from api")
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

			err = p.sendDeleteInUnverifiedOrUpdateInIdentity(msg.RequestId, *user)
			if err != nil {
				p.log.WithError(err).Errorf("failed to send delete unverified or update identity for message action with id `%s`", msg.RequestId)
				return errors.Wrap(err, "failed to send delete unverified or update identity")
			}
		}

		return nil
	})
	if err != nil {
		p.log.WithError(err).Errorf("failed to make remove user transaction for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to make remove user transaction")
	}

	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return data.SUCCESS, nil
}
