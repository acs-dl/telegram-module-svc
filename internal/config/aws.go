package config

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	knox "gitlab.com/distributed_lab/knox/knox-fork/client/external_kms"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type AwsCfg struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
}

func (c *config) Aws() *AwsCfg {
	return c.aws.Do(func() interface{} {

		var cfg AwsCfg
		client := knox.NewKeyManagementClient(c.getter)

		key, err := client.GetKey("aws", "4724180100528296000")
		if err != nil {
			panic(errors.Wrap(err, "failed to get aws key"))
		}

		err = json.Unmarshal(key, &cfg)
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
