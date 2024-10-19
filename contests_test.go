package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestContestType(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.ContestType("0")
	if err != nil {
		t.Error("Fetching ContestType returning error")
	}
	var emptyContestType models.ContestType
	if reflect.DeepEqual(emptyContestType, result) {
		t.Error("Fetch result returns nil")
	}
}
