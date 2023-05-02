package tg_client

import (
	"syscall"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
	pkgErrors "github.com/pkg/errors"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tgInfo) GetUserFromApi(client *telegram.Client, username, phone *string) (*data.User, error) {
	user, err := t.getUserFlow(client, username, phone)
	if err != nil {
		if pkgErrors.Is(err, syscall.EPIPE) {
			cl := NewTgAsInterface(t.cfg, t.ctx).(TelegramClient)
			t.usualClient = cl.GetTg().usualClient
			return t.GetUserFromApi(client, username, phone)
		}

		if tgerr.IsCode(err, 420) {
			duration, ok := tgerr.AsFloodWait(err)
			if !ok {
				return nil, errors.New("failed to convert flood error")
			}
			t.log.Warnf("we need to wait `%s`", duration)
			time.Sleep(duration)
			return t.GetUserFromApi(client, username, phone)
		}

		t.log.WithError(err).Errorf("failed to get user")
		return nil, errors.Wrap(err, "failed to get user")
	}

	t.log.Infof("successfully got user")
	return user, nil
}

func (t *tgInfo) getUserFlow(client *telegram.Client, username, phone *string) (*data.User, error) {
	var user *data.User = nil
	var err error

	if username != nil {
		user, err = t.getUserByUsername(client, *username)
	} else if phone != nil {
		user, err = t.getUserByPhone(client, *phone)
	}
	if err != nil {
		t.log.Errorf("failed to get user")
		return nil, err
	}

	return user, nil
}

func (t *tgInfo) getUserByPhone(client *telegram.Client, phone string) (*data.User, error) {
	imported, err := client.API().ContactsResolvePhone(t.ctx, phone)
	//imported, err := t.usualClient.API().ContactsImportContacts(t.ctx, []tg.InputPhoneContact{
	//	{Phone: phone},
	//})
	if err != nil {
		return nil, err
	}
	if err != nil {
		t.log.Errorf("failed to search contact by phone %s", phone)
		return nil, err
	}

	if phone[0:1] == "+" {
		phone = phone[1:]
	}

	for _, user := range imported.Users {
		converted := user.(*tg.User)
		if converted.Phone == phone {
			return &data.User{
				Username:   &converted.Username,
				Phone:      &converted.Phone,
				FirstName:  converted.FirstName,
				LastName:   converted.LastName,
				TelegramId: converted.ID,
				AccessHash: converted.AccessHash,
			}, nil
		}
	}

	t.log.Errorf("no user was found with phone `%s`", phone)
	return nil, nil
}

func (t *tgInfo) getUserByUsername(client *telegram.Client, username string) (*data.User, error) {
	search, err := client.API().ContactsResolveUsername(t.ctx, username)
	//search, err := t.usualClient.API().ContactsSearch(t.ctx, &tg.ContactsSearchRequest{Q: username, Limit: 10})
	if err != nil {
		t.log.Errorf("failed to search contact by username %s", username)
		return nil, err
	}

	for _, user := range search.Users {
		converted := user.(*tg.User)
		if converted.Username == username {
			return &data.User{
				Username:   &converted.Username,
				Phone:      &converted.Phone,
				FirstName:  converted.FirstName,
				LastName:   converted.LastName,
				TelegramId: converted.ID,
				AccessHash: converted.AccessHash,
			}, nil
		}
	}

	t.log.Errorf("no user was found with username `%s`", username)
	return nil, nil
}
