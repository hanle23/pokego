package pokego

import (
	"fmt"

	"github.com/hanle23/pokego/internal/models"
	"github.com/hanle23/pokego/internal/utils"
)

func (c *Client) Generation(id string) (result models.Generation, err error) {
	err = utils.ArgsValidation(id, false)
	if err != nil {
		return result, err
	}
	targetURL := "generation"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) Generations(offset string, limit string) (result models.Generations, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "generation/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) Pokedex(id string) (result models.Pokedex, err error) {
	err = utils.ArgsValidation(id, false)
	if err != nil {
		return result, err
	}
	targetURL := "pokedex"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) Pokedexes(offset string, limit string) (result models.Pokedexes, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "pokedex/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) Version(id string) (result models.Version, err error) {
	err = utils.ArgsValidation(id, false)
	if err != nil {
		return result, err
	}
	targetURL := "version"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) Versions(offset string, limit string) (result models.Versions, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "version/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) VersionGroup(id string) (result models.VersionGroup, err error) {
	err = utils.ArgsValidation(id, false)
	if err != nil {
		return result, err
	}
	targetURL := "version-group"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) VersionGroups(offset string, limit string) (result models.VersionGroups, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "version-group/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}
