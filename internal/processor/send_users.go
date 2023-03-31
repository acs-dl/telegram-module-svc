package processor

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"time"
)

func (p *processor) sendUsers(uuid string, borderTime time.Time) error {
	users, err := p.usersQ.FilterByTime(borderTime).FilterById(nil).Select()
	if err != nil {
		p.log.WithError(err).Errorf("failed to select users by date `%s`", borderTime.String())
		return errors.Wrap(err, "failed to select users by date")
	}

	unverifiedUsers := make([]data.UnverifiedUser, 0)
	for i := range users {
		if users[i].Id != nil {
			continue
		}

		unverifiedUsers = append(unverifiedUsers, createUnverifiedUserFromModuleUser(users[i]))
	}

	err = p.sender.SendMessageToCustomChannel("unverified-svc", p.buildUnverifiedUserListMessage(uuid, data.UnverifiedPayload{
		Action: SetUsersAction,
		Users:  unverifiedUsers,
	}))
	if err != nil {
		p.log.WithError(err).Errorf("failed to publish users to `unverified-svc`")
		return errors.Wrap(err, "failed to publish users to `unverified-svc`")
	}

	p.resetFilters()
	p.log.Infof("successfully published users to `unverified-svc`")
	return nil
}

func (p *processor) sendDeleteUser(uuid string, user data.User) error {
	unverifiedUsers := make([]data.UnverifiedUser, 0)

	unverifiedUsers = append(unverifiedUsers, createUnverifiedUserFromModuleUser(user))

	err := p.sender.SendMessageToCustomChannel("unverified-svc", p.buildUnverifiedUserListMessage(uuid, data.UnverifiedPayload{
		Action: DeleteUsersAction,
		Users:  unverifiedUsers,
	}))
	if err != nil {
		p.log.WithError(err).Errorf("failed to publish users to `unverified-svc`")
		return errors.Wrap(err, "failed to publish users to `unverified-svc`")
	}

	p.resetFilters()
	p.log.Infof("successfully published users to `unverified-svc`")
	return nil
}

func (p *processor) buildUnverifiedUserListMessage(uuid string, unverifiedPayload data.UnverifiedPayload) *message.Message {
	marshaled, err := json.Marshal(unverifiedPayload)
	if err != nil {
		p.log.WithError(err).Errorf("failed to marshal response")
	}

	return &message.Message{
		UUID:     uuid,
		Metadata: nil,
		Payload:  marshaled,
	}
}

func createUnverifiedUserFromModuleUser(user data.User) data.UnverifiedUser {
	name := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	return data.UnverifiedUser{
		CreatedAt: user.CreatedAt,
		Module:    data.ModuleName,
		ModuleId:  fmt.Sprintf("%d", user.TelegramId),
		Email:     nil,
		Name:      &name,
		Phone:     user.Phone,
		Username:  user.Username,
	}
}
