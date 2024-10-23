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
func TestEncounterMethod_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EncounterMethod("walk")
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

func TestEncounterMethods(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EncounterMethods("0", "20")
	if err != nil {
		t.Error("Fetching Encounter Methods returning error")
	}
	var emptyContestTypes models.EncounterMethods
	if reflect.DeepEqual(emptyContestTypes, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestEncounterCondition_InvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EncounterCondition("0")
	if err == nil {
		t.Error("Fetching invalid Encounter Condition returning no error")
	}
}

func TestEncounterCondition_EmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EncounterCondition("")
	if err == nil {
		t.Error("Fetching Encounter Condition with invalid ID returning no error")
	}
}

func TestEncounterCondition_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EncounterCondition("1")
	if err != nil {
		t.Error("Fetching Encounter Condition with ID returning error")
	}
	var emptyEncounterCondition models.EncounterCondition
	if reflect.DeepEqual(emptyEncounterCondition, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "swarm" {
		t.Errorf("Fetch result returns wrong Name, expect %v, got %v", "swarm", result.Name)
	}
}
func TestEncounterCondition_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EncounterCondition("swarm")
	if err != nil {
		t.Error("Fetching Encounter Condition with ID returning error")
	}
	var emptyEncounterCondition models.EncounterCondition
	if reflect.DeepEqual(emptyEncounterCondition, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "swarm" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestEncounterConditions(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EncounterConditions("0", "20")
	if err != nil {
		t.Error("Fetching Encounter Conditions returning error")
	}
	var emptyEncounterConditions models.EncounterConditions
	if reflect.DeepEqual(emptyEncounterConditions, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestEncounterConditionValue_InvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EncounterConditionValue("0")
	if err == nil {
		t.Error("Fetching invalid Encounter Condition Value returning no error")
	}
}

func TestEncounterConditionValue_EmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EncounterConditionValue("")
	if err == nil {
		t.Error("Fetching Encounter Condition Value with invalid ID returning no error")
	}
}

func TestEncounterConditionValue_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EncounterConditionValue("1")
	if err != nil {
		t.Error("Fetching Encounter Condition Value with ID returning error")
	}
	var emptyEncounterConditionValue models.EncounterConditionValue
	if reflect.DeepEqual(emptyEncounterConditionValue, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Errorf("Fetch result returns wrong ID, expect %v, got %v", 1, result.ID)
	}
	if result.Name != "swarm-yes" {
		t.Errorf("Fetch result returns wrong Name, expect %v, got %v", "swarm-yes", result.Name)
	}
}
func TestEncounterConditionValue_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EncounterConditionValue("swarm-yes")
	if err != nil {
		t.Error("Fetching Encounter Condition Value with ID returning error")
	}
	var emptyEncounterConditionValue models.EncounterConditionValue
	if reflect.DeepEqual(emptyEncounterConditionValue, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "swarm-yes" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestEncounterConditionValues(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EncounterConditionValues("0", "20")
	if err != nil {
		t.Error("Fetching Encounter Condition Values returning error")
	}
	var emptyEncounterConditionValues models.EncounterConditionValues
	if reflect.DeepEqual(emptyEncounterConditionValues, result) {
		t.Error("Fetch result returns nil")
	}
}
