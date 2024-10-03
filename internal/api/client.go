package api

import (
	"net/http"

	"github.com/patrickmn/go-cache"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	config     Config
	cache      *cache.Cache
}

func NewClient(httpClient *http.Client, baseURL string, options []func(*Config)) *Client {
	config := &defaultConfig

	if len(options) > 0 {
		for _, o := range options {
			o(config)
		}
	}

	var newCache *cache.Cache
	client := &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
		config:     *config,
		cache:      newCache,
	}

	if config.useCache == true {
		client.cache = client.NewCache()
	}
	return client
}
