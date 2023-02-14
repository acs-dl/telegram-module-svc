package registrator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/acs/gitlab-module/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func RegisterModule(name, endpoint, innerEndpoint, outerEndpoint string) error {
	request := struct {
		Data resources.Module `json:"data"`
	}{
		Data: resources.Module{
			Attributes: resources.ModuleAttributes{
				Name:     name,
				Endpoint: endpoint,
				Link:     innerEndpoint,
			},
		},
	}

	jsonBody, _ := json.Marshal(request)

	req, err := http.NewRequest(http.MethodPost, outerEndpoint, bytes.NewReader(jsonBody))
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

	return nil
}
