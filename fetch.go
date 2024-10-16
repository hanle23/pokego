package pokego

import (
	"fmt"
	"reflect"
)

// Fetch for manual entered endpoint, or passing empty string to get all endpoints
func (c *Client) Fetch(endpoint string, obj interface{}) error {
	if reflect.ValueOf(obj).Kind() != reflect.Ptr {
		return fmt.Errorf("obj must be a pointer")
	}

	err := c.apiClient.Fetch(endpoint, obj)
	if err != nil {
		return fmt.Errorf("error fetching data: %w", err)
	}

	return nil
}
