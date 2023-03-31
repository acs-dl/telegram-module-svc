package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	Endpoint string
}

func (j *Client) AddAPI(api *Service) error {
	url := fmt.Sprintf("%s/apis", j.Endpoint)
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&api); err != nil {
		return errors.Wrap(err, "failed to marshal body")
	}
	_, err := do("POST", url, &body)
	if err != nil {
		return errors.Wrap(err, "failed to add api")
	}
	return nil
}

func (j *Client) UpdateAPI(name string, api *Service) error {
	url := fmt.Sprintf("%s/apis/%s", j.Endpoint, name)
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&api); err != nil {
		return errors.Wrap(err, "failed to marshal body")
	}
	_, err := do("PUT", url, &body)
	if err != nil {
		return errors.Wrap(err, "failed to modify api")
	}
	return nil
}

func (j *Client) GetAPI(name string) (*Service, error) {
	url := fmt.Sprintf("%s/apis/%s", j.Endpoint, name)
	resp, err := do("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get api")
	}

	if resp == nil {
		return nil, nil
	}

	var result Service
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	return &result, nil
}

func do(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer response.Body.Close()
	bodyBB, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return bodyBB, nil
	case http.StatusNotFound, http.StatusNoContent:
		return nil, nil
	case http.StatusBadRequest:
		return nil, E(
			"request was invalid in some way",
			Response(bodyBB),
			Status(response.StatusCode),
		)
	case http.StatusUnauthorized:
		return nil, E(
			"not allowed",
			Response(bodyBB),
			Status(response.StatusCode),
		)
	default:
		return nil, E(
			"something bad happened",
			Response(bodyBB),
			Status(response.StatusCode),
		)
	}
}
