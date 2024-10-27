package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestEvolutionChains(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EvolutionChains("0", "20")
	if err != nil {
		t.Error("Fetching EvolutionChains returning error")
	}
	var emptyEvolutionChain models.EvolutionChains
	if reflect.DeepEqual(emptyEvolutionChain, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestEvolutionChain_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EvolutionChain("1")
	if err != nil {
		t.Error("Fetching EvolutionChain with ID returning error")
	}
	var emptyEvolutionChain models.EvolutionChain
	if reflect.DeepEqual(emptyEvolutionChain, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Chain.IsBaby != false {
		t.Error("Fetch result returns wrong IsBaby")
	}
	if result.Chain.Species.Name != "bulbasaur" {
		t.Error("Fetch result returns wrong chain species name")
	}
}

func TestEvolutionChain_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EvolutionChain("bulbasaur")
	if err == nil {
		t.Error("Fetching wrong EvolutionChain with Name returning no error")
	}
}

func TestEvolutionChain_FailWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EvolutionChain("0")
	if err == nil {
		t.Error("Fetching wrong EvolutionChain with invalid ID returning no error")
	}
}

func TestEvolutionTriggers(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EvolutionTriggers("0", "20")
	if err != nil {
		t.Error("Fetching EvolutionTriggers returning error")
	}
	var emptyEvolutionTrigger models.EvolutionTriggers
	if reflect.DeepEqual(emptyEvolutionTrigger, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestEvolutionTrigger_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EvolutionTrigger("1")
	if err != nil {
		t.Error("Fetching EvolutionTrigger with ID returning error")
	}
	var emptyEvolutionTrigger models.EvolutionTrigger
	if reflect.DeepEqual(emptyEvolutionTrigger, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "level-up" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestEvolutionTrigger_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EvolutionTrigger("level-up")
	if err != nil {
		t.Error("Fetching wrong EvolutionTrigger with Name returning no error")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "level-up" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestEvolutionTrigger_FailWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EvolutionTrigger("0")
	if err == nil {
		t.Error("Fetching wrong EvolutionTrigger with invalid ID returning no error")
	}
}

func TestEvolutionTrigger_FailWithInvalidName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EvolutionTrigger("invalid")
	if err == nil {
		t.Error("Fetching wrong EvolutionTrigger with invalid Name returning no error")
	}
}

func TestEvolutionTrigger_FailWithEmptyName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EvolutionTrigger("")
	if err == nil {
		t.Error("Fetching wrong EvolutionTrigger with empty Name returning no error")
	}
}
