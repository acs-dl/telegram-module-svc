package cop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Client struct {
	Endpoint string
}

type Service struct {
	Data ServiceData `json:"data"`
}

type ServiceData struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Attributes ServiceAttributes `json:"attributes"`
}

type ServiceAttributes struct {
	Name string `json:"name"`
	Port string `json:"port"`
	Rule string `json:"rule"`
	Url  string `json:"url"`
}

func (c *Client) AddService(service Service) error {
	request, err := json.Marshal(service)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	url := fmt.Sprintf("%s/cop/services", c.Endpoint)
	resp, err := http.DefaultClient.Post(url, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	if resp.StatusCode > 300 {
		return errors.From(errors.New("failed to add service"),
			logan.F{"status": resp.StatusCode,
				"body": string(body),
			})
	}

	return nil
}
