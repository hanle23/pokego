package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Fetch(endpoint string, obj interface{}) error {
	targetURL := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	if c.config.useCache {
		data := c.Retrieve(endpoint)
		if data != nil {
			return json.Unmarshal(data.(CachePackage).Value, obj)
		}
	}
	resp, err := c.httpClient.Get(targetURL)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}
	if c.config.useCache {
		cachePackage := CachePackage{Value: bodyBytes, Etag: resp.Header.Get("Etag")}
		c.SetCache(endpoint, cachePackage, 0)
	}
	err = json.Unmarshal(bodyBytes, obj)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
