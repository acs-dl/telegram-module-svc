package config

import (
	"fmt"
	"os"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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
		cfg := lookupConfigEnv()

		err := cfg.validate()
		if err != nil {
			panic(errors.Wrap(err, "failed to validate telegram config"))
		}

		return cfg
	}).(*TelegramCfg)
}

func lookupConfigEnv() *TelegramCfg {
	superUser := lookupTgData("super_user")
	user := lookupTgData("user")

	return &TelegramCfg{
		superUser,
		user,
	}
}

func lookupTgData(user string) TgData {
	apiIdEnv := user + "_api_id"
	apiHashEnv := user + "_api_hash"
	phoneEnv := user + "_phone"

	apiIdStr, ok := os.LookupEnv(apiIdEnv)
	if !ok {
		panic(errors.New(fmt.Sprintf("no %s env variable", apiIdEnv)))
	}

	apiId, err := strconv.ParseInt(apiIdStr, 10, 64)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("invalid %s env variable (must be integer)", apiIdEnv)))
	}

	apiHash, ok := os.LookupEnv(apiHashEnv)
	if !ok {
		panic(errors.New(fmt.Sprintf("no %s env variable", apiHashEnv)))
	}

	phone, ok := os.LookupEnv(phoneEnv)
	if !ok {
		panic(errors.New(fmt.Sprintf("no %s env variable", phoneEnv)))
	}

	return TgData{
		apiId,
		apiHash,
		phone,
	}
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
