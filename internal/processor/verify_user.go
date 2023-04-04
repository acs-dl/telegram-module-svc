package processor

import (
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateVerifyUser(msg data.ModulePayload) error {
	phoneValidationCase := validation.When(msg.Username == nil, validation.Required.Error("username is required if phone is not set"))
	usernameValidationCase := validation.When(msg.Phone == nil, validation.Required.Error("phone is required if username is not set"))

	return validation.Errors{
		"user_id":  validation.Validate(msg.UserId, validation.Required),
		"username": validation.Validate(msg.Username, usernameValidationCase),
		"phone":    validation.Validate(msg.Phone, phoneValidationCase),
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

	item := p.addFunctionInPqueue(any(p.telegramClient.GetUserFromApi), []any{any(msg.Username), any(msg.Phone)}, 10)
	err = item.Response.Error
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "some error while getting user from api")
	}
	user, err := p.convertUserFromInterfaceAndCheck(item.Response.Value)
	if err != nil {
		p.log.WithError(err).Errorf("something wrong with user for message action with id `%s`", msg.RequestId)
		return errors.Errorf("something wrong with user from api")
	}
	user.Id = &userId

	if err = p.usersQ.Upsert(*user); err != nil {
		p.log.WithError(err).Errorf("failed to upsert user in db for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to upsert user in db")
	}

	err = p.sendUpdateUserTelegram(msg.RequestId, data.ModulePayload{
		RequestId: msg.RequestId,
		UserId:    msg.UserId,
		Username:  msg.Username,
		Phone:     msg.Phone,
		Action:    UpdateTelegramAction,
	})
	if err != nil {
		p.log.WithError(err).Errorf("failed to publish users for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to publish users")
	}

	err = p.SendDeleteUser(msg.RequestId, *user)
	if err != nil {
		p.log.WithError(err).Errorf("failed to publish delete user for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to publish delete user")
	}

	p.resetFilters()
	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return nil
}
