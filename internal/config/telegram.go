package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type TelegramCfg struct {
	AppId       int64  `figure:"appId"`
	AppHash     string `figure:"appHash"`
	PhoneNumber string `figure:"phoneNumber"`
	Host        string `figure:"host"`
}

func (c *config) Telegram() *TelegramCfg {
	return c.telegram.Do(func() interface{} {
		var config TelegramCfg
		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(c.getter, "telegram")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out telegram params from config"))
		}

		return &config
	}).(*TelegramCfg)
}
