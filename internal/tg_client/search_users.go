package tg_client

import (
	"syscall"
	"time"

	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) SearchByFromApi(username, phone *string, amount int) ([]data.User, error) {
	users, err := t.searchByFlow(username, phone, amount)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.userClient = cl.GetTg().userClient
			return t.SearchByFromApi(username, phone, amount)
		}

		duration, isFlood := tgerr.AsFloodWait(err)
		if isFlood {
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.SearchByFromApi(username, phone, amount)
		}

		t.log.WithError(err).Errorf("failed to search for users")
		return nil, errors.Wrap(err, "failed to search for users")
	}

	t.log.Infof("successfully searched users")
	return users, nil
}

func (t *tgInfo) searchByFlow(username, phone *string, amount int) ([]data.User, error) {
	users := make([]data.User, 0)
	var err error

	if username != nil {
		users, err = t.searchUsersByUsername(*username, amount)
	} else if phone != nil {
		users, err = t.searchUsersByPhone(*phone, amount)
	}
	if err != nil {
		t.log.Errorf("failed to search users")
		return nil, err
	}

	return users, nil
}

func (t *tgInfo) searchUsersByPhone(phone string, amount int) ([]data.User, error) {
	var users []data.User

	imported, err := t.userClient.API().ContactsImportContacts(t.ctx, []tg.InputPhoneContact{
		{Phone: phone},
	})
	if err != nil {
		t.log.Errorf("failed to search contact by phone %s", phone)
		return nil, err
	}

	if phone[0:1] == "+" {
		phone = phone[1:]
	}

	for i, user := range imported.Users {
		tgUser := user.(*tg.User)
		users = append(users, data.User{
			Username:   &tgUser.Username,
			Phone:      &tgUser.Phone,
			FirstName:  tgUser.FirstName,
			LastName:   tgUser.LastName,
			TelegramId: tgUser.ID,
			AccessHash: tgUser.AccessHash,
		})
		if i == amount {
			break
		}
	}

	t.log.Infof("found %v users with phone `%s`", len(users), phone)
	return users, nil
}

func (t *tgInfo) searchUsersByUsername(username string, amount int) ([]data.User, error) {
	var users []data.User

	search, err := t.userClient.API().ContactsSearch(t.ctx, &tg.ContactsSearchRequest{Q: username, Limit: amount})
	if err != nil {
		t.log.Errorf("failed to search contact by username %s", username)
		return nil, err
	}

	for _, user := range search.Users {
		tgUser := user.(*tg.User)
		users = append(users, data.User{
			Username:   &tgUser.Username,
			Phone:      &tgUser.Phone,
			FirstName:  tgUser.FirstName,
			LastName:   tgUser.LastName,
			TelegramId: tgUser.ID,
			AccessHash: tgUser.AccessHash,
		})
	}

	t.log.Infof("found %v users with username `%s`", len(users), username)
	return users, nil
}
