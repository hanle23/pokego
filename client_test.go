package pokego

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
}

func TestClientSingleton(t *testing.T) {
	client := NewClient()
	client2 := NewClient()
	if client.apiClient != client2.apiClient {
		t.Error("NewClient did not return a singleton")
	}
}

func TestRemoveClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("NewClient() returned nil")
	}

	initialAPIClient := client.apiClient
	if initialAPIClient == nil {
		t.Fatal("Initial apiClient is nil")
	}

	client.Close()

	currentClient := client.GetClient()
	if currentClient != nil {
		t.Error("GetClient() should return nil after Close()")
	}

	if client.apiClient != nil {
		t.Error("client.apiClient should be nil after Close()")
	}
}

func TestResetClient(t *testing.T) {
	client := NewClient(WithExpireTime(50), WithUseCache(false))
	if client == nil {
		t.Fatal("NewClient() returned nil")
	}

	initialAPIClient := client.apiClient
	if initialAPIClient == nil {
		t.Fatal("Initial apiClient is nil")
	}

	initialUseCache := initialAPIClient.GetUseCache()
	initialExpireTime := initialAPIClient.GetExpireTime()

	client.Reset()

	if client == nil {
		t.Fatal("Client became nil after Reset()")
	}

	newAPIClient := client.apiClient
	if newAPIClient == nil {
		t.Fatal("New apiClient is nil after Reset()")
	}

	if initialAPIClient == newAPIClient {
		t.Error("Reset() did not create new apiClient instance")
	}

	if newAPIClient.GetUseCache() != initialUseCache {
		t.Errorf("Reset() did not maintain useCache setting. Expected: %v, Got: %v",
			initialUseCache, newAPIClient.GetUseCache())
	}

	if newAPIClient.GetExpireTime() != initialExpireTime {
		t.Errorf("Reset() did not maintain expireTime setting. Expected: %v, Got: %v",
			initialExpireTime, newAPIClient.GetExpireTime())
	}
}

func TestSingletonBehavior(t *testing.T) {
	// Create first instance
	client1 := NewClient()
	if client1 == nil {
		t.Fatal("First NewClient() returned nil")
	}

	client2 := NewClient()
	if client2 == nil {
		t.Fatal("Second NewClient() returned nil")
	}

	if client1 != client2 {
		t.Error("NewClient() did not maintain singleton behavior")
	}
}

func TestGetClient(t *testing.T) {
	client := NewClient()
	client2 := client
	if client != client2 {
		t.Error("GetClient did not return the client")
	}
}

func TestGetClientSingleton(t *testing.T) {
	client := NewClient()
	client2 := client.apiClient
	client3 := client.apiClient
	if client2 != client.apiClient {
		t.Error("GetClient did not return the same client")
	}
	if client2 != client3 {
		t.Error("GetClient did not return a singleton")
	}
}
