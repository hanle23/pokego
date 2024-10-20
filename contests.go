package pokego

import (
	"errors"
	"fmt"
	"strconv"
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

func (c *Client) ContestTypes(offset string, limit string) (result models.ContestTypes, err error) {
	_, err = strconv.Atoi(offset)
	if err != nil {
		return result, errors.New("Offset is not a valid number")
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		return result, errors.New("Limit is not a valid number")
	}
	targetURL := "contest-type"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) ContestEffect(id string) (result models.ContestEffect, err error) {
	if strings.TrimSpace(id) == "" {
		return result, errors.New("ContestEffect identifier cannot be empty")
	}
	if _, err := strconv.Atoi(strings.TrimSpace(id)); err != nil {
		return result, errors.New("ContestEffect identifier must be a number")
	}
	targetURL := "contest-effect"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) ContestEffects(offset string, limit string) (result models.ContestEffects, err error) {
	_, err = strconv.Atoi(offset)
	if err != nil {
		return result, errors.New("Offset is not a valid number")
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		return result, errors.New("Limit is not a valid number")
	}
	targetURL := "contest-effect"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) SuperContestEffect(id string) (result models.SuperContestEffect, err error) {
	if strings.TrimSpace(id) == "" {
		return result, errors.New("SuperContestEffect identifier cannot be empty")
	}
	if _, err := strconv.Atoi(strings.TrimSpace(id)); err != nil {
		return result, errors.New("SuperContestEffect identifier must be a number")
	}
	targetURL := "super-contest-effect"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) SuperContestEffects(offset string, limit string) (result models.SuperContestEffects, err error) {
	_, err = strconv.Atoi(offset)
	if err != nil {
		return result, errors.New("Offset is not a valid number")
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		return result, errors.New("Limit is not a valid number")
	}
	targetURL := "super-contest-effect"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}
