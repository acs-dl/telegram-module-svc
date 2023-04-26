package processor

import (
	"fmt"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) getUserFromDbByTelegramId(telegramId int64) (*data.User, error) {
	usersQ := p.usersQ.New()
	user, err := usersQ.FilterByTelegramIds(telegramId).Get()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user from db")
	}

	if user == nil {
		return nil, errors.Errorf("no such user in module")
	}

	return user, nil
}

func (p *processor) resetFilters() {
	p.usersQ = p.usersQ.New()
	p.permissionsQ = p.permissionsQ.New()
}

func (p *processor) sendDeleteInUnverifiedOrUpdateInIdentity(requestId string, user data.User) error {
	if user.Id == nil {
		err := p.SendDeleteUser(requestId, user)
		if err != nil {
			return errors.Wrap(err, "failed to publish delete telegram user in telegram-module")
		}
	} else {
		err := p.sendUpdateUserTelegram(requestId, data.ModulePayload{
			RequestId: requestId,
			UserId:    fmt.Sprintf("%d", *user.Id),
			Action:    RemoveTelegramAction,
		})
		if err != nil {
			return errors.Wrap(err, "failed to publish update telegram user in identity-svc")
		}
	}

	return nil
}
