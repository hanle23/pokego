package pokego

import (
	"strconv"
)

func (c *Client) Berry(value string) {
	identifier, err := strconv.Atoi(value)
	if err == nil {
		c.apiClient.BerryById(identifier)
	} else {
		c.apiClient.BerryByName(value)
	}
}
