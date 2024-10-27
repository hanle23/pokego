package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestMoves(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Moves("0", "20")
	if err != nil {
		t.Error("Fetching Moves returning error")
	}
	var emptyMoves models.Moves
	if reflect.DeepEqual(emptyMoves, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestMove_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Move("1")
	if err != nil {
		t.Error("Fetching Move with ID returning error")
	}
	var emptyMove models.Move
	if reflect.DeepEqual(emptyMove, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "pound" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMove_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Move("pound")
	if err != nil {
		t.Error("Fetching Move with ID returning error")
	}
	var emptyMove models.Move
	if reflect.DeepEqual(emptyMove, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "pound" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMove_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Move("0")
	if err == nil {
		t.Error("Fetching wrong Move with ID returning no error")
	}
}

func TestMove_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Move("wrong")
	if err == nil {
		t.Error("Fetching wrong Move with Name returning no error")
	}
}

func TestMove_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Move("")
	if err == nil {
		t.Error("Fetching wrong Move with empty ID returning no error")
	}
}

func TestMoveAilments(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveAilments("0", "20")
	if err != nil {
		t.Error("Fetching MoveAilments returning error")
	}
	var emptyMoveAilments models.MoveAilments
	if reflect.DeepEqual(emptyMoveAilments, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestMoveAilment_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveAilment("1")
	if err != nil {
		t.Error("Fetching MoveAilment with ID returning error")
	}
	var emptyMoveAilment models.MoveAilment
	if reflect.DeepEqual(emptyMoveAilment, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "paralysis" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveAilment_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveAilment("paralysis")
	if err != nil {
		t.Error("Fetching MoveAilment with ID returning error")
	}
	var emptyMoveAilment models.MoveAilment
	if reflect.DeepEqual(emptyMoveAilment, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "paralysis" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveAilment_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveAilment("-2")
	if err == nil {
		t.Error("Fetching wrong MoveAilment with ID returning no error")
	}
}

func TestMoveAilment_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveAilment("wrong")
	if err == nil {
		t.Error("Fetching wrong MoveAilment with Name returning no error")
	}
}

func TestMoveAilment_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveAilment("")
	if err == nil {
		t.Error("Fetching wrong MoveAilment with empty ID returning no error")
	}
}

func TestMoveBattleStyles(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveBattleStyles("0", "20")
	if err != nil {
		t.Error("Fetching MoveBattleStyles returning error")
	}
	var emptyMoveBattleStyles models.MoveBattleStyles
	if reflect.DeepEqual(emptyMoveBattleStyles, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestMoveBattleStyle_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveBattleStyle("1")
	if err != nil {
		t.Error("Fetching MoveBattleStyle with ID returning error")
	}
	var emptyMoveBattleStyle models.MoveBattleStyle
	if reflect.DeepEqual(emptyMoveBattleStyle, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "attack" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveBattleStyle_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveBattleStyle("attack")
	if err != nil {
		t.Error("Fetching MoveBattleStyle with ID returning error")
	}
	var emptyMoveBattleStyle models.MoveBattleStyle
	if reflect.DeepEqual(emptyMoveBattleStyle, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "attack" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveBattleStyle_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveBattleStyle("0")
	if err == nil {
		t.Error("Fetching wrong MoveBattleStyle with ID returning no error")
	}
}

func TestMoveBattleStyle_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveBattleStyle("wrong")
	if err == nil {
		t.Error("Fetching wrong MoveBattleStyle with Name returning no error")
	}
}

func TestMoveBattleStyle_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveBattleStyle("")
	if err == nil {
		t.Error("Fetching wrong MoveBattleStyle with empty ID returning no error")
	}
}

func TestMoveCategories(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveCategories("0", "20")
	if err != nil {
		t.Error("Fetching MoveCategories returning error")
	}
	var emptyMoveCategories models.MoveCategories
	if reflect.DeepEqual(emptyMoveCategories, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestMoveCategory_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveCategory("1")
	if err != nil {
		t.Error("Fetching MoveCategory with ID returning error")
	}
	var emptyMoveCategory models.MoveCategory
	if reflect.DeepEqual(emptyMoveCategory, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "ailment" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveCategory_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveCategory("ailment")
	if err != nil {
		t.Error("Fetching MoveCategory with ID returning error")
	}
	var emptyMoveCategory models.MoveCategory
	if reflect.DeepEqual(emptyMoveCategory, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "ailment" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveCategory_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveCategory("-1")
	if err == nil {
		t.Error("Fetching wrong MoveCategory with ID returning no error")
	}
}

func TestMoveCategory_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveCategory("wrong")
	if err == nil {
		t.Error("Fetching wrong MoveCategory with Name returning no error")
	}
}

func TestMoveCategory_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveCategory("")
	if err == nil {
		t.Error("Fetching wrong MoveCategory with empty ID returning no error")
	}
}

func TestMoveDamageClasses(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveDamageClasses("0", "20")
	if err != nil {
		t.Error("Fetching MoveDamangeClasses returning error")
	}
	var emptyMoveDamageClasses models.MoveDamageClasses
	if reflect.DeepEqual(emptyMoveDamageClasses, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestMoveDamageClass_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveDamageClass("1")
	if err != nil {
		t.Error("Fetching MoveDamageClass with ID returning error")
	}
	var emptyMoveDamageClass models.MoveDamageClass
	if reflect.DeepEqual(emptyMoveDamageClass, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "status" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveDamageClass_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveDamageClass("status")
	if err != nil {
		t.Error("Fetching MoveDamageClass with ID returning error")
	}
	var emptyMoveDamageClass models.MoveDamageClass
	if reflect.DeepEqual(emptyMoveDamageClass, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "status" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveDamageClass_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveDamageClass("0")
	if err == nil {
		t.Error("Fetching wrong MoveDamageClass with ID returning no error")
	}
}

func TestMoveDamageClass_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveDamageClass("wrong")
	if err == nil {
		t.Error("Fetching wrong MoveDamageClass with Name returning no error")
	}
}

func TestMoveDamageClass_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveDamageClass("")
	if err == nil {
		t.Error("Fetching wrong MoveDamageClass with empty ID returning no error")
	}
}

func TestMoveLearnMethods(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveLearnMethods("0", "20")
	if err != nil {
		t.Error("Fetching MoveLearnMethods returning error")
	}
	var emptyMoveLearnMethods models.MoveLearnMethods
	if reflect.DeepEqual(emptyMoveLearnMethods, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestMoveLearnMethod_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveLearnMethod("1")
	if err != nil {
		t.Error("Fetching MoveLearnMethod with ID returning error")
	}
	var emptyMoveLearnMethod models.MoveLearnMethod
	if reflect.DeepEqual(emptyMoveLearnMethod, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "level-up" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveLearnMethod_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveLearnMethod("level-up")
	if err != nil {
		t.Error("Fetching MoveLearnMethod with ID returning error")
	}
	var emptyMoveLearnMethod models.MoveLearnMethod
	if reflect.DeepEqual(emptyMoveLearnMethod, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "level-up" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveLearnMethod_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveLearnMethod("0")
	if err == nil {
		t.Error("Fetching wrong MoveLearnMethod with ID returning no error")
	}
}

func TestMoveLearnMethod_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveLearnMethod("wrong")
	if err == nil {
		t.Error("Fetching wrong MoveLearnMethod with Name returning no error")
	}
}

func TestMoveLearnMethod_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveLearnMethod("")
	if err == nil {
		t.Error("Fetching wrong MoveLearnMethod with empty ID returning no error")
	}
}

func TestMoveTargets(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveTargets("0", "20")
	if err != nil {
		t.Error("Fetching MoveTargets returning error")
	}
	var emptyMoveTargets models.MoveTargets
	if reflect.DeepEqual(emptyMoveTargets, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestMoveTarget_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveTarget("1")
	if err != nil {
		t.Error("Fetching MoveTarget with ID returning error")
	}
	var emptyMoveTarget models.MoveTarget
	if reflect.DeepEqual(emptyMoveTarget, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "specific-move" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveTarget_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.MoveTarget("specific-move")
	if err != nil {
		t.Error("Fetching MoveTarget with ID returning error")
	}
	var emptyMoveTarget models.MoveTarget
	if reflect.DeepEqual(emptyMoveTarget, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "specific-move" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestMoveTarget_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveTarget("0")
	if err == nil {
		t.Error("Fetching wrong MoveTarget with ID returning no error")
	}
}

func TestMoveTarget_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveTarget("wrong")
	if err == nil {
		t.Error("Fetching wrong MoveTarget with Name returning no error")
	}
}

func TestMoveTarget_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.MoveTarget("")
	if err == nil {
		t.Error("Fetching wrong MoveTarget with empty ID returning no error")
	}
}
