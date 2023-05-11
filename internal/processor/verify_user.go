package processor

import (
	"strconv"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/helpers"
	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	validation "github.com/go-ozzo/ozzo-validation/v4"
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

func (p *processor) HandleVerifyUserAction(msg data.ModulePayload) (string, error) {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateVerifyUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to validate fields")
	}

	userId, err := strconv.ParseInt(msg.UserId, 10, 64)
	if err != nil {
		p.log.WithError(err).Errorf("failed to parse user id `%s` for message action with id `%s`", msg.UserId, msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to parse user id")
	}

	user, err := helpers.GetUser(p.pqueues.UserPQueue,
		any(p.telegramClient.GetUserFromApi),
		[]any{
			any(p.telegramClient.GetSuperClient()),
			any(msg.Username),
			any(msg.Phone),
		},
		pqueue.NormalPriority,
	)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user from api for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to get user from api")
	}

	if user == nil {
		p.log.Errorf("no user was found for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.New("no user was found")
	}

	user.Id = &userId

	if err = p.usersQ.Upsert(*user); err != nil {
		p.log.WithError(err).Errorf("failed to upsert user in db for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to upsert user in db")
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
		return data.FAILURE, errors.Wrap(err, "failed to publish users")
	}

	err = p.SendDeleteUser(msg.RequestId, *user)
	if err != nil {
		p.log.WithError(err).Errorf("failed to publish delete user for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to publish delete user")
	}

	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return data.SUCCESS, nil
}
