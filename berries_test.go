package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestBerry(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Berries("0", "20")
	if err != nil {
		t.Error("Fetching Berries returning error")
	}
	var emptyBerry models.Berries
	if reflect.DeepEqual(emptyBerry, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestBerry_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Berry("1")
	if err != nil {
		t.Error("Fetching Berry with ID returning error")
	}
	var emptyBerry models.Berry
	if reflect.DeepEqual(emptyBerry, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "cheri" {
		t.Error("Fetch result returns wrong Name")
	}
	if result.GrowthTime != 3 {
		t.Error("Fetch result returns wrong GrowthTime")
	}
	if result.MaxHarvest != 5 {
		t.Error("Fetch result returns wrong MaxHarvest")
	}
	if result.NaturalGiftPower != 60 {
		t.Error("Fetch result returns wrong NaturalGiftPower")
	}
	if result.Size != 20 {
		t.Error("Fetch result returns wrong Size")
	}
	if result.Smoothness != 25 {
		t.Error("Fetch result returns wrong Smoothness")
	}
	if result.SoilDryness != 15 {
		t.Error("Fetch result returns wrong SoilDryness")
	}
}

func TestBerry_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Berry("cheri")
	if err != nil {
		t.Error("Fetching Berry with ID returning error")
	}
	var emptyBerry models.Berry
	if reflect.DeepEqual(emptyBerry, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "cheri" {
		t.Error("Fetch result returns wrong Name")
	}
	if result.GrowthTime != 3 {
		t.Error("Fetch result returns wrong GrowthTime")
	}
	if result.MaxHarvest != 5 {
		t.Error("Fetch result returns wrong MaxHarvest")
	}
	if result.NaturalGiftPower != 60 {
		t.Error("Fetch result returns wrong NaturalGiftPower")
	}
	if result.Size != 20 {
		t.Error("Fetch result returns wrong Size")
	}
	if result.Smoothness != 25 {
		t.Error("Fetch result returns wrong Smoothness")
	}
	if result.SoilDryness != 15 {
		t.Error("Fetch result returns wrong SoilDryness")
	}
}

func TestBerry_Fail(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Berry("0")
	if err == nil {
		t.Error("Fetching Berry with ID 0 should return error")
	}
}

func TestBerryFirmness_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.BerryFirmness("1")
	if err != nil {
		t.Error("Fetching Berry Firmness returning error")
	}
	var emptyBerryFirmness models.BerryFirmness
	if reflect.DeepEqual(emptyBerryFirmness, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "very-soft" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestBerryFirmness_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.BerryFirmness("very-soft")
	if err != nil {
		t.Error("Fetching Berry Firmness returning error")
	}
	var emptyBerryFirmness models.BerryFirmness
	if reflect.DeepEqual(emptyBerryFirmness, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "very-soft" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestBerryFirmness_Invalid(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.BerryFirmness("abcs")
	if err == nil {
		t.Error("Fetching Berry Firmness with random string should return error")
	}
}

func TestBerryFlavor_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.BerryFlavor("1")
	if err != nil {
		t.Error("Fetching Berry Flavor returning error")
	}
	var emptyBerryFlavor models.BerryFlavor
	if reflect.DeepEqual(emptyBerryFlavor, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "spicy" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestBerryFlavor_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.BerryFlavor("spicy")
	if err != nil {
		t.Error("Fetching Berry Flavor returning error")
	}
	var emptyBerryFlavor models.BerryFlavor
	if reflect.DeepEqual(emptyBerryFlavor, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "spicy" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestBerryFlavor_Invalid(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.BerryFlavor("abcs")
	if err == nil {
		t.Error("Fetching Berry Flavor with random string should return error")
	}
}
