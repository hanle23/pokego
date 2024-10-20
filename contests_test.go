package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestContestType_InvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ContestType("0")
	if err == nil {
		t.Error("Fetching invalid ContestType returning no error")
	}
}

func TestContestType_EmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ContestType("")
	if err == nil {
		t.Error("Fetching ContestType with invalid ID returning no error")
	}
}

func TestContestType_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ContestType("1")
	if err != nil {
		t.Error("Fetching ContestType with ID returning error")
	}
	var emptyContestType models.ContestType
	if reflect.DeepEqual(emptyContestType, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "cool" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestContestType_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ContestType("cool")
	if err != nil {
		t.Error("Fetching ContestType with ID returning error")
	}
	var emptyContestType models.ContestType
	if reflect.DeepEqual(emptyContestType, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "cool" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestContestTypes(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ContestTypes("0", "20")
	if err != nil {
		t.Error("Fetching ContestTypes returning error")
	}
	var emptyContestTypes models.ContestTypes
	if reflect.DeepEqual(emptyContestTypes, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestContestEffect_InvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ContestEffect("0")
	if err == nil {
		t.Error("Fetching invalid ContestEffect returning no error")
	}
}

func TestContestEffect_InvalidID2(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ContestEffect("string")
	if err == nil {
		t.Error("Fetching invalid ContestEffect with string returning no error")
	}
}

func TestContestEffect_EmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ContestEffect("")
	if err == nil {
		t.Error("Fetching ContestEffect with invalid ID returning no error")
	}
}

func TestContestEffect_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ContestEffect("1")
	if err != nil {
		t.Error("Fetching ContestEffect with ID returning error")
	}
	var emptyContestEffect models.ContestEffect
	if reflect.DeepEqual(emptyContestEffect, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Appeal != 4 {
		t.Error("Fetch result returns wrong Appeal")
	}
}

func TestContestEffects(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ContestEffects("0", "20")
	if err != nil {
		t.Error("Fetching ContestEffects returning error")
	}
	var emptyContestEffects models.ContestEffects
	if reflect.DeepEqual(emptyContestEffects, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestSuperContestEffect_InvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.SuperContestEffect("0")
	if err == nil {
		t.Error("Fetching invalid SuperContestEffect returning no error")
	}
}

func TestSuperContestEffect_InvalidID2(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.SuperContestEffect("string")
	if err == nil {
		t.Error("Fetching invalid SuperContestEffect with string returning no error")
	}
}

func TestSuperContestEffect_EmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.SuperContestEffect("")
	if err == nil {
		t.Error("Fetching SuperContestEffect with invalid ID returning no error")
	}
}

func TestSuperContestEffect_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.SuperContestEffect("1")
	if err != nil {
		t.Error("Fetching SuperContestEffect with ID returning error")
	}
	var emptySuperContestEffect models.SuperContestEffect
	if reflect.DeepEqual(emptySuperContestEffect, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Appeal != 2 {
		t.Error("Fetch result returns wrong Appeal")
	}
}

func TestSuperContestEffects(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.SuperContestEffects("0", "20")
	if err != nil {
		t.Error("Fetching SuperContestEffects returning error")
	}
	var emptySuperContestEffects models.SuperContestEffects
	if reflect.DeepEqual(emptySuperContestEffects, result) {
		t.Error("Fetch result returns nil")
	}
}
