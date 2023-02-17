package tg

import (
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

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
