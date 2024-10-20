package models

type Location struct {
	ID             int            `json:"id"`
	Locations      []Location     `json:"locations"`
	Name           string         `json:"name"`
	Names          []Name         `json:"names"`
	MainGeneration Generation     `json:"main_generation"`
	Pokedexes      []Pokedex      `json:"pokedexes"`
	VersionGroups  []VersionGroup `json:"version_groups"`
}

type LocationArea struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             Location              `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod EncounterMethod          `json:"encounter_method"`
	VersionDetails  []VersionEncounterDetail `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type PalParkArea struct {
	ID                int                       `json:"id"`
	Name              string                    `json:"name"`
	Names             []Name                    `json:"names"`
	PokemonEncounters []PalParkEncounterSpecies `json:"pokemon_encounters"`
}

type PalParkEncounterSpecies struct {
	BaseScore      int `json:"base_score"`
	Rate           int `json:"rate"`
	PokemonSpecies `json:"pokemon_species"`
}

type Region struct {
	ID             int            `json:"id"`
	Locations      []Location     `json:"locations"`
	Name           string         `json:"name"`
	Names          []Name         `json:"names"`
	MainGeneration Generation     `json:"main_generation"`
	Pokedexes      []Pokedex      `json:"pokedexes"`
	VersionGroups  []VersionGroup `json:"version_groups"`
}
