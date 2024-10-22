package pokego

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hanle23/pokego/internal/models"
)

func (c *Client) Berry(id string) (result models.Berry, err error) {
	if strings.TrimSpace(id) == "" {
		return result, errors.New("Berry identifier cannot be empty")
	}
	targetURL := "berry"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) Berries(offset string, limit string) (result models.Berries, err error) {
	_, err = strconv.Atoi(offset)
	if err != nil {
		return result, errors.New("Offset is not a valid number")
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		return result, errors.New("Limit is not a valid number")
	}
	targetURL := "berry/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) BerryFirmness(id string) (result models.BerryFirmness, err error) {
	if strings.TrimSpace(id) == "" {
		return result, errors.New("Berry identifier cannot be empty")
	}
	targetURL := "berry-firmness"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) BerryFirmnesses(offset string, limit string) (result models.BerryFirmnesses, err error) {
	_, err = strconv.Atoi(offset)
	if err != nil {
		return result, errors.New("Offset is not a valid number")
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		return result, errors.New("Limit is not a valid number")
	}
	targetURL := "berry-firmness/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) BerryFlavor(id string) (result models.BerryFlavor, err error) {
	if strings.TrimSpace(id) == "" {
		return result, errors.New("Berry identifier cannot be empty")
	}
	targetURL := "berry-flavor"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) BerryFlavors(offset string, limit string) (result models.BerryFlavors, err error) {
	_, err = strconv.Atoi(offset)
	if err != nil {
		return result, errors.New("Offset is not a valid number")
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		return result, errors.New("Limit is not a valid number")
	}
	targetURL := "berry-flavor/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}
