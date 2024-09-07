package api

import (
	"fmt"
)

func (c *Client) BerryByName(name string) {
	fmt.Println(name)
}

func (c *Client) BerryById(id int) {
	fmt.Println(id)
}
