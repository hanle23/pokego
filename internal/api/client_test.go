package api

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	httpClient := &http.Client{}
	baseURL := "https://api.example.com"
	client := NewClient(httpClient, baseURL)

	if client.httpClient != httpClient {
		t.Errorf("expected %v, got %v", httpClient, client.httpClient)
	}
	if client.baseURL != baseURL {
		t.Errorf("expected %v, got %v", baseURL, client.baseURL)
	}
}
