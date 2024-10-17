package api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/hanle23/pokego/internal/models"
)

func TestFetch(t *testing.T) {
	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := httpClient.Get("https://pokeapi.co/api/v2/berry/?offset=0&limit=20")
	if err != nil {
		t.Fatalf("Failed to make HTTP request: %v", err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	defer resp.Body.Close()
	var resultFetch models.Berries
	err = json.Unmarshal(bodyBytes, &resultFetch)
	if err != nil {
		t.Fatalf("Failed to marshal result to JSON: %v", err)
	}

	var resultAPI models.Berries
	err = client.Fetch("berry/?offset=0&limit=20", &resultAPI)
	if err != nil {
		t.Fatalf("Failed to fetch through API: %v", err)
	}
	if !reflect.DeepEqual(resultAPI, resultFetch) {
		t.Fatalf("Data mismatch")
	}
}
