package tg_client

import (
	"context"

	"github.com/gotd/contrib/bg"
	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func getSessionFromFile(sessionFileName string) (*session.StorageMemory, error) {
	sessionFile := telegram.FileSessionStorage{
		Path: sessionFileName,
	}

	data, err := sessionFile.LoadSession(context.Background())
	if err != nil {
		return nil, err
	}

	var storage = new(session.StorageMemory)

	err = storage.StoreSession(context.Background(), data)
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func authFlow(client *telegram.Client, storage *session.StorageMemory, ctx context.Context, phone, sessionFileName string) error {
	sendCodeClass, err := client.Auth().SendCode(ctx, phone, auth.SendCodeOptions{})
	if err != nil {
		return err
	}

	code, ok := sendCodeClass.(*tg.AuthSentCode)
	if !ok {
		return errors.New("wrong sent code type")
	}

	enteredCode := enter("code")

	_, err = client.Auth().SignIn(ctx, phone, enteredCode, code.PhoneCodeHash)
	if err == nil {
		err = storage.WriteFile(sessionFileName, 0644)
		if err != nil {
			return err
		}

		return nil
	}

	if err != auth.ErrPasswordAuthNeeded {
		return errors.Wrap(err, "unexpected error")
	}

	return handle2FAFlow(client, storage, ctx, sessionFileName)
}

func handle2FAFlow(client *telegram.Client, storage *session.StorageMemory, ctx context.Context, sessionFileName string) error {
	password := enter("Password")

	_, err := client.Auth().Password(ctx, password)
	if err != nil {
		return err
	}

	err = storage.WriteFile(sessionFileName, 0644)
	if err != nil {
		return err
	}

	return nil
}

func login(ctx context.Context, data config.TgData, sessionFileName string) (*telegram.Client, error) {
	storage, err := getSessionFromFile(sessionFileName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get session file")
	}

	client := telegram.NewClient(int(data.ApiId), data.ApiHash, telegram.Options{
		SessionStorage: storage,
	})

	// bg.Connect will call Run in background.
	// Call stop() to disconnect.
	_, err = bg.Connect(client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to run telegram client in background")
	}

	status, err := client.Auth().Status(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get auth status")
	}

	if status.Authorized {
		return client, nil
	}

	err = authFlow(client, storage, ctx, data.PhoneNumber, sessionFileName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to authenticate")
	}

	return client, nil
}
