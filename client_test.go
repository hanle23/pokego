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
		t.Error("NewClient() returned nil")
	}

	client.Close()

	copyClient := client.apiClient
	if copyClient != nil {
		t.Error("GetClient() did not return nil after removal")
	}
}

func TestResetClient(t *testing.T) {
	client := NewClient(WithExpireTime(50), WithUseCache(false))
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	copyClient := client.apiClient
	client.Reset()
	client2 := client.apiClient
	if copyClient == client2 {
		t.Error("ResetClient did not reset the client")
	}
	if copyClient.GetUseCache() != client2.GetUseCache() {
		t.Error("ResetClient did not properly copy config useCache")
	}
	if copyClient.GetExpireTime() != client2.GetExpireTime() {
		t.Error("ResetClient did not properly copy config expireTime")
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
