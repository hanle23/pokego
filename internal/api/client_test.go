package api

import (
	"net/http"
	"testing"
	"time"
)

func TestDefaultNewClient(t *testing.T) {
	httpClient := &http.Client{}
	baseURL := "https://api.example.com"
	var options []func(*Config)
	client := NewClient(httpClient, baseURL, options)

	if client.httpClient != httpClient {
		t.Errorf("expected %v, got %v", httpClient, client.httpClient)
	}
	if client.baseURL != baseURL {
		t.Errorf("expected %v, got %v", baseURL, client.baseURL)
	}

	if client.GetUseCache() != true {
		t.Errorf("expected %v, got %v", true, client.GetUseCache())
	}

	if client.GetExpireTime() != time.Duration(64908*float64(time.Second)) {
		t.Errorf("expected %v, got %v", time.Duration(64908*float64(time.Second)), client.GetExpireTime())
	}
}
