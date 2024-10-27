package pokego

import (
	"fmt"

	"github.com/hanle23/pokego/internal/models"
	"github.com/hanle23/pokego/internal/utils"
)

func (c *Client) Machine(id string) (result models.Machine, err error) {
	err = utils.ArgsValidation(id, true)
	if err != nil {
		return result, err
	}
	targetURL := "machine"
	fullURL := fmt.Sprintf("%s/%s/", targetURL, id)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}

func (c *Client) Machines(offset string, limit string) (result models.Machines, err error) {
	err = utils.PaginationArgsValidation(offset, limit)
	if err != nil {
		return result, err
	}
	targetURL := "machine/"
	fullURL := fmt.Sprintf("%s?offset=%s&limit=%s", targetURL, offset, limit)
	err = c.apiClient.Fetch(fullURL, &result)
	return result, err
}
