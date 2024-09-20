package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Fetch(endpoint string, obj interface{}) error {
	targetURL := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	resp, err := c.httpClient.Get(targetURL)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(obj)
	if err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	return nil
}
