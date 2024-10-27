package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestLocations(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Locations("0", "20")
	if err != nil {
		t.Error("Fetching Locations returning error")
	}
	var emptyLocations models.Locations
	if reflect.DeepEqual(emptyLocations, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestLocation_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Location("1")
	if err != nil {
		t.Error("Fetching Location with ID returning error")
	}
	var emptyLocation models.Location
	if reflect.DeepEqual(emptyLocation, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "canalave-city" {
		t.Error("Fetch result returns wrong location name")
	}
}

func TestLocation_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Location("canalave-city")
	if err != nil {
		t.Error("Fetching Location with ID returning error")
	}
	var emptyLocation models.Location
	if reflect.DeepEqual(emptyLocation, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "canalave-city" {
		t.Error("Fetch result returns wrong location name")
	}
}

func TestLocation_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Location("abcd")
	if err == nil {
		t.Error("Fetching wrong Location with Name returning no error")
	}
}

func TestLocation_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Location("")
	if err == nil {
		t.Error("Fetching wrong Location with empty ID returning no error")
	}
}

func TestLocation_FailWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Location("0")
	if err == nil {
		t.Error("Fetching wrong Location with empty Name returning no error")
	}
}

func TestLocationAreas(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.LocationAreas("0", "20")
	if err != nil {
		t.Error("Fetching Locations returning error")
	}
	var emptyLocationAreas models.LocationAreas
	if reflect.DeepEqual(emptyLocationAreas, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestLocationArea_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.LocationArea("1")
	if err != nil {
		t.Error("Fetching Location with ID returning error")
	}
	var emptyLocationArea models.LocationArea
	if reflect.DeepEqual(emptyLocationArea, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "canalave-city-area" {
		t.Error("Fetch result returns wrong location name")
	}
}

func TestLocationArea_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.LocationArea("canalave-city-area")
	if err != nil {
		t.Error("Fetching Location with ID returning error")
	}
	var emptyLocationArea models.LocationArea
	if reflect.DeepEqual(emptyLocationArea, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "canalave-city-area" {
		t.Error("Fetch result returns wrong location name")
	}
}

func TestLocationArea_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.LocationArea("abcd")
	if err == nil {
		t.Error("Fetching wrong Location with Name returning no error")
	}
}

func TestLocationArea_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.LocationArea("")
	if err == nil {
		t.Error("Fetching wrong Location with empty ID returning no error")
	}
}

func TestLocationArea_FailWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.LocationArea("0")
	if err == nil {
		t.Error("Fetching wrong Location with empty Name returning no error")
	}
}

func TestPalParkAreas(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PalParkAreas("0", "20")
	if err != nil {
		t.Error("Fetching Locations returning error")
	}
	var emptyPalParkAreas models.PalParkAreas
	if reflect.DeepEqual(emptyPalParkAreas, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPalParkArea_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PalParkArea("1")
	if err != nil {
		t.Error("Fetching Location with ID returning error")
	}
	var emptyPalParkArea models.PalParkArea
	if reflect.DeepEqual(emptyPalParkArea, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "forest" {
		t.Error("Fetch result returns wrong location name")
	}
}

func TestPalParkArea_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PalParkArea("forest")
	if err != nil {
		t.Error("Fetching Location with ID returning error")
	}
	var emptyPalParkArea models.PalParkArea
	if reflect.DeepEqual(emptyPalParkArea, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "forest" {
		t.Error("Fetch result returns wrong location name")
	}
}

func TestPalParkArea_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PalParkArea("abcd")
	if err == nil {
		t.Error("Fetching wrong Location with Name returning no error")
	}
}

func TestPalParkArea_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PalParkArea("")
	if err == nil {
		t.Error("Fetching wrong Location with empty ID returning no error")
	}
}

func TestPalParkArea_FailWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PalParkArea("0")
	if err == nil {
		t.Error("Fetching wrong Location with empty Name returning no error")
	}
}

func TestRegions(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Regions("0", "20")
	if err != nil {
		t.Error("Fetching Regions returning error")
	}
	var emptyRegions models.Regions
	if reflect.DeepEqual(emptyRegions, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestRegion_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Region("1")
	if err != nil {
		t.Error("Fetching Region with ID returning error")
	}
	var emptyRegion models.Region
	if reflect.DeepEqual(emptyRegion, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "kanto" {
		t.Error("Fetch result returns wrong region name")
	}
}

func TestRegion_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Region("kanto")
	if err != nil {
		t.Error("Fetching Region with ID returning error")
	}
	var emptyRegion models.Region
	if reflect.DeepEqual(emptyRegion, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "kanto" {
		t.Error("Fetch result returns wrong region name")
	}
}

func TestRegion_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Region("abcd")
	if err == nil {
		t.Error("Fetching wrong Region with Name returning no error")
	}
}

func TestRegion_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Region("")
	if err == nil {
		t.Error("Fetching wrong Region with empty ID returning no error")
	}
}

func TestRegion_FailWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Region("0")
	if err == nil {
		t.Error("Fetching wrong Region with empty Name returning no error")
	}
}
