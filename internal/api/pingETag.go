package api

import (
	"fmt"
	"net/http"
)

func (c *Client) PingETag(endpoint string, etag string) (isChanged bool, error error) {
	targetURL := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	req, err := http.NewRequest("HEAD", targetURL, nil)
	if err != nil {
		return true, err
	}
	req.Header = http.Header{
		"If-None-Match": {etag},
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return true, error
	}
	if resp.StatusCode == 304 {
		return false, nil
	}
	return true, nil
}
