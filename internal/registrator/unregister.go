package registrator

import (
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (r *registrar) UnregisterModule() error {
	r.logger.Infof("started unregister module `%s`", data.ModuleName)

	req, err := http.NewRequest(http.MethodDelete, r.config.OuterUrl+"/"+data.ModuleName, nil)
	if err != nil {
		return errors.Wrap(err, "couldn't create request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error making http request")
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("error in response, status %s", res.Status))
	}

	r.logger.Infof("finished unregister module `%s`", data.ModuleName)
	return nil
}
