package registrator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (r *registrar) registerModule(_ context.Context) error {
	r.logger.Infof("started register module `%s`", data.ModuleName)

	request := struct {
		Data resources.Module `json:"data"`
	}{
		Data: resources.Module{
			Attributes: resources.ModuleAttributes{
				Name:     data.ModuleName,
				Topic:    r.config.Topic,
				Link:     r.config.InnerUrl,
				Title:    r.config.Title,
				Prefix:   r.config.Prefix,
				IsModule: r.config.IsModule,
			},
		},
	}

	jsonBody, _ := json.Marshal(request)

	req, err := http.NewRequest(http.MethodPost, r.config.OuterUrl, bytes.NewReader(jsonBody))
	if err != nil {
		return errors.Wrap(err, "couldn't create request")
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error making http request")
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("error in response, status %s", res.Status))
	}

	r.logger.Infof("finished register module `%s`", data.ModuleName)
	return nil
}
