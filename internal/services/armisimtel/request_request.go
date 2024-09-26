package armisimtel

import (
	"bytes"
	"errors"
	"net/http"
)

func (request *Request) Request(method string, url string, data []byte) (*http.Response, error) {
	token, err := request.Init()
	if err != nil {
		return nil, err
	}

	bearer := "Bearer " + token

	req, errorReq := http.NewRequest(method, BASE_URL+url, bytes.NewBuffer(data))

	if errorReq != nil {
		return nil, errorReq
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	resp, errorResp := client.Do(req)

	if errorResp != nil {
		return nil, errorResp
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return resp, errorResp
}
