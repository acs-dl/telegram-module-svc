package processor

import (
	"fmt"
	"strconv"

	"github.com/acs-dl/telegram-module-svc/internal/data"
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

func ConvertIdentifiersStringsToInt(userIdStr, submoduleIdStr string, submoduleAccessHashStr *string) (userId, submoduleId int64, submoduleAccessHash *int64, err error) {
	userId, err = strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return 0, 0, nil, errors.Wrap(err, "failed to parse user id")
	}

	submoduleId, err = strconv.ParseInt(submoduleIdStr, 10, 64)
	if err != nil {
		return 0, 0, nil, errors.Wrap(err, "failed to parse submodule id")
	}

	if submoduleAccessHashStr == nil {
		return userId, submoduleId, nil, nil
	}

	tmp, err := strconv.ParseInt(*submoduleAccessHashStr, 10, 64)
	if err != nil {
		return 0, 0, nil, errors.Wrap(err, "failed to parse submodule id")
	}
	submoduleAccessHash = &tmp

	return userId, submoduleId, submoduleAccessHash, nil
}
