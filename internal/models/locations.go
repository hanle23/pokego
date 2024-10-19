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

type Region struct {
	ID             int            `json:"id"`
	Locations      []Location     `json:"locations"`
	Name           string         `json:"name"`
	Names          []Name         `json:"names"`
	MainGeneration Generation     `json:"main_generation"`
	Pokedexes      []Pokedex      `json:"pokedexes"`
	VersionGroups  []VersionGroup `json:"version_groups"`
}
