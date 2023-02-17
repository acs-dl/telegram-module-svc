package processor

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strconv"
	"time"
)

func (p *processor) validateAddUser(msg data.ModulePayload) error {
	return validation.Errors{
		"link":     validation.Validate(msg.Link, validation.Required),
		"username": validation.Validate(msg.Username, validation.Required),
		"phone":    validation.Validate(msg.Phone, validation.Required),
		"user_id":  validation.Validate(msg.UserId, validation.Required),
	}.Filter()
}

func (p *processor) handleAddUserAction(msg data.ModulePayload) error {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateAddUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to validate fields")
	}

	userId, err := strconv.ParseInt(msg.UserId, 10, 64)
	if err != nil {
		p.log.WithError(err).Errorf("failed to parse user id `%s` for message action with id `%s`", msg.UserId, msg.RequestId)
		return errors.Wrap(err, "failed to parse user id")
	}

	err = p.telegramClient.AddUserInChatFromApi(&msg.Username, &msg.Phone, msg.Link)
	if err != nil {
		p.log.WithError(err).Errorf("failed to add user from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to add user from api")
	}

	user, err := p.telegramClient.GetChatUserFromApi(&msg.Username, &msg.Phone, msg.Link)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get chat user from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get chat user from api")
	}
	if user == nil {
		p.log.WithError(err).Errorf("something wrong with user from api for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "something wrong with user from api")
	}
	user.CreatedAt = time.Now()
	user.Id = &userId

	err = p.managerQ.Transaction(func() error {
		if err = p.usersQ.Upsert(*user); err != nil {
			p.log.WithError(err).Errorf("failed to upsert user in db for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "failed to upsert user in db")
		}

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
		p.log.WithError(err).Errorf("failed to make add user transaction for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to make add user transaction")
	}

	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return nil
}
