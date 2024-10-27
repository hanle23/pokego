package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestItems(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Items("0", "20")
	if err != nil {
		t.Error("Fetching Items returning error")
	}
	var emptyItems models.Items
	if reflect.DeepEqual(emptyItems, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestItem_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Item("1")
	if err != nil {
		t.Error("Fetching Item with ID returning error")
	}
	var emptyItem models.Item
	if reflect.DeepEqual(emptyItem, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "master-ball" {
		t.Error("Fetch result returns wrong name")
	}
	if result.Cost != 0 {
		t.Error("Fetch result returns wrong cost")
	}
}

func TestItemAttributes(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemAttributes("0", "20")
	if err != nil {
		t.Error("Fetching EvolutionChains returning error")
	}
	var emptyItemAttributes models.ItemAttributes
	if reflect.DeepEqual(emptyItemAttributes, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestItemAttribute_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemAttribute("1")
	if err != nil {
		t.Error("Fetching ItemAttribute with ID returning error")
	}
	var emptyItemAttribute models.ItemAttribute
	if reflect.DeepEqual(emptyItemAttribute, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "countable" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestItemAttribute_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemAttribute("countable")
	if err != nil {
		t.Error("Fetching ItemAttribute with ID returning error")
	}
	var emptyItemAttribute models.ItemAttribute
	if reflect.DeepEqual(emptyItemAttribute, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "countable" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestItemAttribute_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ItemAttribute("abcd")
	if err == nil {
		t.Error("Fetching wrong ItemAttribute with Name returning no error")
	}
}

func TestItemAttribute_FailWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ItemAttribute("0")
	if err == nil {
		t.Error("Fetching wrong ItemAttribute with Name returning no error")
	}
}

func TestItemCategories(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemCategories("0", "20")
	if err != nil {
		t.Error("Fetching ItemCategories returning error")
	}
	var emptyItemCategories models.ItemCategories
	if reflect.DeepEqual(emptyItemCategories, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestItemCategory_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemCategory("1")
	if err != nil {
		t.Error("Fetching ItemCategory with ID returning error")
	}
	var emptyItemCategory models.ItemCategory
	if reflect.DeepEqual(emptyItemCategory, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "stat-boosts" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestItemCategory_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemCategory("stat-boosts")
	if err != nil {
		t.Error("Fetching ItemCategory with ID returning error")
	}
	var emptyItemCategory models.ItemCategory
	if reflect.DeepEqual(emptyItemCategory, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "stat-boosts" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestItemCategory_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ItemCategory("abcd")
	if err == nil {
		t.Error("Fetching wrong ItemCategory with Name returning no error")
	}
}

func TestItemCategory_FailWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ItemCategory("0")
	if err == nil {
		t.Error("Fetching wrong ItemCategory with Name returning no error")
	}
}

func TestItemFlingEffects(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemFlingEffects("0", "20")
	if err != nil {
		t.Error("Fetching ItemFlingEffects returning error")
	}
	var emptyItemFlingEffects models.ItemFlingEffects
	if reflect.DeepEqual(emptyItemFlingEffects, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestItemFlingEffect_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemFlingEffect("1")
	if err != nil {
		t.Error("Fetching ItemFlingEffect with ID returning error")
	}
	var emptyItemFlingEffect models.ItemFlingEffect
	if reflect.DeepEqual(emptyItemFlingEffect, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "badly-poison" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestItemFlingEffect_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemFlingEffect("badly-poison")
	if err != nil {
		t.Error("Fetching ItemFlingEffect with ID returning error")
	}
	var emptyItemFlingEffect models.ItemFlingEffect
	if reflect.DeepEqual(emptyItemFlingEffect, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "badly-poison" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestItemFlingEffect_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ItemFlingEffect("abcd")
	if err == nil {
		t.Error("Fetching wrong ItemFlingEffect with Name returning no error")
	}
}

func TestItemFlingEffect_FailWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ItemFlingEffect("0")
	if err == nil {
		t.Error("Fetching wrong ItemFlingEffect with Name returning no error")
	}
}

func TestItemPockets(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemPockets("0", "20")
	if err != nil {
		t.Error("Fetching ItemPockets returning error")
	}
	var emptyItemPockets models.ItemPockets
	if reflect.DeepEqual(emptyItemPockets, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestItemPocket_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemPocket("1")
	if err != nil {
		t.Error("Fetching ItemPocket with ID returning error")
	}
	var emptyItemPocket models.ItemPocket
	if reflect.DeepEqual(emptyItemPocket, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "misc" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestItemPocket_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ItemPocket("misc")
	if err != nil {
		t.Error("Fetching ItemPocket with ID returning error")
	}
	var emptyItemPocket models.ItemPocket
	if reflect.DeepEqual(emptyItemPocket, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "misc" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestItemPocket_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ItemPocket("abcd")
	if err == nil {
		t.Error("Fetching wrong ItemPocket with Name returning no error")
	}
}

func TestItemPocket_FailWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.ItemPocket("0")
	if err == nil {
		t.Error("Fetching wrong ItemPocket with Name returning no error")
	}
}
