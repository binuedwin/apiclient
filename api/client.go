package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"thunes-client/pkg"
)

type ThunesClient struct {
	baseUrl   string
	apiKey    string
	apiSecret string
}

func NewThunesClient(baseUrl, apiKey, apiSecret string) *ThunesClient {
	return &ThunesClient{
		baseUrl:   baseUrl,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (tc *ThunesClient) NewRequest(method, url string, body io.Reader, expectedStatus int) ([]byte, error) {
	// construct the request
	req, err := http.NewRequest(http.MethodGet, tc.baseUrl+url, body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// set the auth
	req.SetBasicAuth(tc.apiKey, tc.apiSecret)

	// make the http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// handle the error response case
	if resp.StatusCode != expectedStatus {
		return nil, respError(resp.Body)
	}

	// return the bytes from the response if present
	return ioutil.ReadAll(resp.Body)
}

func respError(body io.Reader) error {
	var errors pkg.Errors
	if err := json.NewDecoder(body).Decode(&errors); err != nil {
		return err
	}

	return errors.ToError()
}
