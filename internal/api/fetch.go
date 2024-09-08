package api

import (
	"fmt"
)

func (c *Client) Fetch(endpoint string) {
	targetURL := c.baseURL + endpoint
	fmt.Println("Fetch", targetURL)
}
