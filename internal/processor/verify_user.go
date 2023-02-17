package processor

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strconv"
)

func (p *processor) validateVerifyUser(msg data.ModulePayload) error {
	return validation.Errors{
		"user_id":  validation.Validate(msg.UserId, validation.Required),
		"username": validation.Validate(msg.Username, validation.Required),
		"phone":    validation.Validate(msg.Phone, validation.Required),
	}.Filter()
}

func (p *processor) handleVerifyUserAction(msg data.ModulePayload) error {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateVerifyUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to validate fields")
	}

	userId, err := strconv.ParseInt(msg.UserId, 10, 64)
	if err != nil {
		p.log.WithError(err).Errorf("failed to parse user id `%s` for message action with id `%s`", msg.UserId, msg.RequestId)
		return errors.Wrap(err, "failed to parse user id")
	}

	user, err := p.telegramClient.GetUserFromApi(&msg.Username, &msg.Phone)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "some error while getting user from api")
	}
	user.Id = &userId

	if err = p.usersQ.Upsert(*user); err != nil {
		p.log.WithError(err).Errorf("failed to upsert user in db for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to upsert user in db")
	}

	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return nil
}
