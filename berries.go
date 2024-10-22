package pokego

import (
	"fmt"

	"github.com/hanle23/pokego/internal/models"
	"github.com/hanle23/pokego/internal/utils"
)

func (c *Client) Berry(id string) (result models.Berry, err error) {
	err = utils.ArgsValidation(id, false)
	if err != nil {
		return result, err
	}
	targetURL := "berry"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) Berries(offset string, limit string) (result models.Berries, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "berry/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) BerryFirmness(id string) (result models.BerryFirmness, err error) {
	err = utils.ArgsValidation(id, false)
	if err != nil {
		return result, err
	}
	targetURL := "berry-firmness"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) BerryFirmnesses(offset string, limit string) (result models.BerryFirmnesses, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "berry-firmness/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) BerryFlavor(id string) (result models.BerryFlavor, err error) {
	err = utils.ArgsValidation(id, false)
	if err != nil {
		return result, err
	}
	targetURL := "berry-flavor"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) BerryFlavors(offset string, limit string) (result models.BerryFlavors, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "berry-flavor/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}
