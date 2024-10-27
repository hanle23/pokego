package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestLanguages(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Languages("0", "20")
	if err != nil {
		t.Error("Fetching Languages returning error")
	}
	var emptyLanguages models.Languages
	if reflect.DeepEqual(emptyLanguages, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestLanguage_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Language("1")
	if err != nil {
		t.Error("Fetching Language with ID returning error")
	}
	var emptyLanguage models.Language
	if reflect.DeepEqual(emptyLanguage, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "ja-Hrkt" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestLanguage_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Language("ja-Hrkt")
	if err != nil {
		t.Error("Fetching Language with Name returning error")
	}
	var emptyLanguage models.Language
	if reflect.DeepEqual(emptyLanguage, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "ja-Hrkt" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestLanguage_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Language("0")
	if err == nil {
		t.Error("Fetching wrong Language with ID returning no error")
	}
}

func TestLanguage_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Language("ja-Hrktt")
	if err == nil {
		t.Error("Fetching wrong Language with Name returning no error")
	}
}

func TestLanguage_FailWithEmptyName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Language("")
	if err == nil {
		t.Error("Fetching wrong Language with empty Name returning no error")
	}
}
