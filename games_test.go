package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestGenerations(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Generations("0", "20")
	if err != nil {
		t.Error("Fetching Generations returning error")
	}
	var emptyGenerations models.Generations
	if reflect.DeepEqual(emptyGenerations, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestGeneration_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Generation("1")
	if err != nil {
		t.Error("Fetching Generation with ID returning error")
	}
	var emptyGeneration models.Generation
	if reflect.DeepEqual(emptyGeneration, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.MainRegion.Name != "kanto" {
		t.Error("Fetch result returns wrong MainRegion")
	}
	if result.Name != "generation-i" {
		t.Error("Fetch result returns wrong Name")
	}
	if result.PokemonSpecies[0].Name != "bulbasaur" {
		t.Error("Fetch result returns wrong PokemonSpecies")
	}
	if result.Types[0].Name != "normal" {
		t.Error("Fetch result returns wrong Types")
	}
}

func TestGeneration_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Generation("generation-i")
	if err != nil {
		t.Error("Fetching Generation with ID returning error")
	}
	var emptyGeneration models.Generation
	if reflect.DeepEqual(emptyGeneration, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.MainRegion.Name != "kanto" {
		t.Error("Fetch result returns wrong MainRegion")
	}
	if result.Name != "generation-i" {
		t.Error("Fetch result returns wrong Name")
	}
	if result.PokemonSpecies[0].Name != "bulbasaur" {
		t.Error("Fetch result returns wrong PokemonSpecies")
	}
	if result.Types[0].Name != "normal" {
		t.Error("Fetch result returns wrong Types")
	}
}

func TestGeneration_Fail(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Generation("0")
	if err == nil {
		t.Error("Fetching Generation with ID 0 should return error")
	}
}

func TestGeneration_FailEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Generation("")
	if err == nil {
		t.Error("Fetching Generation with empty ID should return error")
	}
}

func TestPokedexes(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Pokedexes("0", "20")
	if err != nil {
		t.Error("Fetching Pokedexes returning error")
	}
	var emptyGenerations models.Generations
	if reflect.DeepEqual(emptyGenerations, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPokedex_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Pokedex("1")
	if err != nil {
		t.Error("Fetching Pokedex with ID returning error")
	}
	var emptyGeneration models.Generation
	if reflect.DeepEqual(emptyGeneration, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "national" {
		t.Error("Fetch result returns wrong Name")
	}
	if result.PokemonEntries[0].EntryNumber != 1 {
		t.Error("Fetch result returns wrong EntryNumber")
	}
	if result.PokemonEntries[0].PokemonSpecies.Name != "bulbasaur" {
		t.Error("Fetch result returns wrong PokemonSpecies")
	}
}

func TestPokedex_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Pokedex("national")
	if err != nil {
		t.Error("Fetching Pokedex with ID returning error")
	}
	var emptyGeneration models.Generation
	if reflect.DeepEqual(emptyGeneration, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "national" {
		t.Error("Fetch result returns wrong Name")
	}
	if result.PokemonEntries[0].EntryNumber != 1 {
		t.Error("Fetch result returns wrong EntryNumber")
	}
	if result.PokemonEntries[0].PokemonSpecies.Name != "bulbasaur" {
		t.Error("Fetch result returns wrong PokemonSpecies")
	}
}

func TestPokedex_Fail(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Pokedex("0")
	if err == nil {
		t.Error("Fetching Pokedex with ID 0 should return error")
	}
}

func TestPokedex_FailEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Pokedex("")
	if err == nil {
		t.Error("Fetching Pokedex with empty ID should return error")
	}
}

func TestVersions(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Versions("0", "20")
	if err != nil {
		t.Error("Fetching Versions returning error")
	}
	var emptyVersions models.Versions
	if reflect.DeepEqual(emptyVersions, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestVersion_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Version("1")
	if err != nil {
		t.Error("Fetching Version with ID returning error")
	}
	var emptyVersion models.Version
	if reflect.DeepEqual(emptyVersion, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "red" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestVersion_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Version("red")
	if err != nil {
		t.Error("Fetching Version with ID returning error")
	}
	var emptyVersion models.Version
	if reflect.DeepEqual(emptyVersion, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "red" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestVersion_Fail(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Version("0")
	if err == nil {
		t.Error("Fetching Version with ID 0 should return error")
	}
}

func TestVersion_FailEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Version("")
	if err == nil {
		t.Error("Fetching Version with empty ID should return error")
	}
}

func TestVersionGroups(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.VersionGroups("0", "20")
	if err != nil {
		t.Error("Fetching Version Groups returning error")
	}
	var emptyVersionGroups models.VersionGroups
	if reflect.DeepEqual(emptyVersionGroups, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestVersionGroup_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.VersionGroup("1")
	if err != nil {
		t.Error("Fetching Version with ID returning error")
	}
	var emptyVersionGroup models.VersionGroup
	if reflect.DeepEqual(emptyVersionGroup, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "red-blue" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestVersionGroup_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.VersionGroup("red-blue")
	if err != nil {
		t.Error("Fetching Version with ID returning error")
	}
	var emptyVersionGroup models.VersionGroup
	if reflect.DeepEqual(emptyVersionGroup, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "red-blue" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestVersionGroup_Fail(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.VersionGroup("0")
	if err == nil {
		t.Error("Fetching Version Group with ID 0 should return error")
	}
}

func TestVersionGroup_FailEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.VersionGroup("")
	if err == nil {
		t.Error("Fetching Version Group with empty ID should return error")
	}
}
