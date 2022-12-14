package gooremote

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	defer body.Close()

	err = json.NewDecoder(body).Decode(&pr)
	if err != nil {
		return "", err
	}

	if !pr.Ok {
		return "", errors.New("payment could not be created")
	}

	return pr.Result.ID, nil
}

func GetGatewayURL(id string) string {
	return "https://api.oremote.org/pay/v1/pay/" + id
}

func (i *instance) GetRawGatewayURL(id string) (string, error) {
	var res getRawGatewayURLResponse
	url := fmt.Sprintf("https://api.oremote.org/pay/v1/zget/%s", id)

	body, err := i.post(url, nil)
	if err != nil {
		return "", err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&res)
	if err != nil {
		return "", err
	}

	if !res.Ok {
		return "", fmt.Errorf("err code: %d, err desc: %s", res.ErrorCode, res.ErrorDescription)
	}

	return res.Result.URL, nil
}

func (i *instance) Verify(id string) (paid bool, err error) {
	var vr verifyResponse

	url := fmt.Sprintf("https://api.oremote.org/pay/v1/verify/%s", id)

	body, err := i.post(url, nil)
	if err != nil {
		return false, err
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&vr)
	if err != nil {
		return false, err
	}

	if !vr.Ok {
		return false, fmt.Errorf("code: %d, desc: %s", vr.ErrorCode, vr.ErrorDescription)
	}

	return vr.Result.Paid, nil
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
