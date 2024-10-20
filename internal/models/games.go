package models

type Generation struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	Abilities      []Ability        `json:"abilities"`
	Names          []Name           `json:"names"`
	MainRegion     *Region          `json:"main_region"`
	Moves          []Move           `json:"moves"`
	PokemonSpecies []PokemonSpecies `json:"pokemon_species"`
	Types          []Type           `json:"types"`
	VersionGroups  []VersionGroup   `json:"version_groups"`
}

type Pokedex struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	IsMainSeries   bool           `json:"is_main_series"`
	Descriptions   []Description  `json:"descriptions"`
	Names          []Name         `json:"names"`
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
	Region         []Region       `json:"region"`
	VersionGroups  []VersionGroup `json:"version_groups"`
}

type PokemonEntry struct {
	EntryNumber    int `json:"entry_number"`
	PokemonSpecies `json:"pokemon_species"`
}

type Version struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Names        []Name `json:"names"`
	VersionGroup `json:"version_group"`
}

type VersionGroup struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Order            int    `json:"order"`
	Generation       `json:"generation"`
	MoveLearnMethods []MoveLearnMethod `json:"move_learn_methods"`
	Pokedexes        []Pokedex         `json:"pokedexes"`
	Regions          []Region          `json:"regions"`
	Versions         []Version         `json:"versions"`
}
