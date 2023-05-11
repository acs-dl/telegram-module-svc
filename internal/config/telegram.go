package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type TelegramCfg struct {
	SuperUser TgData `figure:"super_user"`
	User      TgData `figure:"user"`
}

type TgData struct {
	ApiId       int64  `figure:"api_id"`
	ApiHash     string `figure:"api_hash"`
	PhoneNumber string `figure:"phone_number"`
}

func (c *config) Telegram() *TelegramCfg {
	return c.telegram.Do(func() interface{} {
		var cfg TelegramCfg
		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "telegram")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out telegram params from config"))
		}

		return &cfg
	}).(*TelegramCfg)
}
