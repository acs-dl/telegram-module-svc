package tg

import (
	"syscall"
	"time"

	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tg) GetUserFromApi(username, phone *string) (*data.User, error) {
	user, err := t.getUserFlow(username, phone)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTg(t.tgCfg, t.log)
			t.client = cl.GetClient()
			return t.GetUserFromApi(username, phone)
		}

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
		return nil, errors.Wrap(err, "failed to get user")
	}

	t.log.Infof("successfully got user")
	return user, nil
}

func (t *tg) getUserFlow(username, phone *string) (*data.User, error) {
	var user *data.User = nil
	var err error

	if username != nil {
		user, err = t.getUserByUsername(*username)
	} else if phone != nil {
		user, err = t.getUserByPhone(*phone)
	}
	if err != nil {
		t.log.Errorf("failed to get user")
		return nil, err
	}

	return user, nil
}

func (t *tg) getUserByPhone(phone string) (*data.User, error) {
	imported, err := t.client.ContactsImportContacts([]*telegram.InputPhoneContact{{
		Phone: phone,
	}})
	if err != nil {
		t.log.Errorf("failed to search contact by phone %s", phone)
		return nil, err
	}

	if phone[0:1] == "+" {
		phone = phone[1:]
	}

	for _, user := range imported.Users {
		converted := user.(*telegram.UserObj)
		if converted.Phone == phone {
			return &data.User{
				Username:   &converted.Username,
				Phone:      &converted.Phone,
				FirstName:  converted.FirstName,
				LastName:   converted.LastName,
				TelegramId: int64(converted.ID),
				AccessHash: converted.AccessHash,
			}, nil
		}
	}

	t.log.Errorf("no user was found with phone `%s`", phone)
	return nil, nil
}

func (t *tg) getUserByUsername(username string) (*data.User, error) {
	search, err := t.client.ContactsSearch(username, 100)
	if err != nil {
		t.log.Errorf("failed to search contact by username %s", username)
		return nil, err
	}

	for _, user := range search.Users {
		converted := user.(*telegram.UserObj)
		if converted.Username == username {
			return &data.User{
				Username:   &converted.Username,
				Phone:      &converted.Phone,
				FirstName:  converted.FirstName,
				LastName:   converted.LastName,
				TelegramId: int64(converted.ID),
				AccessHash: converted.AccessHash,
			}, nil
		}
	}

	t.log.Errorf("no user was found with username `%s`", username)
	return nil, nil
}
