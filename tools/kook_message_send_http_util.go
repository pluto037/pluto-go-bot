package tools

import (
	"bytes"
	"crypto/tls"
	"errors"
	"github.com/pluto037/pluto-go-bot/config/kook"
	"io/ioutil"
	"net/http"
)

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient() *HTTPClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // This allows insecure SSL connections, use it for testing only
	}

	client := &http.Client{
		Transport: transport,
	}

	return &HTTPClient{client}
}

func (c *HTTPClient) SendMessage(body []byte) ([]byte, error) {
	if c.client == nil {
		return nil, errors.New("HTTP client is not initialized")
	}
	url := kook.BASE_URL + kook.CHANNEL_SEND_MESSAGE
	headers := make(map[string]string)
	headers["Authorization"] = kook.TOKEN
	headers["Content-Type"] = "application/json"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// 设置额外的头信息
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
