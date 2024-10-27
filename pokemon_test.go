package pokego_test

import (
	"reflect"
	"testing"

	"github.com/hanle23/pokego"
	"github.com/hanle23/pokego/internal/models"
)

func TestAbilities(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Abilities("0", "20")
	if err != nil {
		t.Error("Fetching Abilities returning error")
	}
	var emptyAbilities models.Abilities
	if reflect.DeepEqual(emptyAbilities, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestAbility_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Ability("1")
	if err != nil {
		t.Error("Fetching Ability with ID returning error")
	}
	var emptyAbility models.Ability
	if reflect.DeepEqual(emptyAbility, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "stench" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestAbility_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Ability("stench")
	if err != nil {
		t.Error("Fetching Ability with Name returning error")
	}
	var emptyAbility models.Ability
	if reflect.DeepEqual(emptyAbility, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "stench" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestAbility_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Ability("0")
	if err == nil {
		t.Error("Fetching wrong Ability with ID returning no error")
	}
}

func TestAbility_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Ability("wrong")
	if err == nil {
		t.Error("Fetching wrong Ability with Name returning no error")
	}
}

func TestAbility_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Ability("")
	if err == nil {
		t.Error("Fetching empty Ability with ID returning no error")
	}
}

func TestCharacteristics(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Characteristics("0", "20")
	if err != nil {
		t.Error("Fetching Characteristics returning error")
	}
	var emptyCharacteristics models.Characteristics
	if reflect.DeepEqual(emptyCharacteristics, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestCharacteristic_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Characteristic("1")
	if err != nil {
		t.Error("Fetching Characteristic with ID returning error")
	}
	var emptyCharacteristic models.Characteristic
	if reflect.DeepEqual(emptyCharacteristic, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.GeneModulo != 0 {
		t.Error("Fetch result returns wrong GeneModulo")
	}
}

func TestCharacteristic_FailWithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Characteristic("wrong")
	if err == nil {
		t.Error("Fetching wrong Characteristic with Name returning no error")
	}
}

func TestCharacteristic_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Characteristic("")
	if err == nil {
		t.Error("Fetching empty Characteristic with ID returning no error")
	}
}

func TestEggGroups(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EggGroups("0", "20")
	if err != nil {
		t.Error("Fetching EggGroups returning error")
	}
	var emptyEggGroups models.EggGroups
	if reflect.DeepEqual(emptyEggGroups, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestEggGroup_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EggGroup("1")
	if err != nil {
		t.Error("Fetching EggGroup with ID returning error")
	}
	var emptyEggGroup models.EggGroup
	if reflect.DeepEqual(emptyEggGroup, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "monster" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestEggGroup_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.EggGroup("monster")
	if err != nil {
		t.Error("Fetching EggGroup with Name returning error")
	}
	var emptyEggGroup models.EggGroup
	if reflect.DeepEqual(emptyEggGroup, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "monster" {
		t.Error("Fetch result returns wrong name")
	}
}

func TestEggGroup_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EggGroup("0")
	if err == nil {
		t.Error("Fetching wrong EggGroup with ID returning no error")
	}
}

func TestEggGroup_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EggGroup("wrong")
	if err == nil {
		t.Error("Fetching wrong EggGroup with Name returning no error")
	}
}

func TestEggGroup_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.EggGroup("")
	if err == nil {
		t.Error("Fetching empty EggGroup with ID returning no error")
	}
}

func TestGenders(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Genders("0", "20")
	if err != nil {
		t.Error("Fetching EggGroups returning error")
	}
	var emptyGenders models.Genders
	if reflect.DeepEqual(emptyGenders, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestGender_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Gender("1")
	if err != nil {
		t.Error("Fetching Gender with ID returning error")
	}
	var emptyGender models.Gender
	if reflect.DeepEqual(emptyGender, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "female" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestGender_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Gender("female")
	if err != nil {
		t.Error("Fetching Gender with ID returning error")
	}
	var emptyGender models.Gender
	if reflect.DeepEqual(emptyGender, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "female" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestGender_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Gender("0")
	if err == nil {
		t.Error("Fetching wrong Gender with invalid ID returning no error")
	}
}

func TestGender_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Gender("wrong")
	if err == nil {
		t.Error("Fetching wrong Gender with invalid ID returning no error")
	}
}

func TestGender_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Gender("")
	if err == nil {
		t.Error("Fetching empty Gender with invalid ID returning no error")
	}
}

func TestGrowthRates(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.GrowthRates("0", "20")
	if err != nil {
		t.Error("Fetching GrowthRates returning error")
	}
	var emptyGrowthRates models.GrowthRates
	if reflect.DeepEqual(emptyGrowthRates, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestGrowthRate_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.GrowthRate("1")
	if err != nil {
		t.Error("Fetching GrowthRate with ID returning error")
	}
	var emptyGrowthRate models.GrowthRate
	if reflect.DeepEqual(emptyGrowthRate, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "slow" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestGrowthRate_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.GrowthRate("slow")
	if err != nil {
		t.Error("Fetching GrowthRate with Name returning error")
	}
	var emptyGrowthRate models.GrowthRate
	if reflect.DeepEqual(emptyGrowthRate, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "slow" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestGrowthRate_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.GrowthRate("0")
	if err == nil {
		t.Error("Fetching wrong GrowthRate with invalid ID returning no error")
	}
}

func TestGrowthRate_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.GrowthRate("wrong")
	if err == nil {
		t.Error("Fetching wrong GrowthRate with invalid Name returning no error")
	}
}

func TestGrowthRate_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.GrowthRate("")
	if err == nil {
		t.Error("Fetching empty GrowthRate with invalid ID returning no error")
	}
}

func TestNatures(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Natures("0", "20")
	if err != nil {
		t.Error("Fetching Natures returning error")
	}
	var emptyNatures models.Natures
	if reflect.DeepEqual(emptyNatures, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestNature_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Nature("1")
	if err != nil {
		t.Error("Fetching Nature with ID returning error")
	}
	var emptyNature models.Nature
	if reflect.DeepEqual(emptyNature, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "hardy" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestNature_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Nature("hardy")
	if err != nil {
		t.Error("Fetching Nature with Name returning error")
	}
	var emptyNature models.Nature
	if reflect.DeepEqual(emptyNature, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "hardy" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestNature_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Nature("0")
	if err == nil {
		t.Error("Fetching wrong Nature with invalid ID returning no error")
	}
}

func TestNature_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Nature("wrong")
	if err == nil {
		t.Error("Fetching wrong Nature with invalid Name returning no error")
	}
}

func TestNature_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Nature("")
	if err == nil {
		t.Error("Fetching empty Nature with invalid ID returning no error")
	}
}

func TestPokeathlonStats(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokeathlonStats("0", "20")
	if err != nil {
		t.Error("Fetching PokeathlonStats returning error")
	}
	var emptyPokeathlonStats models.PokeathlonStats
	if reflect.DeepEqual(emptyPokeathlonStats, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPokeathlonStat_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokeathlonStat("1")
	if err != nil {
		t.Error("Fetching PokeathlonStat with ID returning error")
	}
	var emptyPokeathlonStat models.PokeathlonStat
	if reflect.DeepEqual(emptyPokeathlonStat, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "speed" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestPokeathlonStat_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokeathlonStat("speed")
	if err != nil {
		t.Error("Fetching PokeathlonStat with Name returning error")
	}
	var emptyPokeathlonStat models.PokeathlonStat
	if reflect.DeepEqual(emptyPokeathlonStat, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "speed" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestPokeathlonStat_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokeathlonStat("0")
	if err == nil {
		t.Error("Fetching wrong PokeathlonStat with invalid ID returning no error")
	}
}

func TestPokeathlonStat_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokeathlonStat("wrong")
	if err == nil {
		t.Error("Fetching wrong PokeathlonStat with invalid Name returning no error")
	}
}

func TestPokeathlonStat_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokeathlonStat("")
	if err == nil {
		t.Error("Fetching empty PokeathlonStat with invalid ID returning no error")
	}
}

func TestPokemons(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Pokemons("0", "20")
	if err != nil {
		t.Error("Fetching Pokemon returning error")
	}
	var emptyPokemons models.Pokemons
	if reflect.DeepEqual(emptyPokemons, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPokemon_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Pokemon("1")
	if err != nil {
		t.Errorf("Fetching Pokemon with ID returning error: %v", err)
	}
	var emptyPokemon models.Pokemon
	if reflect.DeepEqual(emptyPokemon, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "bulbasaur" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestPokemon_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Pokemon("bulbasaur")
	if err != nil {
		t.Error("Fetching Pokemon with Name returning error")
	}
	var emptyPokemon models.Pokemon
	if reflect.DeepEqual(emptyPokemon, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "bulbasaur" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestPokemon_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Pokemon("0")
	if err == nil {
		t.Error("Fetching wrong Pokemon with invalid ID returning no error")
	}
}

func TestPokemon_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Pokemon("wrong")
	if err == nil {
		t.Error("Fetching wrong Pokemon with invalid Name returning no error")
	}
}

func TestPokemon_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Pokemon("")
	if err == nil {
		t.Error("Fetching empty Pokemon with invalid ID returning no error")
	}
}

func TestPokemonLocationAreas_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Pokemon("1")
	if err != nil {
		t.Errorf("Fetching Pokemon with ID returning error: %v", err)
	}
	result, err := client.PokemonLocationAreas("1")
	if err != nil {
		t.Error("Fetching PokemonLocationAreas with ID returning error")
	}
	var emptyPokemonLocationAreas []models.LocationAreaEncounter
	if reflect.DeepEqual(emptyPokemonLocationAreas, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPokemonLocationAreas_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Pokemon("bulbasaur")
	if err != nil {
		t.Error("Fetching Pokemon with Name returning error")
	}
	result, err := client.PokemonLocationAreas("bulbasaur")
	if err != nil {
		t.Error("Fetching PokemonLocationAreas with Name returning error")
	}
	var emptyPokemonLocationAreas []models.LocationAreaEncounter
	if reflect.DeepEqual(emptyPokemonLocationAreas, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPokemonLocationAreas_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonLocationAreas("0")
	if err == nil {
		t.Error("Fetching wrong PokemonLocationAreas with invalid ID returning no error")
	}
}

func TestPokemonLocationAreas_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonLocationAreas("wrong")
	if err == nil {
		t.Error("Fetching wrong PokemonLocationAreas with invalid Name returning no error")
	}
}

func TestPokemonColors(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonColors("0", "20")
	if err != nil {
		t.Error("Fetching Pokemon Colors returning error")
	}
	var emptyPokemonColors models.PokemonColors
	if reflect.DeepEqual(emptyPokemonColors, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPokemonColor_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonColor("1")
	if err != nil {
		t.Error("Fetching PokemonColor with ID returning error")
	}
	var emptyPokemonColor models.PokemonColor
	if reflect.DeepEqual(emptyPokemonColor, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "black" {
		t.Error("Fetch result returns wrong IsBaby")
	}
}

func TestPokemonColor_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonColor("black")
	if err != nil {
		t.Error("Fetching PokemonColor with ID returning error")
	}
	var emptyPokemonColor models.PokemonColor
	if reflect.DeepEqual(emptyPokemonColor, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "black" {
		t.Error("Fetch result returns wrong IsBaby")
	}
}

func TestPokemonColor_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonColor("0")
	if err == nil {
		t.Error("Fetching wrong PokemonColor with invalid ID returning no error")
	}
}

func TestPokemonColor_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonColor("wrong")
	if err == nil {
		t.Error("Fetching wrong PokemonColor with invalid Name returning no error")
	}
}

func TestPokemonColor_FailedWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonColor("")
	if err == nil {
		t.Error("Fetching wrong PokemonColor with invalid Name returning no error")
	}
}

func TestPokemonForms(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonForms("0", "20")
	if err != nil {
		t.Error("Fetching PokemonForms returning error")
	}
	var emptyPokemonForms models.PokemonForms
	if reflect.DeepEqual(emptyPokemonForms, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPokemonForm_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonForm("1")
	if err != nil {
		t.Error("Fetching PokemonForm with ID returning error")
	}
	var emptyPokemonForm models.PokemonForm
	if reflect.DeepEqual(emptyPokemonForm, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "bulbasaur" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestPokemonForm_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonForm("bulbasaur")
	if err != nil {
		t.Error("Fetching PokemonForm with ID returning error")
	}
	var emptyPokemonForm models.PokemonForm
	if reflect.DeepEqual(emptyPokemonForm, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "bulbasaur" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestPokemonForm_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonForm("0")
	if err == nil {
		t.Error("Fetching wrong PokemonForm with invalid ID returning no error")
	}
}

func TestPokemonForm_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonForm("wrong")
	if err == nil {
		t.Error("Fetching wrong PokemonForm with invalid Name returning no error")
	}
}

func TestPokemonForm_FailedWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonForm("")
	if err == nil {
		t.Error("Fetching wrong PokemonForm with invalid Name returning no error")
	}
}

func TestPokemonHabitats(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonHabitats("0", "20")
	if err != nil {
		t.Error("Fetching PokemonHabitats returning error")
	}
	var emptyPokemonHabitats models.PokemonHabitats
	if reflect.DeepEqual(emptyPokemonHabitats, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPokemonHabitat_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonHabitat("1")
	if err != nil {
		t.Error("Fetching PokemonHabitat with ID returning error")
	}
	var emptyPokemonHabitat models.PokemonHabitat
	if reflect.DeepEqual(emptyPokemonHabitat, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "cave" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestPokemonHabitat_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonHabitat("cave")
	if err != nil {
		t.Error("Fetching PokemonHabitat with ID returning error")
	}
	var emptyPokemonHabitat models.PokemonHabitat
	if reflect.DeepEqual(emptyPokemonHabitat, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "cave" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestPokemonHabitat_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonHabitat("0")
	if err == nil {
		t.Error("Fetching wrong PokemonHabitat with invalid ID returning no error")
	}
}

func TestPokemonHabitat_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonHabitat("wrong")
	if err == nil {
		t.Error("Fetching wrong PokemonHabitat with invalid Name returning no error")
	}
}

func TestPokemonHabitat_FailedWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonHabitat("")
	if err == nil {
		t.Error("Fetching wrong PokemonHabitat with invalid Name returning no error")
	}
}

func TestPokemonShapes(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonShapes("0", "20")
	if err != nil {
		t.Error("Fetching PokemonShapes returning error")
	}
	var emptyPokemonShapes models.PokemonShapes
	if reflect.DeepEqual(emptyPokemonShapes, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPokemonShape_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonShape("1")
	if err != nil {
		t.Error("Fetching PokemonShape with ID returning error")
	}
	var emptyPokemonShape models.PokemonShape
	if reflect.DeepEqual(emptyPokemonShape, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "ball" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestPokemonShape_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonShape("ball")
	if err != nil {
		t.Error("Fetching PokemonShape with ID returning error")
	}
	var emptyPokemonShape models.PokemonShape
	if reflect.DeepEqual(emptyPokemonShape, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "ball" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestPokemonShape_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonShape("0")
	if err == nil {
		t.Error("Fetching wrong PokemonShape with invalid ID returning no error")
	}
}

func TestPokemonShape_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonShape("wrong")
	if err == nil {
		t.Error("Fetching wrong PokemonShape with invalid ID returning no error")
	}
}

func TestPokemonShape_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonShape("")
	if err == nil {
		t.Error("Fetching wrong PokemonShape with invalid ID returning no error")
	}
}

func TestPokemonSpecies(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonSpeciesPagination("0", "20")
	if err != nil {
		t.Error("Fetching PokemonSpeciesPagination returning error")
	}
	var emptyPokemonSpeciesPagination models.PokemonSpeciesPagination
	if reflect.DeepEqual(emptyPokemonSpeciesPagination, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestPokemonSpecies_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonSpecies("1")
	if err != nil {
		t.Error("Fetching PokemonSpecies with ID returning error")
	}
	var emptyPokemonSpecies models.PokemonSpecies
	if reflect.DeepEqual(emptyPokemonSpecies, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "bulbasaur" {
		t.Error("Fetch result returns wrong IsBaby")
	}
}

func TestPokemonSpecies_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.PokemonSpecies("bulbasaur")
	if err != nil {
		t.Error("Fetching PokemonSpecies with ID returning error")
	}
	var emptyPokemonSpecies models.PokemonSpecies
	if reflect.DeepEqual(emptyPokemonSpecies, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "bulbasaur" {
		t.Error("Fetch result returns wrong IsBaby")
	}
}

func TestPokemonSpecies_FailWithWrongID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonSpecies("0")
	if err == nil {
		t.Error("Fetching wrong PokemonSpecies with invalid ID returning no error")
	}
}

func TestPokemonSpecies_FailWithWrongName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonSpecies("wrong")
	if err == nil {
		t.Error("Fetching wrong PokemonSpecies with invalid ID returning no error")
	}
}
func TestPokemonSpecies_FailWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.PokemonSpecies("")
	if err == nil {
		t.Error("Fetching wrong PokemonSpecies with invalid ID returning no error")
	}
}

func TestStats(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Stats("0", "20")
	if err != nil {
		t.Error("Fetching Stats returning error")
	}
	var emptyStats models.Stats
	if reflect.DeepEqual(emptyStats, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestStat_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Stat("1")
	if err != nil {
		t.Error("Fetching Stat with ID returning error")
	}
	var emptyStat models.Stat
	if reflect.DeepEqual(emptyStat, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "hp" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestStat_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Stat("hp")
	if err != nil {
		t.Error("Fetching Stat with Name returning error")
	}
	var emptyStat models.Stat
	if reflect.DeepEqual(emptyStat, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "hp" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestStat_FailedWithInvalidId(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Stat("0")
	if err == nil {
		t.Error("Fetching wrong Stat with invalid ID returning no error")
	}
}

func TestStat_FailedWithInvalidName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Stat("wrong")
	if err == nil {
		t.Error("Fetching wrong Stat with invalid Name returning no error")
	}
}

func TestStat_FailedWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Stat("")
	if err == nil {
		t.Error("Fetching wrong Stat with empty Name returning no error")
	}
}

func TestTypes(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Types("0", "20")
	if err != nil {
		t.Error("Fetching Types returning error")
	}
	var emptyTypes models.Types
	if reflect.DeepEqual(emptyTypes, result) {
		t.Error("Fetch result returns nil")
	}
}

func TestType_WithID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Type("1")
	if err != nil {
		t.Error("Fetching Type with ID returning error")
	}
	var emptyType models.Type
	if reflect.DeepEqual(emptyType, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "normal" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestType_WithName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	result, err := client.Type("normal")
	if err != nil {
		t.Error("Fetching Type with ID returning error")
	}
	var emptyType models.Type
	if reflect.DeepEqual(emptyType, result) {
		t.Error("Fetch result returns nil")
	}
	if result.ID != 1 {
		t.Error("Fetch result returns wrong ID")
	}
	if result.Name != "normal" {
		t.Error("Fetch result returns wrong Name")
	}
}

func TestType_FailedWithInvalidID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Type("0")
	if err == nil {
		t.Error("Fetching wrong Type with invalid ID returning no error")
	}
}

func TestType_FailedWithInvalidName(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Type("wrong")
	if err == nil {
		t.Error("Fetching wrong Type with invalid Name returning no error")
	}
}

func TestType_FailedWithEmptyID(t *testing.T) {
	client := pokego.NewClient()
	if client == nil {
		t.Error("NewClient() returned nil")
	}
	_, err := client.Type("")
	if err == nil {
		t.Error("Fetching wrong Type with empty ID returning no error")
	}
}
