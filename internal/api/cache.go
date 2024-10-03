package api

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func (c *Client) NewCache() *cache.Cache {
	minTime := c.config.expireTime
	cache := cache.New(minTime*time.Minute, 30*time.Minute)
	return cache
}

func ClearCache(ca cache.Cache) {
	ca.Flush()
}

func (c *Client) SetCache(endpoint string, body []byte, duration *int) {
	if c.cache == nil {
		return
	}
	var expireTime time.Duration
	if duration != nil {
		expireTime = time.Duration(*duration) * time.Minute
	} else {
		expireTime = c.config.expireTime
	}
	c.cache.Set(endpoint, body, expireTime)
}
