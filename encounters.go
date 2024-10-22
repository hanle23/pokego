package pokego

import (
	"fmt"

	"github.com/hanle23/pokego/internal/models"
	"github.com/hanle23/pokego/internal/utils"
)

func (c *Client) EncounterMethod(id string) (result models.EncounterMethod, err error) {
	err = utils.ArgsValidation(id, false)
	if err != nil {
		return result, err
	}
	targetURL := "encounter-method"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) EncounterMethods(offset string, limit string) (result models.EncounterMethods, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "berry/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}
