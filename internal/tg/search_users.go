package tg

import (
	"fmt"
	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"time"
)

func (t *tg) SearchByFromApi(username, phone *string, amount int64) ([]data.User, error) {
	users, err := t.searchByFlow(username, phone, amount)
	if err != nil {
		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			t.log.WithError(err).Errorf("failed to search for users, some strange error")
			return nil, errors.Wrap(err, "failed to search for users, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.SearchByFromApi(username, phone, amount)
		}

		t.log.WithError(err).Errorf("failed to search for users")
		return nil, errors.Wrap(err, fmt.Sprintf("failed to search for users"))
	}

	t.log.Infof("successfully searched users")
	return users, nil
}

func (t *tg) searchByFlow(username, phone *string, amount int64) ([]data.User, error) {
	var users []data.User
	var err error
	if username != nil {
		users, err = t.searchUsersByUsername(*username, amount)
	}
	if phone != nil {
		users, err = t.searchUsersByPhone(*phone, amount)
	}
	if err != nil {
		t.log.WithError(err).Errorf("failed to search users")
		return nil, err
	}

	return users, nil
}

func (t *tg) searchUsersByPhone(phone string, amount int64) ([]data.User, error) {
	var users []data.User

	imported, err := t.client.ContactsImportContacts([]*telegram.InputPhoneContact{{
		Phone: phone,
	}})
	if err != nil {
		t.log.WithError(err).Errorf("failed to search contact by phone %s", phone)
		return nil, err
	}

	if phone[0:1] == "+" {
		phone = phone[1:]
	}

	for i, user := range imported.Users {
		tgUser := user.(*telegram.UserObj)
		users = append(users, data.User{
			Username:   tgUser.Username,
			Phone:      tgUser.Phone,
			FirstName:  tgUser.FirstName,
			LastName:   tgUser.LastName,
			TelegramId: int64(tgUser.ID),
			AccessHash: tgUser.AccessHash,
		})
		if int64(i) == amount {
			break
		}
	}

	t.log.Errorf("found %v users with phone `%s`", len(users), phone)
	return users, nil
}

func (t *tg) searchUsersByUsername(username string, amount int64) ([]data.User, error) {
	var users []data.User

	search, err := t.client.ContactsSearch(username, int32(amount))
	if err != nil {
		t.log.WithError(err).Errorf("failed to search contact by username %s", username)
		return nil, err
	}

	for _, user := range search.Users {
		tgUser := user.(*telegram.UserObj)
		users = append(users, data.User{
			Username:   tgUser.Username,
			Phone:      tgUser.Phone,
			FirstName:  tgUser.FirstName,
			LastName:   tgUser.LastName,
			TelegramId: int64(tgUser.ID),
			AccessHash: tgUser.AccessHash,
		})
	}

	t.log.Errorf("found %v users with username `%s`", len(users), username)
	return users, nil
}
