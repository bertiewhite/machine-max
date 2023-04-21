package lorawan

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	baseUrl                = "http://europe-west1-machinemax-dev-d524.cloudfunctions.net"
	registerDevEUIEndpoint = "/sensor-onboarding-sample"

	wrongIdFormatError  = "invalid id format"
	unexpectedHttpError = "unexpected http response status code"
)

type client struct {
	client  *http.Client
	baseUrl string
}

type registerDevEUIRequest struct {
	Deveui string `json:"deveui"`
}

func NewLorawanClient(c *http.Client) LoRaWAN {
	return &client{
		client:  c,
		baseUrl: baseUrl,
	}
}

func (c *client) RegisterDevEUI(id string) (bool, error) {
	if id == "" {
		// TODO: this could be expanded to catch any ids in the wrong format
		return false, errors.New(wrongIdFormatError)
	}

	reqBody := registerDevEUIRequest{
		Deveui: id,
	}

	var buffer bytes.Buffer

	err := json.NewEncoder(&buffer).Encode(reqBody)
	if err != nil {
		return false, err
	}

	request, err := http.NewRequest("POST", c.baseUrl+registerDevEUIEndpoint, &buffer)

	// Probably not necessary but feels like the morally right decision to do
	request.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(request)
	if err != nil {
		return false, err
	}

	if response.StatusCode == 200 {
		return true, nil
	}

	if response.StatusCode == 422 {
		return false, nil
	}

	return false, errors.New(unexpectedHttpError)
}
