package pokego

import (
	"time"

	"github.com/hanle23/pokego/internal/api"
)

func (c *Client) GetCurrentConfig() *api.Config {
	configCopy := c.apiClient.GetCurrentConfig()
	return &configCopy
}

func WithUseCache(useCache bool) func(*api.Config) {
	return func(s *api.Config) {
		s.ChangeUseCache(useCache)
	}
}

func WithExpireTime(expireTime time.Duration) func(*api.Config) {
	return func(s *api.Config) {
		s.ChangeExpireTime(expireTime)
	}
}
