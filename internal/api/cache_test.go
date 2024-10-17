package api_test

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hanle23/pokego/internal/api"
	"github.com/hanle23/pokego/internal/models"
)

var (
	client  *api.Client
	network bytes.Buffer
)

func setup() {
	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}
	baseURL := "https://pokeapi.co/api/v2/"
	var config []func(*api.Config)
	client = api.NewClient(httpClient, baseURL, config)
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
	cache := client.GetUseCache()
	if !cache {
		t.Error("Cache returned nil on enable")
	}
}

func TestSetCache_defaultTime(t *testing.T) {
	enc := gob.NewEncoder(&network)
	testData := models.Berry{Name: "Cheri"}
	err := enc.Encode(testData)
	if err != nil {
		t.Errorf("Encode error: %v", err)
	}
	testPackage := api.CachePackage{Value: network.Bytes(), Etag: "abc123"}
	client.SetCache("berry", testPackage, 0)
	data, expiryTime, found := client.GetCacheInstance().GetWithExpiration("berry")
	if !found {
		t.Errorf("Expected found, got: %v", found)
	}
	expectedTime := time.Now().Add(time.Second * 64908)
	if !expiryTime.Before(expectedTime) {
		t.Errorf("Expected %v before %v", expiryTime, expectedTime)
	}
	for i := range data.(api.CachePackage).Value {
		if data.(api.CachePackage).Value[i] != testPackage.Value[i] {
			t.Errorf("Expected similar %v, got: %v", data.(api.CachePackage).Value[i], testPackage.Value[i])
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
	testPackage := api.CachePackage{Value: network.Bytes(), Etag: "abc123"}
	client.SetCache("berry", testPackage, customTime)
	data, expiryTime, found := client.GetCacheInstance().GetWithExpiration("berry")
	if !found {
		t.Errorf("Expected found, got: %v", found)
	}
	withErrorPoint := customTime + 1
	expectedTime := time.Now().Add(time.Second * time.Duration(withErrorPoint))
	if !expiryTime.Before(expectedTime) {
		t.Errorf("Expected %v before %v", expiryTime, expectedTime)
	}
	for i := range data.(api.CachePackage).Value {
		if data.(api.CachePackage).Value[i] != testPackage.Value[i] {
			t.Errorf("Expected similar %v, got: %v", data.(api.CachePackage).Value[i], testPackage.Value[i])
			break
		}
	}
}

func TestGetCache_invalidEndpoint(t *testing.T) {
	data, expireTime, found := client.GetCache("testing")
	if data != nil {
		t.Errorf("Expected data to be %v, got %v", nil, data)
	}
	if !expireTime.IsZero() {
		t.Errorf("Expected is zero returns %v, got %v", true, expireTime.IsZero())
	}
	if found == true {
		t.Errorf("Expected found to be %v, got %v", false, found)
	}
}

func TestGetCache_validEndpoint(t *testing.T) {
	enc := gob.NewEncoder(&network)
	testData := models.Berry{Name: "Cheri"}
	customTime := 60
	err := enc.Encode(testData)
	if err != nil {
		t.Errorf("Encode error: %v", err)
	}
	testPackage := api.CachePackage{Value: network.Bytes(), Etag: "abc123"}
	client.SetCache("berry", testPackage, customTime)
	data, expireTime, found := client.GetCache("berry")
	if data == nil {
		t.Errorf("Expected data to be %v, got %v", data, nil)
	}
	if expireTime.IsZero() {
		t.Errorf("Expected is zero returns %v, got %v", false, expireTime.IsZero())
	}
	if found != true {
		t.Errorf("Expected found to be %v, got %v", true, found)
	}
}

func TestRetrieve_emptyEndpoint(t *testing.T) {
	result := client.Retrieve("")
	if result != nil {
		t.Errorf("Expected nil, got: %v", result)
	}
}

func TestRetrieve_ValidEndpointNonExistData(t *testing.T) {
	result := client.Retrieve("invalid-key")
	if result != nil {
		t.Errorf("Expected nil, got: %v", result)
	}
}

func TestRetrieve_ValidEndpointExistData(t *testing.T) {
	enc := gob.NewEncoder(&network)
	testData := models.Berry{Name: "Cheri"}
	err := enc.Encode(testData)
	if err != nil {
		t.Errorf("Encode error: %v", err)
	}
	testPackage := api.CachePackage{Value: network.Bytes(), Etag: "abc123"}
	client.SetCache("berry", testPackage, 0)
	result := client.Retrieve("berry")
	if result == nil {
		t.Errorf("Expected data, got nil")
	}
}

func TestRetrieve_ValidEndpointExistDataExpired(t *testing.T) {
	enc := gob.NewEncoder(&network)
	testData := models.Berry{Name: "Cheri"}
	err := enc.Encode(testData)
	if err != nil {
		t.Errorf("Encode error: %v", err)
	}
	testPackage := api.CachePackage{Value: network.Bytes(), Etag: "abc123"}
	client.SetCache("berry", testPackage, 0)
	client.ClearCache()
	result := client.Retrieve("berry")
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
