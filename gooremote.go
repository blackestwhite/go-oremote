package gooremote

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func New(AccessToken string) instance {
	return instance{Key: AccessToken}
}

func (i *instance) NewPayment(amount int, next, webhook, description string) (string, error) {
	var pr newPaymentResponse

	p := payment{
		Amount:      amount,
		Next:        next,
		Webhook:     webhook,
		Description: description,
	}

	body, err := i.post("https://api.oremote.org/pay/v1/new", p)
	if err != nil {
		return "", nil
	}

	err = json.NewDecoder(body).Decode(&pr)
	if err != nil {
		return "", err
	}

	if !pr.Ok {
		return "", err
	}

	return pr.Result.ID, nil
}

func GetGatewayURL(id string) string {
	return "https://api.oremote.org/pay/v1/pay/" + id
}

func (i *instance) post(url string, data interface{}) (io.ReadCloser, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"Content-Type":               {"application/json"},
		"x-oremote-api-access-token": {i.Key},
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
