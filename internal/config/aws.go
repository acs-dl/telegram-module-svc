package config

import (
	"encoding/json"
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type AwsCfg struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
}

func (c *config) Aws() *AwsCfg {
	return c.aws.Do(func() interface{} {
		var cfg AwsCfg
		value, ok := os.LookupEnv("aws")
		if !ok {
			panic(errors.New("no aws env variable"))
		}

		err := json.Unmarshal([]byte(value), &cfg)
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out aws params from env variable"))
		}

		err = cfg.validate()
		if err != nil {
			panic(errors.Wrap(err, "failed to validate aws config"))
		}

		return &cfg
	}).(*AwsCfg)
}

func (a *AwsCfg) validate() error {
	return validation.Errors{
		"id":     validation.Validate(a.Id, validation.Required),
		"secret": validation.Validate(a.Secret, validation.Required),
	}.Filter()
}
