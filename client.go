package pokego

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/hanle23/pokego/internal/api"
)

var lock = &sync.Mutex{}

type Client struct {
	apiClient *api.Client
}

var client *Client

func NewClient() *Client {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		if client == nil {
			httpClient := &http.Client{
				Timeout: time.Second * 30,
			}
			var baseURL string = "https://pokeapi.co/api/v2/"

			apiClient := api.NewClient(httpClient, baseURL)
			client = &Client{
				apiClient: apiClient,
			}
			return client
		} else {
			fmt.Println("Client was already created.")
		}
		fmt.Println("Client was already created.")
	}

	return client
}
