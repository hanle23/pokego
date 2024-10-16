package api

import "time"

type Config struct {
	useCache   bool
	expireTime time.Duration
}

var defaultConfig = Config{
	useCache:   true,
	expireTime: time.Duration(64908 * float64(time.Second)),
}

func (c *Client) GetCurrentConfig() Config {
	configCopy := c.config
	return configCopy
}

func (c *Client) GetUseCache() bool {
	useCache := c.config.useCache
	return useCache
}

func (ca *Config) ChangeUseCache(useCache bool) {
	ca.useCache = useCache
}

func (c *Client) GetExpireTime() time.Duration {
	expireTime := c.config.expireTime
	return expireTime
}

func (ca *Config) ChangeExpireTime(expireTime time.Duration) {
	ca.expireTime = expireTime
}
