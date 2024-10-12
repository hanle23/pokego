package api

import (
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	client *Client
)

func setup() {
	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}
	baseURL := "https://pokeapi.co/api/v2/"
	var config []func(*Config)
	client = NewClient(httpClient, baseURL, config)
}

func shutdown() {
	client = nil
}

func TestClient(t *testing.T) {
	if client == nil {
		t.Error("NewClient() returned nil")
	}
}

func TestCacheEnabled(t *testing.T) {
	cache := client.cache
	if cache == nil {
		t.Error("Cache returned nil on enable")
	}
}

func TestEmptyEndpoint(t *testing.T) {
	result := client.Retrieve("")
	if result != nil {
		t.Errorf("Expected nil, got: %v", result)
	}
}

func TestValidEndpointNonExistData(t *testing.T) {
	result := client.Retrieve("invalid-key")
	if result != nil {
		t.Errorf("Expected nil, got: %v", result)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
