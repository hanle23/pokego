package api

import (
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	config     Config
}

func NewClient(httpClient *http.Client, baseURL string, options []func(*Config)) *Client {
	config := &defaultConfig

	if len(options) > 0 {
		for _, o := range options {
			o(config)
		}
	}
	client := &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
		config:     *config,
	}
	return client
}
