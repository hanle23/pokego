package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestMachines(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Machines("0", "20")
	if err != nil {
		t.Error("Fetching Machines returning error")
	}
	var emptyMachines models.Machines
	if reflect.DeepEqual(emptyMachines, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestMachine_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Machine("1")
	if err != nil {
		t.Error("Fetching Machine with ID returning error")
	}
	var emptyMachine models.Machine
	if reflect.DeepEqual(emptyMachine, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Item.Name != "tm00" {
		t.Error("Fetch result returns wrong Item Name")
	}
	if result.Move.Name != "mega-punch" {
		t.Error("Fetch result returns wrong Move name")
	}
}

func TestMachine_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Machine("test")
	if err == nil {
		t.Error("Fetching Machine with Fail didn't return error")
	}
}

func TestMachine_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Machine("")
	if err == nil {
		t.Error("Fetching Machine with empty ID didn't return error")
	}
}

func TestMachine_FailWithNegativeID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Machine("-1")
	if err == nil {
		t.Error("Fetching Machine with negative ID didn't return error")
	}
}
