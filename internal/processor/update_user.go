package processor

import (
	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/helpers"
	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	"github.com/acs-dl/telegram-module-svc/internal/tg_client"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateUpdateUser(msg data.ModulePayload) error {
	phoneValidationCase := validation.When(msg.Username == nil, validation.Required.Error("phone is required if username is not set"))
	usernameValidationCase := validation.When(msg.Phone == nil, validation.Required.Error("username is required if phone is not set"))

	return validation.Errors{
		"link":         validation.Validate(msg.Link, validation.Required),
		"id":           validation.Validate(msg.SubmoduleId, validation.Required),
		"username":     validation.Validate(msg.Username, usernameValidationCase),
		"phone":        validation.Validate(msg.Phone, phoneValidationCase),
		"access_level": validation.Validate(msg.AccessLevel, validation.Required),
	}.Filter()
}

func (p *processor) HandleUpdateUserAction(msg data.ModulePayload) (string, error) {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateUpdateUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to validate fields")
	}
	user, err := p.checkUserExistence(msg.Username, msg.Phone)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to get user")
	}

	err = p.updateRemotePermission(msg, *user)
	if err != nil {
		p.log.WithError(err).Errorf("failed to update permission from API for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to update permission from api")
	}

	if err = p.permissionsQ.UpdateAccessLevel(data.Permission{
		TelegramId:  user.TelegramId,
		AccessLevel: msg.AccessLevel,
		Link:        msg.Link,
	}); err != nil {
		p.log.WithError(err).Errorf("failed to update user in db for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to update user in db")
	}

	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return data.SUCCESS, nil
}

func (p *processor) updateRemotePermission(msg data.ModulePayload, user data.User) error {
	chat, err := p.getChatForUser(msg, user)
	if err != nil {
		return errors.Wrap(err, "failed to get chat for user from api")
	}

	err = helpers.GetRequestError(p.pqueues.SuperUserPQueue, any(p.telegramClient.UpdateUserInChatFromApi), []any{any(user), any(*chat)}, pqueue.NormalPriority)
	if err != nil {
		return errors.Wrap(err, "failed to update user from api")
	}

	return nil
}

func (p *processor) getChatForUser(msg data.ModulePayload, user data.User) (*tg_client.Chat, error) {
	chats, err := helpers.GetChats(p.pqueues.SuperUserPQueue, any(p.telegramClient.GetChatFromApi), []any{any(msg.Link)}, pqueue.NormalPriority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chats from api")
	}

	chat := helpers.RetrieveChat(chats, msg)

	if chat == nil {
		return nil, errors.New("no chat was found")
	}

	chatUser, err := helpers.GetUser(p.pqueues.SuperUserPQueue, any(p.telegramClient.GetChatUserFromApi), []any{any(user), any(*chat)}, pqueue.NormalPriority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user from api")
	}

	if chatUser == nil {
		return nil, errors.New("user is not in chat")
	}

	return chat, nil
}
