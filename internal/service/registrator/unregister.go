package registrator

import (
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

func UnregisterModule(name, outerEndpoint string) error {
	req, err := http.NewRequest(http.MethodDelete, outerEndpoint+"/"+name, nil)
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

	return nil
}
