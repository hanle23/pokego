package pokego

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego/internal/models"
)

func TestFetchBerry(t *testing.T) {
	client := NewClient()
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
