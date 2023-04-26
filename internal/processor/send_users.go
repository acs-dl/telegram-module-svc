package processor

import (
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) sendUsers(uuid string, users []data.User) error {
	unverifiedUsers := make([]data.UnverifiedUser, 0)
	for i := range users {
		if users[i].Id != nil {
			continue
		}
		permission, err := p.permissionsQ.
			FilterByTelegramIds(users[i].TelegramId).
			FilterByGreaterTime(users[i].CreatedAt).
			Get()
		if err != nil {
			p.log.WithError(err).Errorf("failed to select permissions by date `%s`", users[i].CreatedAt.String())
			return errors.Wrap(err, "failed to select permissions by date")
		}
		p.resetFilters()

		if permission == nil {
			continue
		}

		unverifiedUsers = append(unverifiedUsers, createUnverifiedUserFromModuleUser(users[i], permission.Link))
	}

	marshaledPayload, err := json.Marshal(data.UnverifiedPayload{
		Action: SetUsersAction,
		Users:  unverifiedUsers,
	})
	if err != nil {
		p.log.WithError(err).Errorf("failed to marshal unverified users list")
		return errors.Wrap(err, "failed to marshal unverified users list")
	}

	err = p.sender.SendMessageToCustomChannel(data.UnverifiedService, p.buildMessage(uuid, marshaledPayload))
	if err != nil {
		p.log.WithError(err).Errorf("failed to publish users to `telegram-module`")
		return errors.Wrap(err, "failed to publish users to `telegram-module`")
	}

	p.log.Infof("successfully published users to `telegram-module`")
	return nil
}

func (p *processor) SendDeleteUser(uuid string, user data.User) error {
	unverifiedUsers := make([]data.UnverifiedUser, 0)

	unverifiedUsers = append(unverifiedUsers, createUnverifiedUserFromModuleUser(user, ""))

	marshaledPayload, err := json.Marshal(data.UnverifiedPayload{
		Action: DeleteUsersAction,
		Users:  unverifiedUsers,
	})
	if err != nil {
		p.log.WithError(err).Errorf("failed to marshal unverified users list")
		return errors.Wrap(err, "failed to marshal unverified users list")
	}

	err = p.sender.SendMessageToCustomChannel(data.UnverifiedService, p.buildMessage(uuid, marshaledPayload))
	if err != nil {
		p.log.WithError(err).Errorf("failed to publish users to `unverified-svc`")
		return errors.Wrap(err, "failed to publish users to `unverified-svc`")
	}

	p.resetFilters()
	p.log.Infof("successfully published users to `telegram-module`")
	return nil
}

func (p *processor) buildMessage(uuid string, payload []byte) *message.Message {
	return &message.Message{
		UUID:     uuid,
		Metadata: nil,
		Payload:  payload,
	}
}

func createUnverifiedUserFromModuleUser(user data.User, submodule string) data.UnverifiedUser {
	fullName := user.FirstName + " " + user.LastName
	return data.UnverifiedUser{
		CreatedAt: user.CreatedAt,
		Module:    data.ModuleName,
		Submodule: submodule,
		ModuleId:  fmt.Sprintf("%d", user.TelegramId),
		Email:     nil,
		Name:      &fullName,
		Phone:     user.Phone,
		Username:  user.Username,
	}
}

func (p *processor) sendUpdateUserTelegram(uuid string, msg data.ModulePayload) error {
	marshaledPayload, err := json.Marshal(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to marshal update telegram info")
		return errors.Wrap(err, "failed to marshal update telegram info")
	}

	err = p.sender.SendMessageToCustomChannel(data.IdentityService, p.buildMessage(uuid, marshaledPayload))
	if err != nil {
		p.log.WithError(err).Errorf("failed to publish users to `identity-svc`")
		return errors.Wrap(err, "failed to publish users to `identity-svc`")
	}

	p.log.Infof("successfully published user to `identity-svc`")
	return nil
}
