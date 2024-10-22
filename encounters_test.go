package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestEncounterMethod_InvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EncounterMethod("0")
	if err == nil {
		t.Error("Fetching invalid Encounter Method returning no error")
	}
}

func TestEncounterMethod_EmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EncounterMethod("")
	if err == nil {
		t.Error("Fetching Encounter Method with invalid ID returning no error")
	}
}

func TestEncounterMethod_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EncounterMethod("1")
	if err != nil {
		t.Error("Fetching Encounter Method with ID returning error")
	}
	var emptyEncounterMethod models.EncounterMethod
	if reflect.DeepEqual(emptyEncounterMethod, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "walk" {
		t.Error("Fetch result returns wrong Name")
	}
}
