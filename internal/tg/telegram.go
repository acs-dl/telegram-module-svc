package tg

import (
	"os"
	"path/filepath"
	"time"

	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type TelegramClient interface {
	GetUsersFromApi(title string) ([]data.User, error)
	GetUserFromApi(username, phone *string) (*data.User, error)
	GetChatUserFromApi(username, phone *string, title string) (*data.User, error)
	SearchByFromApi(username, phone *string, amount int64) ([]data.User, error)

	GetChatFromApi(title string) (*Chat, error)

	AddUserInChatFromApi(username, phone *string, title string) error
	UpdateUserInChatFromApi(username, phone *string, title string) error
	DeleteFromChatFromApi(username, phone *string, title string) error

	GetClient() *telegram.Client
}

type tg struct {
	client *telegram.Client
	log    *logan.Entry
	tgCfg  *config.TelegramCfg
}

type Chat struct {
	id         *int32
	accessHash *int64
}

func NewTg(cfg *config.TelegramCfg, log *logan.Entry) TelegramClient {
	currentDir, err := os.Getwd()
	if err != nil {
		log.WithError(err).Errorf("failed to get current directory path")
		panic(errors.Wrap(err, "failed to get current directory path"))
	}

	sessionFile := filepath.Join(currentDir, "session.json")
	publicKeys := filepath.Join(currentDir, "tg_public_keys.pem")

	client, err := telegram.NewClient(telegram.ClientConfig{
		SessionFile:    sessionFile,
		ServerHost:     cfg.Host,
		PublicKeysFile: publicKeys,
		AppID:          int(cfg.ApiId),
		AppHash:        cfg.ApiHash,
	})
	if err != nil {
		log.WithError(err).Errorf("failed to create client")
		panic(errors.Wrap(err, "failed to create client"))
	}

	// Please, don't spam auth too often, if you have session file, don't repeat auth process.
	signedIn, err := IsSessionRegistered(client)
	if err != nil {
		log.WithError(err).Errorf("failed to check that session is registered")
		panic(errors.Wrap(err, "failed to check that session is registered"))
	}

	if signedIn {
		println("You've already signed in!")
		return &tg{
			client: client,
			log:    log,
			tgCfg:  cfg,
		}
	}

	setCode, err := client.AuthSendCode(
		cfg.PhoneNumber, int32(cfg.ApiId), cfg.ApiHash, &telegram.CodeSettings{},
	)

	if err != nil {
		errResponse := &mtproto.ErrResponseCode{}
		if !pkgErrors.As(err, &errResponse) {
			log.WithError(err).Errorf("got strange error")
			panic(errors.Wrap(err, "got strange error"))
		}

		switch errResponse.Message {
		case "AUTH_RESTART":
			log.Warnf("You accidentally restart authorization process!\n" +
				"You should login only once, if you'll spam 'AuthSendCode' method, you can be\n" +
				"timeouted to long long time. You warned.")
		case "FLOOD_WAIT_X":
			timeoutDuration := time.Second * time.Duration(errResponse.AdditionalInfo.(int))
			log.Warnf("you've reached flood timeout\n" +
				"Repeat after " + timeoutDuration.String())
		default:
			log.Errorf("Unknown error occurred: %s", errResponse.Error())
		}

		panic(errors.Wrap(err, "failed to send auth code"))
	}

	code := enter("Auth code")

	_, err = client.AuthSignIn(
		cfg.PhoneNumber,
		setCode.PhoneCodeHash,
		code,
	)
	if err == nil {
		isCreated, err := checkIfActiveSessionCreated(client, log)
		if err != nil {
			log.WithError(err).Errorf("failed to check if active session created")
			panic(errors.Wrap(err, "failed to check if active session created"))
		}
		if !isCreated {
			log.Errorf("active session wasn't created")
			panic(errors.Errorf("active session wasn't created"))
		}

		log.Infof("Success! You've signed in!")
		return &tg{
			client: client,
			log:    log,
			tgCfg:  cfg,
		}
	}

	//in case of 2FA we need to do some more steps

	errResponse := &mtproto.ErrResponseCode{}
	ok := pkgErrors.As(err, &errResponse)
	if !ok || errResponse.Message != "SESSION_PASSWORD_NEEDED" {
		log.WithError(err).Errorf("sign in process failed")
		panic(errors.Wrap(err, "sign in process failed"))
	}

	password := enter("Password")

	accountPassword, err := client.AccountGetPassword()
	if err != nil {
		log.WithError(err).Errorf("failed to get password")
		panic(errors.Wrap(err, "failed to get password"))
	}

	inputCheck, err := telegram.GetInputCheckPassword(password, accountPassword)
	if err != nil {
		log.WithError(err).Errorf("failed to create password object")
		panic(errors.Wrap(err, "failed to create password object"))
	}

	_, err = client.AuthCheckPassword(inputCheck)
	if err != nil {
		log.WithError(err).Errorf("failed to check password")
		panic(errors.Wrap(err, "failed to check password"))
	}

	log.Infof("Success! You've signed in!")

	isCreated, err := checkIfActiveSessionCreated(client, log)
	if err != nil {
		log.WithError(err).Errorf("failed to check if active session created")
		panic(errors.Wrap(err, "failed to check if active session created"))
	}
	if !isCreated {
		log.Errorf("active session wasn't created")
		panic(errors.Errorf("active session wasn't created"))
	}

	return &tg{
		client: client,
		log:    log,
		tgCfg:  cfg,
	}
}

func checkIfActiveSessionCreated(client *telegram.Client, log *logan.Entry) (bool, error) {
	for i := 0; i < 5; i++ {
		isRegistered, err := IsSessionRegistered(client)
		if err != nil {
			return false, err
		}
		if isRegistered {
			return isRegistered, nil
		}
		log.Warnf("active session wasn't created, wait a minute")
		time.Sleep(time.Minute)
	}

	return false, nil
}
