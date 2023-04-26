package processor

import (
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/helpers"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateAddUser(msg data.ModulePayload) error {
	phoneValidationCase := validation.When(msg.Username == nil, validation.Required.Error("phone is required if username is not set"))
	usernameValidationCase := validation.When(msg.Phone == nil, validation.Required.Error("username is required if phone is not set"))

	return validation.Errors{
		"link":     validation.Validate(msg.Link, validation.Required),
		"username": validation.Validate(msg.Username, usernameValidationCase),
		"phone":    validation.Validate(msg.Phone, phoneValidationCase),
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

	user, err := helpers.GetUser(p.pqueues.UsualPQueue,
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

	err = helpers.GetRequestError(p.pqueues.SuperPQueue, any(p.telegramClient.AddUserInChatFromApi), []any{any(*user), any(*chat)}, pqueue.NormalPriority)
	if err != nil {
		p.log.WithError(err).Errorf("failed to add user in chat from API for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to add user in chat from api")
	}

	//when we add user is always member
	user.AccessLevel = data.Member
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
		p.log.WithError(err).Errorf("failed to publish users for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to publish users")
	}

	p.resetFilters()
	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return nil
}
