package api

import (
	"fmt"
)

func (c *Client) Fetch(endpoint string, obj interface{}) error {
	targetURL := c.baseURL + endpoint
	fmt.Println("Fetch", targetURL)
	return nil
}
