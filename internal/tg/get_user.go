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

func (t *tg) GetUserFromApi(username, phone *string) (*data.User, error) {
	user, err := t.getUserFlow(username, phone)
	if err != nil {
		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			t.log.WithError(err).Errorf("failed to get user, some strange error")
			return nil, errors.Wrap(err, "failed to get user, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.GetUserFromApi(username, phone)
		}

		t.log.WithError(err).Errorf("failed to get user")
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get user"))
	}

	t.log.Infof("successfully got user")
	return user, nil
}

func (t *tg) getUserFlow(username, phone *string) (*data.User, error) {
	inputUser, err := t.getInputUser(username, phone)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get input user")
		return nil, err
	}

	tgFullUser, err := t.client.UsersGetFullUser(inputUser)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get full user")
		return nil, err
	}

	tgUser := tgFullUser.User.(*telegram.UserObj)

	return &data.User{
		Username:   &tgUser.Username,
		Phone:      &tgUser.Phone,
		FirstName:  tgUser.FirstName,
		LastName:   tgUser.LastName,
		TelegramId: int64(tgUser.ID),
		AccessHash: tgUser.AccessHash,
	}, nil
}

func (t *tg) getUserByPhone(phone string) (*telegram.InputUserObj, error) {
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

	for _, user := range imported.Users {
		converted := user.(*telegram.UserObj)
		if converted.Phone == phone {
			return &telegram.InputUserObj{
				UserID:     converted.ID,
				AccessHash: converted.AccessHash,
			}, nil
		}
	}

	t.log.Errorf("no user was found with phone `%s`", phone)
	return nil, errors.Errorf("no user was found with phone `%s`", phone)
}

func (t *tg) getUserByUsername(username string) (*telegram.InputUserObj, error) {
	search, err := t.client.ContactsSearch(username, 100)
	if err != nil {
		t.log.WithError(err).Errorf("failed to search contact by username %s", username)
		return nil, err
	}

	for _, user := range search.Users {
		converted := user.(*telegram.UserObj)
		if converted.Username == username {
			return &telegram.InputUserObj{
				UserID:     converted.ID,
				AccessHash: converted.AccessHash,
			}, nil
		}
	}

	t.log.Errorf("no user was found with username `%s`", username)
	return nil, errors.Errorf("no user was found with username `%s`", username)
}
