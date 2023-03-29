package processor

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateGetUsers(msg data.ModulePayload) error {
	return validation.Errors{
		"link": validation.Validate(msg.Link, validation.Required),
	}.Filter()
}

func (p *processor) handleGetUsersAction(msg data.ModulePayload) error {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateGetUsers(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to validate fields")
	}

	users, err := p.telegramClient.GetUsersFromApi(msg.Link)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get users from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "some error while getting users from api")
	}

	if len(users) == 0 {
		p.log.Warnf("no user was found for message action with id `%s`", msg.RequestId)
		return nil
	}

	usersToUnverified := make([]data.User, 0)

	for _, user := range users {
		user.CreatedAt = time.Now()
		err = p.managerQ.Transaction(func() error {
			if err = p.usersQ.Upsert(user); err != nil {
				p.log.WithError(err).Errorf("failed to create user in db for message action with id `%s`", msg.RequestId)
				return errors.Wrap(err, "failed to create user in user db")
			}

			dbUser, err := p.getUserFromDbByTelegramId(user.TelegramId)
			if err != nil {
				p.log.WithError(err).Errorf("failed to get user from db for message action with id `%s`", msg.RequestId)
				return errors.Wrap(err, "failed to get user from")
			}

			user.Id = dbUser.Id
			usersToUnverified = append(usersToUnverified, user)

			if err = p.permissionsQ.Upsert(data.Permission{
				RequestId:   msg.RequestId,
				TelegramId:  user.TelegramId,
				AccessLevel: user.AccessLevel,
				Link:        msg.Link,
				CreatedAt:   user.CreatedAt,
			}); err != nil {
				p.log.WithError(err).Errorf("failed to upsert permission in db for message action with id `%s`", msg.RequestId)
				return errors.Wrap(err, "failed to upsert permission in db")
			}

			return nil
		})
		if err != nil {
			p.log.WithError(err).Errorf("failed to make get users transaction for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "failed to make get users transaction")
		}
	}

	err = p.sendUsers(msg.RequestId, usersToUnverified)
	if err != nil {
		p.log.WithError(err).Errorf("failed to publish users for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to publish users")
	}

	p.resetFilters()
	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return nil
}

func (p *processor) getUserFromDbByTelegramId(telegramId int64) (*data.User, error) {
	usersQ := p.usersQ.New()
	user, err := usersQ.FilterByTelegramIds(telegramId).Get()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user from db")
	}

	if user == nil {
		return nil, errors.Errorf("no such user in module")
	}

	return user, nil
}
