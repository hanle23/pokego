package api

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type CachePackage struct {
	value []byte
	etag  string
}

func (c *Client) NewCache() *cache.Cache {
	minTime := c.config.expireTime
	cache := cache.New(minTime, 30*time.Minute)
	return cache
}

func ClearCache(ca cache.Cache) {
	ca.Flush()
}

func (c *Client) SetCache(endpoint string, body CachePackage, duration *int) {
	if c.cache == nil {
		return
	}
	var expireTime time.Duration
	if duration != nil {
		expireTime = time.Duration(*duration) * time.Second
	} else {
		expireTime = c.config.expireTime
	}
	c.cache.Set(endpoint, body, expireTime)
}

func (c *Client) GetCache(endpoint string) (data interface{}, expireTime time.Time, found bool) {
	data, expireTime, found = c.cache.GetWithExpiration(endpoint)
	if !found {
		return nil, expireTime, found
	}
	return data, expireTime, true
}
