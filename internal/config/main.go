package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	// base
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer

	// connectors

	// other config values
	Telegram() *TelegramCfg
	Amqp() *AmqpData
	Cache() *CacheData
	JwtParams() *JwtCfg

	// Registrator config
	Registrator() RegistratorConfig
}

type config struct {
	// base
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	getter kv.Getter
	comfig.Listenerer

	// connectors

	// other config values
	telegram    comfig.Once
	amqp        comfig.Once
	cache       comfig.Once
	registrator comfig.Once
	jwtCfg      comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:     getter,
		Databaser:  pgdb.NewDatabaser(getter),
		Copuser:    copus.NewCopuser(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Listenerer: comfig.NewListenerer(getter),
	}
}
