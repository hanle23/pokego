package api

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type CachePackage struct {
	Value []byte
	Etag  string
}

func (c *Client) NewCache() *cache.Cache {
	minTime := c.config.expireTime
	cache := cache.New(minTime, 30*time.Minute)
	return cache
}

func (c *Client) ClearCache() {
	c.cache.Flush()
}

func (c *Client) SetCache(endpoint string, body CachePackage, duration int) {
	if c.cache == nil {
		return
	}
	var expireTime time.Duration
	if duration != 0 {
		expireTime = time.Duration(duration) * time.Second
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

func (c *Client) Retrieve(endpoint string) (data interface{}) {
	if endpoint == "" {
		return nil
	}
	data, expireTime, found := c.GetCache(endpoint)
	if !found || data == nil {
		return nil
	}
	if time.Now().Before(expireTime) {
		return data
	}
	etag := data.(CachePackage).Etag
	isChanged, err := c.PingETag(endpoint, etag)
	if err != nil {
		return nil
	}
	if isChanged {
		return nil
	}
	return data
}

func (c *Client) GetCacheInstance() cache.Cache {
	cache := c.cache
	return *cache
}
