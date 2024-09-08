package pokego

// Fetch for manual entered endpoint, or passing empty string to get all endpoints
func (c *Client) Fetch(endpoint string) {
	c.apiClient.Fetch(endpoint)
}
