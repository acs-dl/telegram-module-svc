package tg

import (
	"fmt"
	"time"

	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (t *tg) AddUserInChatFromApi(username, phone *string, title string) error {
	err := t.addUserFlow(username, phone, title)
	if err != nil {
		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			t.log.WithError(err).Errorf("failed to add user, some strange error")
			return errors.Wrap(err, "failed to add user, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.AddUserInChatFromApi(username, phone, title)
		}

		t.log.WithError(err).Errorf("failed to add user in %s", title)
		return errors.Wrap(err, fmt.Sprintf("failed to add userin %s", title))
	}

	t.log.Infof("successfully add user in %s", title)
	return nil
}

func (t *tg) addUserFlow(username, phone *string, title string) error {
	id, accessHash, err := t.findChatByTitle(title)
	if err != nil {
		t.log.WithError(err).Errorf("failed to find chat %s", title)
		return err
	}

	inputUser, err := t.getInputUser(username, phone)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get input user")
		return err
	}

	if err = t.addUser(inputUser, *id, accessHash); err != nil {
		t.log.WithError(err).Errorf("failed to add user")
		return err
	}

	return nil
}

func (t *tg) addUser(user *telegram.InputUserObj, id int32, hashID *int64) error {
	if hashID != nil {
		_, err := t.client.ChannelsInviteToChannel(&telegram.InputChannelObj{
			ChannelID:  id,
			AccessHash: *hashID,
		}, []telegram.InputUser{user})
		if err != nil {
			t.log.WithError(err).Errorf("failed to invite to channel")
			return err
		}
	} else {
		_, err := t.client.MessagesAddChatUser(id, user, 0)
		if err != nil {
			t.log.WithError(err).Errorf("failed to add chat user")
			return err
		}
	}

	return nil
}
