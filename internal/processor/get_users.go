package processor

import (
	"strings"
	"time"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/helpers"
	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateGetUsers(msg data.ModulePayload) error {
	return validation.Errors{
		"link": validation.Validate(msg.Link, validation.Required),
	}.Filter()
}

func (p *processor) HandleGetUsersAction(msg data.ModulePayload) error {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateGetUsers(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to validate fields")
	}

	chats, err := helpers.GetChats(p.pqueues.SuperUserPQueue, any(p.telegramClient.GetChatFromApi), []any{any(msg.Link)}, pqueue.LowPriority)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get chat from api for message action with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to get chat from api")
	}

	for _, chat := range chats {
		err = p.handleDBChat(&chat)
		if err != nil {
			p.log.WithError(err).Errorf("failed to handle db chat flow for message action with id `%s`", msg.RequestId)
			return errors.Wrap(err, "failed to handle db chat flow")
		}

		users, err := helpers.GetUsers(p.pqueues.SuperUserPQueue, any(p.telegramClient.GetChatUsersFromApi), []any{any(chat)}, pqueue.LowPriority)
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
					RequestId:           msg.RequestId,
					TelegramId:          user.TelegramId,
					AccessLevel:         user.AccessLevel,
					Link:                msg.Link,
					CreatedAt:           user.CreatedAt,
					SubmoduleAccessHash: chat.AccessHash,
					SubmoduleId:         chat.Id,
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

	}
	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return nil
}

func (p *processor) handleDBChat(chat *tg_client.Chat) error {
	dbChat, err := p.chatsQ.FilterByIds(chat.Id).FilterByAccessHash(chat.AccessHash).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get chat from database")
	}

	var photoName *string = nil
	if dbChat != nil {
		photoName = dbChat.PhotoName
	}

	chatPhotoLink, err := helpers.GetString(
		p.pqueues.SuperUserPQueue,
		any(p.telegramClient.GetChatPhotoFromApi),
		[]any{any(photoName), any(*chat)},
		pqueue.LowPriority,
	)
	if err != nil {
		return errors.Wrap(err, "failed to get chat photo from api")
	}

	var photoLink *string = nil
	if chatPhotoLink != "" {
		splitLink := strings.Split(chatPhotoLink, "/")
		splitName := strings.Split(splitLink[len(splitLink)-1], ".")

		photoName = &splitName[0]
		photoLink = &chatPhotoLink
	}

	err = p.chatsQ.Upsert(data.Chat{
		Title:         chat.Title,
		Id:            chat.Id,
		AccessHash:    chat.AccessHash,
		MembersAmount: chat.MembersAmount,
		PhotoName:     photoName,
		PhotoLink:     photoLink,
	})
	if err != nil {
		return errors.Wrap(err, "failed to upsert chat")
	}

	return nil
}
