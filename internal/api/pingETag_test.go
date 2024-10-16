package api

import (
	"net/http"
	"testing"
)

func TestPingETag(t *testing.T) {
	resp, err := client.httpClient.Get("https://pokeapi.co/api/v2/pokemon/1")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	etag := resp.Header.Get("Etag")
	if etag == "" {
		t.Error("Etag is empty")
	}
	isChanged, err := client.PingETag("pokemon/1", etag)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if isChanged {
		t.Error("Expected isChanged to be false")
	}
}
