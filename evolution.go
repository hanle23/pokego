package pokego

import (
	"fmt"

	"github.com/hanle23/pokego/internal/models"
	"github.com/hanle23/pokego/internal/utils"
)

func (c *Client) EvolutionChain(id string) (result models.EvolutionChain, err error) {
	err = utils.ArgsValidation(id, true)
	if err != nil {
		return result, err
	}
	targetURL := "evolution-chain"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) EvolutionChains(offset string, limit string) (result models.EvolutionChains, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "evolution-chain/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) EvolutionTrigger(id string) (result models.EvolutionTrigger, err error) {
	err = utils.ArgsValidation(id, false)
	if err != nil {
		return result, err
	}
	targetURL := "evolution-trigger"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) EvolutionTriggers(offset string, limit string) (result models.EvolutionTriggers, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "evolution-trigger/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}
