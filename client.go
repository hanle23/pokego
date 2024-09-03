package pokego

import (
	"net/http"
	"time"

	"github.com/hanle23/pokego/internal/api"
)

type Client struct {
	apiClient *api.Client
}

func NewClient(baseURL string) *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}

	apiClient := api.NewClient(httpClient, baseURL)

	return &Client{
		apiClient: apiClient,
	}
}
