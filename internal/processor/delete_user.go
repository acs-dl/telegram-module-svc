package processor

import (
	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/helpers"
	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateDeleteUser(msg data.ModulePayload) error {
	phoneValidationCase := validation.When(msg.Username == nil, validation.Required.Error("phone is required if username is not set"))
	usernameValidationCase := validation.When(msg.Phone == nil, validation.Required.Error("username is required if phone is not set"))

	return validation.Errors{
		"username": validation.Validate(msg.Username, usernameValidationCase),
		"phone":    validation.Validate(msg.Phone, phoneValidationCase),
	}.Filter()
}

func (p *processor) HandleDeleteUserAction(msg data.ModulePayload) (string, error) {
	p.log.Infof("start handle message action with id `%s`", msg.RequestId)

	err := p.validateDeleteUser(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate fields for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to validate fields")
	}

	user, err := p.checkUserExistence(msg.Username, msg.Phone)
	if err != nil {
		p.log.WithError(err).Errorf("failed to get user for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to get user")
	}

	permissions, err := p.permissionsQ.FilterByTelegramIds(user.TelegramId).Select()
	if err != nil {
		p.log.WithError(err).Errorf("failed to select permissions by telegram id `%d` for message action with id `%s`", user.TelegramId, msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to select permissions")
	}

	for _, permission := range permissions {
		err = p.deleteRemotePermission(permission.Link, permission.SubmoduleId, permission.SubmoduleAccessHash, *user)
		if err != nil {
			p.log.WithError(err).Errorf("failed to remove permission from API for message action with id `%s`", msg.RequestId)
			return data.FAILURE, errors.Wrap(err, "some error while removing permission from api")
		}

		if err = p.permissionsQ.FilterByTelegramIds(permission.TelegramId).FilterByLinks(permission.Link).Delete(); err != nil {
			p.log.WithError(err).Errorf("failed to delete permission from db for message action with id `%s`", msg.RequestId)
			return data.FAILURE, errors.Wrap(err, "failed to delete permission")
		}
	}

	err = p.usersQ.FilterByTelegramIds(user.TelegramId).Delete()
	if err != nil {
		p.log.WithError(err).Errorf("failed to delete user by telegram id `%d` for message action with id `%s`", user.TelegramId, msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to delete user")
	}

	err = p.sendDeleteInUnverifiedOrUpdateInIdentity(msg.RequestId, *user)
	if err != nil {
		p.log.WithError(err).Errorf("failed to send delete unverified or update identity for message action with id `%s`", msg.RequestId)
		return data.FAILURE, errors.Wrap(err, "failed to send delete unverified or update identity")
	}

	p.log.Infof("finish handle message action with id `%s`", msg.RequestId)
	return data.SUCCESS, nil
}

func (p *processor) checkUserExistence(username, phone *string) (*data.User, error) {
	user, err := helpers.GetUser(p.pqueues.UserPQueue,
		any(p.telegramClient.GetUserFromApi),
		[]any{
			any(p.telegramClient.GetSuperClient()),
			any(username),
			any(phone),
		},
		pqueue.NormalPriority,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user from api")
	}

	if user == nil {
		return nil, errors.New("no user was found")
	}

	dbUser, err := p.getUserFromDbByTelegramId(user.TelegramId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user from service")
	}

	return dbUser, nil
}

func (p *processor) deleteRemotePermission(link string, submoduleId int64, submoduleAccessHash *int64, user data.User) error {
	chat, err := p.getChatForUser(link, submoduleId, submoduleAccessHash, user)
	if err != nil {
		return errors.Wrap(err, "failed to get chat for user from api")
	}

	err = helpers.GetRequestError(p.pqueues.SuperUserPQueue, any(p.telegramClient.DeleteFromChatFromApi), []any{any(user), any(*chat)}, pqueue.NormalPriority)
	if err != nil {
		return errors.Wrap(err, "some error while removing user from api")
	}

	return nil
}
