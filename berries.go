package pokego

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (c *Client) Berry(value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New("Berry identifier cannot be empty")
	}
	targetURL := "berry/" + value + "/"
	c.apiClient.Fetch(targetURL)
	return nil
}

func (c *Client) Berries(offset string, limit string) error {
	_, err := strconv.Atoi(offset)
	if err != nil {
		return errors.New("Offset is not a valid number")
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		return errors.New("Limit is not a valid number")
	}
	targetURL := "berry/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	c.apiClient.Fetch(fullURL)
	return nil
}
