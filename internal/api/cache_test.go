package api

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hanle23/pokego/internal/models"
)

var (
	client  *Client
	network bytes.Buffer
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

func TestSetCache_defaultTime(t *testing.T) {
	enc := gob.NewEncoder(&network)
	testData := models.Berry{Name: "Cheri"}
	err := enc.Encode(testData)
	if err != nil {
		t.Errorf("Encode error: %v", err)
	}
	testPackage := CachePackage{value: network.Bytes(), etag: "abc123"}
	client.SetCache("berry", testPackage, 0)
	data, expiryTime, found := client.cache.GetWithExpiration("berry")
	if !found {
		t.Errorf("Expected found, got: %v", found)
	}
	expectedTime := time.Now().Add(time.Second * 64908)
	if !expiryTime.Before(expectedTime) {
		t.Errorf("Expected %v before %v", expiryTime, expectedTime)
	}
	for i := range data.(CachePackage).value {
		if data.(CachePackage).value[i] != testPackage.value[i] {
			t.Errorf("Expected similar %v, got: %v", data.(CachePackage).value[i], testPackage.value[i])
			break
		}
	}
}

func TestSetCache_customTime(t *testing.T) {
	enc := gob.NewEncoder(&network)
	testData := models.Berry{Name: "Cheri"}
	customTime := 60
	err := enc.Encode(testData)
	if err != nil {
		t.Errorf("Encode error: %v", err)
	}
	testPackage := CachePackage{value: network.Bytes(), etag: "abc123"}
	client.SetCache("berry", testPackage, customTime)
	data, expiryTime, found := client.cache.GetWithExpiration("berry")
	if !found {
		t.Errorf("Expected found, got: %v", found)
	}
	withErrorPoint := customTime + 1
	expectedTime := time.Now().Add(time.Second * time.Duration(withErrorPoint))
	if !expiryTime.Before(expectedTime) {
		t.Errorf("Expected %v before %v", expiryTime, expectedTime)
	}
	for i := range data.(CachePackage).value {
		if data.(CachePackage).value[i] != testPackage.value[i] {
			t.Errorf("Expected similar %v, got: %v", data.(CachePackage).value[i], testPackage.value[i])
			break
		}
	}
}

func TestCacheHitValidData(t *testing.T) {
	// testData := models.Berry{Name:"Cheri"}
	// client.SetCache("berry", testData, time.Now().Add(1*time.Hour))
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
