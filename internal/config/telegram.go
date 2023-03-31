package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type TelegramCfg struct {
	ApiId       int64  `figure:"api_id"`
	ApiHash     string `figure:"api_hash"`
	PhoneNumber string `figure:"phone_number"`
	Host        string `figure:"host"`
}

func validateTelegramCfgValues(cfg TelegramCfg) error {
	return validation.Errors{
		"api_id":       validation.Validate(cfg.ApiId, validation.Required),
		"api_hash":     validation.Validate(cfg.ApiHash, validation.Required),
		"phone_number": validation.Validate(cfg.PhoneNumber, validation.Required),
		"host":         validation.Validate(cfg.Host, validation.Required),
	}.Filter()
}

func (c *config) Telegram() *TelegramCfg {
	return c.telegram.Do(func() interface{} {
		var cfg TelegramCfg
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "telegram")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out telegram params from config"))
		}

		err = validateTelegramCfgValues(cfg)
		if err != nil {
			panic(errors.Wrap(err, "failed to validate telegram params"))
		}

		return &cfg
	}).(*TelegramCfg)
}
