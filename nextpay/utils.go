package nextpay

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Map map[string]string

func formURLEncodedPayload(data Map) io.Reader {
	requestPayloads := url.Values{}
	for key, value := range data {
		requestPayloads[key] = []string{value}
	}
	return bytes.NewReader([]byte(requestPayloads.Encode()))
}

func post[T any](endpoint string, data Map) (*T, error) {
	resp, err := http.Post(BaseURL, "application/x-www-form-urlencoded", formURLEncodedPayload(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData T
	return &respData, json.NewDecoder(resp.Body).Decode(respData)
}
