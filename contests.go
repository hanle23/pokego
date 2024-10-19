package pokego

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hanle23/pokego/internal/models"
)

func (c *Client) ContestType(id string) (result models.ContestType, err error) {
	if strings.TrimSpace(id) == "" {
		return result, errors.New("ContestType identifier cannot be empty")
	}
	targetURL := "contest-type"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) ContestEffect(id string) (result models.ContestEffect, err error) {
	if strings.TrimSpace(id) == "" {
		return result, errors.New("ContestEffect identifier cannot be empty")
	}
	targetURL := "contest-effect"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) SuperContestEffect(id string) (result models.SuperContestEffect, err error) {
	if strings.TrimSpace(id) == "" {
		return result, errors.New("SuperContestEffect identifier cannot be empty")
	}
	targetURL := "super-contest-effect"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}
