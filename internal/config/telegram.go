package config

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	knox "gitlab.com/distributed_lab/knox/knox-fork/client/external_kms"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type TelegramCfg struct {
	SuperUser TgData `json:"super_user"`
	User      TgData `json:"user"`
}

type TgData struct {
	ApiId       int64  `json:"api_id"`
	ApiHash     string `json:"api_hash"`
	PhoneNumber string `json:"phone_number"`
}

func (c *config) Telegram() *TelegramCfg {
	return c.telegram.Do(func() interface{} {
		var cfg TelegramCfg

		client := knox.NewKeyManagementClient(c.getter)

		key, err := client.GetKey("super_user", "1436751686134996000")
		if err != nil {
			panic(errors.Wrap(err, "failed to get super user key"))
		}

		var usr TgData
		err = json.Unmarshal(key, &usr)
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out super user params from vault"))
		}
		cfg.SuperUser = usr

		key, err = client.GetKey("user", "4999510296215657000")
		if err != nil {
			panic(errors.Wrap(err, "failed to get user key"))
		}

		err = json.Unmarshal(key, &usr)
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out user params from vault"))
		}
		cfg.User = usr

		err = cfg.validate()
		if err != nil {
			panic(errors.Wrap(err, "failed to validate telegram config"))
		}

		return &cfg
	}).(*TelegramCfg)
}

func (tg *TelegramCfg) validate() error {
	return validation.Errors{
		"super_user_api_id":   validation.Validate(tg.SuperUser.ApiId, validation.Required),
		"super_user_api_hash": validation.Validate(tg.SuperUser.ApiHash, validation.Required),
		"super_user_phone":    validation.Validate(tg.SuperUser.PhoneNumber, validation.Required),
		"user_api_id":         validation.Validate(tg.User.ApiId, validation.Required),
		"user_api_hash":       validation.Validate(tg.User.ApiHash, validation.Required),
		"user_phone":          validation.Validate(tg.User.PhoneNumber, validation.Required),
	}.Filter()
}
