package pokego

import (
	"net/http"
	"sync"
	"time"

	"github.com/hanle23/pokego/internal/api"
)

var (
	instance *Client
	once     sync.Once
	mu       sync.RWMutex
)

type Client struct {
	apiClient  *api.Client
	baseURL    string
	httpClient *http.Client
}

func DefaultConfig() (*http.Client, string) {
	return &http.Client{
		Timeout: time.Second * 30,
	}, "https://pokeapi.co/api/v2/"
}

func NewClient(options ...func(*api.Config)) *Client {
	once.Do(func() {
		httpClient, baseURL := DefaultConfig()
		apiClient := api.NewClient(httpClient, baseURL, options)

		instance = &Client{
			apiClient:  apiClient,
			httpClient: httpClient,
			baseURL:    baseURL,
		}
	})

	return instance
}

func (c *Client) Reset() {
	if c == nil {
		return
	}

	mu.Lock()
	defer mu.Unlock()

	currentConfig := c.apiClient.GetCurrentConfig()
	c.apiClient = api.NewClientWithConfig(c.httpClient, c.baseURL, currentConfig)
	once = sync.Once{}
}

func (c *Client) Close() {
	if c == nil {
		return
	}

	mu.Lock()
	defer mu.Unlock()

	instance = nil
	c.apiClient = nil
	once = sync.Once{}
}

func (c *Client) GetClient() *Client {
	if c == nil {
		return nil
	}

	mu.RLock()
	defer mu.RUnlock()

	return instance
}
