package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Fetch(endpoint string, obj interface{}) error {
	targetURL := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	data := c.Retrieve(endpoint)
	if data != nil {
		// Handling unMarshal and return obj
	}
	resp, err := c.httpClient.Get(targetURL)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	fmt.Println(resp.Header.Get("Etag"))
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}
	cachePackage := CachePackage{value: bodyBytes, etag: resp.Header.Get("Etag")}
	c.SetCache(endpoint, cachePackage, nil)
	return json.Unmarshal(bodyBytes, obj)
}
