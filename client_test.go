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

	RemoveClient()

	// Check that the GetClient() returns a different instance
	copyClient := GetClient()
	if copyClient != nil {
		t.Error("GetClient() did not return nil after removal")
	}

}

func TestGetClient(t *testing.T) {
	client := NewClient()
	client2 := GetClient()
	if client != client2 {
		t.Error("GetClient did not return the client")
	}
}

func TestGetClientSingleton(t *testing.T) {
	client := NewClient()
	client2 := GetClient()
	client3 := GetClient()
	if client2 != client {
		t.Error("GetClient did not return the same client")
	}
	if client2 != client3 {
		t.Error("GetClient did not return a singleton")
	}
}
