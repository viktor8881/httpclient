package simplehttp

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type SimpleClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewSimpleClient(baseURL string, timeout time.Duration, transport http.RoundTripper) *SimpleClient {
	return &SimpleClient{
		httpClient: &http.Client{
			Timeout:   timeout,
			Transport: transport,
		},
		baseURL: baseURL,
	}
}

func (c *SimpleClient) Get(ctx context.Context, endpoint string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.httpClient.Do(req)
}

func (c *SimpleClient) Post(ctx context.Context, endpoint string, body interface{}, headers map[string]string) (*http.Response, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.httpClient.Do(req)
}
