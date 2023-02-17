package tg

import (
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
)

func (t *tg) checkUserInChat(id int32, hashID *int64, userId int64) (*data.User, error) {
	var users, err = t.getAllUsers(id, hashID)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get all users")
		return nil, err
	}

	for _, user := range users {
		if userId == user.TelegramId {
			return &user, nil
		}
	}

	return nil, nil
}
