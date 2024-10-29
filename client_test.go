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
	if client != client2 {
		t.Error("NewClient did not return a singleton")
	}
}

func TestRemoveClient(t *testing.T) {
	// Create the initial client instance
	client := NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}

	client.RemoveClient()

	// Check that the GetClient() returns a different instance
	copyClient := client.GetClient()
	if copyClient != nil {
		t.Error("GetClient() did not return nil after removal")
	}
}

func TestResetClient(t *testing.T) {
	client := NewClient(WithExpireTime(50), WithUseCache(false))
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	copyClient := client.GetClient()
	client.ResetClient()
	client2 := client.GetClient()
	if copyClient == client2 {
		t.Error("ResetClient did not reset the client")
	}
	if copyClient.apiClient.GetUseCache() != client2.apiClient.GetUseCache() {
		t.Error("ResetClient did not properly copy config useCache")
	}
	if copyClient.apiClient.GetExpireTime() != client2.apiClient.GetExpireTime() {
		t.Error("ResetClient did not properly copy config expireTime")
	}
}

func TestGetClient(t *testing.T) {
	client := NewClient()
	client2 := client.GetClient()
	if client != client2 {
		t.Error("GetClient did not return the client")
	}
}

func TestGetClientSingleton(t *testing.T) {
	client := NewClient()
	client2 := client.GetClient()
	client3 := client.GetClient()
	if client2 != client {
		t.Error("GetClient did not return the same client")
	}
	if client2 != client3 {
		t.Error("GetClient did not return a singleton")
	}
}
