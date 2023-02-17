package tg

import (
	"fmt"
	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"time"
)

func (t *tg) DeleteFromChatFromApi(username, phone *string, title string) error {
	err := t.kickUserFlow(title, username, phone)
	if err != nil {
		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			t.log.WithError(err).Errorf("failed to kick user, some strange error")
			return errors.Wrap(err, "failed to kick user, some strange error")
		}
		if errResponse.Message == "FLOOD_WAIT_X" {
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			t.log.Warnf("we need to wait `%s`", timeoutDuration.String())
			time.Sleep(timeoutDuration)
			return t.DeleteFromChatFromApi(username, phone, title)
		}

		t.log.WithError(err).Errorf("failed to kick user from %s", title)
		return errors.Wrap(err, fmt.Sprintf("failed to kick user from %s", title))
	}

	t.log.Infof("successfully kicked user from %s", title)
	return nil
}

func (t *tg) kickUserFlow(title string, username, phone *string) error {
	id, accessHash, err := t.findChatByTitle(title)
	if err != nil {
		t.log.WithError(err).Errorf("failed to find chat %s", title)
		return err
	}

	var inputUser *telegram.InputUserObj
	if username != nil {
		inputUser, err = t.getUserByUsername(*username)
	}
	if phone != nil {
		inputUser, err = t.getUserByPhone(*phone)
	}
	if err != nil {
		t.log.WithError(err).Errorf("failed to get user")
		return err
	}

	user, err := t.checkUserInChat(*id, accessHash, int64(inputUser.UserID))
	if err != nil {
		t.log.WithError(err).Errorf("failed to check chat and user")
		return err
	}
	if user == nil {
		t.log.Errorf("user is not in `%s`", title)
		return errors.Errorf("user is not in `%s`", title)
	}
	if err = t.kickUser(inputUser, *id, accessHash); err != nil {
		t.log.WithError(err).Errorf("failed to kick user")
		return err
	}

	return nil
}

func (t *tg) kickUser(user *telegram.InputUserObj, id int32, hashID *int64) error {
	if hashID != nil {
		_, err := t.client.ChannelsEditBanned(&telegram.InputChannelObj{
			ChannelID:  id,
			AccessHash: *hashID,
		}, user, &telegram.ChatBannedRights{
			ViewMessages: true,
			SendMessages: true,
			SendMedia:    true,
			SendStickers: true,
			SendGifs:     true,
			SendGames:    true,
			SendInline:   true,
			EmbedLinks:   true,
			SendPolls:    true,
			ChangeInfo:   true,
			InviteUsers:  true,
			PinMessages:  true,
			UntilDate:    0,
		})
		if err != nil {
			t.log.WithError(err).Errorf("failed to ban channel user")
			return err
		}
	} else {
		_, err := t.client.MessagesDeleteChatUser(id, user)
		if err != nil {
			t.log.WithError(err).Errorf("failed to delete chat user")
			return err
		}
	}

	return nil
}
