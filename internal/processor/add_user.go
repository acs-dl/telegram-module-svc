package processor

import (
	"strconv"
	"time"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/helpers"
	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	"github.com/acs-dl/telegram-module-svc/internal/tg_client"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateAddUser(msg data.ModulePayload) error {
	phoneValidationCase := validation.When(msg.Username == nil, validation.Required.Error("phone is required if username is not set"))
	usernameValidationCase := validation.When(msg.Phone == nil, validation.Required.Error("username is required if phone is not set"))

	return validation.Errors{
		"link":     validation.Validate(msg.Link, validation.Required),
		"id":       validation.Validate(msg.SubmoduleId, validation.Required),
		"username": validation.Validate(msg.Username, usernameValidationCase),
		"phone":    validation.Validate(msg.Phone, phoneValidationCase),
		"user_id":  validation.Validate(msg.UserId, validation.Required),
	}.Filter()
}

func (p *processor) HandleAddUserAction(msg data.ModulePayload) (string, error) {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateAddUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to validate fields")
	}

	userId, err := strconv.ParseInt(msg.UserId, 10, 64)
	if err != nil {
		p.log.WithError(err).Errorf("failed to parse user id `%s` for message action with id `%s`", msg.UserId, msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to parse user id")
	}

	requestStatus, user, err := p.addUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to add user for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to add user")
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
			RequestId:           msg.RequestId,
			TelegramId:          user.TelegramId,
			AccessLevel:         user.AccessLevel,
			Link:                msg.Link,
			CreatedAt:           user.CreatedAt,
			SubmoduleAccessHash: msg.SubmoduleAccessHash,
			SubmoduleId:         msg.SubmoduleId,
		}); err != nil {
			p.log.WithError(err).Errorf("failed to upsert permission in db for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "failed to upsert permission in db")
		}

		return nil
	})
	if err != nil {
		p.log.WithError(err).Errorf("failed to make add user transaction for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to make add user transaction")
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
		p.log.WithError(err).Errorf("failed to publish users for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to publish users")
	}

	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return requestStatus, nil
}

func (p *processor) addUser(msg data.ModulePayload) (string, *data.User, error) {
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
		return data.FAILURE, nil, errors.Wrap(err, "failed to get user from api")
	}

	if user == nil {
		return data.FAILURE, nil, errors.New("no user was found")
	}

	chats, err := helpers.GetChats(p.pqueues.SuperUserPQueue, any(p.telegramClient.GetChatFromApi), []any{any(msg.Link)}, pqueue.NormalPriority)
	if err != nil {
		return data.FAILURE, nil, errors.Wrap(err, "failed to get chat from api")
	}

	chat := helpers.RetrieveChat(chats, msg)

	if chat == nil {
		return data.FAILURE, nil, errors.New("no chat was found")
	}

	err = helpers.GetRequestError(p.pqueues.SuperUserPQueue, any(p.telegramClient.AddUserInChatFromApi), []any{any(*user), any(*chat)}, pqueue.NormalPriority)
	if err == nil {
		return data.SUCCESS, user, nil
	}

	if !tgerr.Is(err, tg.ErrUserNotMutualContact, tg.ErrUserPrivacyRestricted) {
		return data.FAILURE, nil, errors.Wrap(err, "failed to add user in chat from api")
	}

	err = p.sendInviteMessage(msg.Link, *user, *chat)
	if err != nil {
		return data.FAILURE, nil, errors.Wrap(err, "failed to send invite in chat from api")
	}

	return data.INVITED, user, nil
}
func (p *processor) sendInviteMessage(link string, user data.User, chat tg_client.Chat) error {
	inviteLink, err := helpers.GetString(
		p.pqueues.SuperUserPQueue,
		any(p.telegramClient.GenerateChatLinkFromApi),
		[]any{any(chat)}, pqueue.NormalPriority)
	if err != nil {
		return err
	}

	var userName string
	switch {
	case len(user.FirstName+user.LastName) != 0:
		userName = user.FirstName + " " + user.LastName
	case user.Username != nil:
		userName = *user.Username
	case user.Phone != nil:
		userName = *user.Phone
	default:
		p.log.Warnf("failed to get user name")
		userName = "from ACS"
	}

	err = helpers.GetRequestError(p.pqueues.SuperUserPQueue, any(p.telegramClient.SendMessageFromApi),
		[]any{
			any(data.MessageInfo{
				Data: map[string]interface{}{
					"Name":       userName,
					"GroupName":  link,
					"InviteLink": inviteLink,
				},
				MessageTemplate: data.InviteMessageTemplate,
				User:            user,
			}),
		}, pqueue.NormalPriority)
	if err != nil {
		return err
	}

	return nil
}
