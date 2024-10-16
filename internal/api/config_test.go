package api

import (
	"net/http"
	"testing"
)

func TestCustomConfig(t *testing.T) {
	httpClient := &http.Client{}
	baseURL := "https://api.example.com"
	options := []func(*Config){func(s *Config) {
		s.ChangeUseCache(false)
	}, func(s *Config) {
		s.ChangeExpireTime(1000)
	}}
	client := NewClient(httpClient, baseURL, options)
	if client.GetUseCache() != false {
		t.Errorf("expected %v, got %v", false, client.GetUseCache())
	}
	if client.GetExpireTime() != 1000 {
		t.Errorf("expected %v, got %v", 1000, client.GetExpireTime())
	}
}
