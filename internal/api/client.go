package api

import (
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(httpClient *http.Client, baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}
