package pokego

import (
	"net/http"
	"sync"
	"time"

	"github.com/hanle23/pokego/internal/api"
)

var (
	client     *Client
	clientOnce sync.Once
	clientLock sync.RWMutex
)

type Client struct {
	apiClient *api.Client
}

func NewClient(options ...func(*api.Config)) *Client {
	clientOnce.Do(func() {
		httpClient := &http.Client{
			Timeout: time.Second * 30,
		}
		baseURL := "https://pokeapi.co/api/v2/"
		apiClient := api.NewClient(httpClient, baseURL, options)
		client = &Client{
			apiClient: apiClient,
		}
	})
	return client
}

func RemoveClient() {
	if client == nil {
		return
	}
	clientLock.Lock()
	defer clientLock.Unlock()
	client = nil
	clientOnce = sync.Once{}
}

func ResetClient() {
	if client == nil {
		NewClient()
	}
	clientLock.Lock()
	defer clientLock.Unlock()
	currentConfig := client.apiClient.GetCurrentConfig()
	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}
	apiClient := api.NewClientWithConfig(httpClient, "https://pokeapi.co/api/v2/", currentConfig)
	client = &Client{
		apiClient: apiClient,
	}
	clientOnce = sync.Once{}
}

func GetClient() *Client {
	if client == nil {
		return nil
	}
	clientLock.RLock()
	defer clientLock.RUnlock()
	return client
}
