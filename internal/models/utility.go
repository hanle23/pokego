package models

type Language struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	Iso639   string `json:"iso639"`
	Iso3166  string `json:"iso3166"`
	Names    []Name `json:"names"`
}

type APIResource struct {
	URL string `json:"url"`
}

type Description struct {
	Description string `json:"description"`
	Language    `json:"language"`
}

type Effect struct {
	Effect   string `json:"effect"`
	Language `json:"language"`
}

type Encounter struct {
	MinLevel        int                       `json:"min_level"`
	MaxLevel        int                       `json:"max_level"`
	ConditionValues []EncounterConditionValue `json:"condition_values"`
	Chance          int                       `json:"chance"`
	Method          EncounterMethod           `json:"method"`
}

type FlavorText struct {
	FlavorText string `json:"flavor_text"`
	Language   `json:"language"`
	Version    `json:"version"`
}

type GenerationGameIndex struct {
	GameIndex  int `json:"game_index"`
	Generation `json:"generation"`
}

type MachineVersionDetail struct {
	Machine      `json:"machine"`
	VersionGroup `json:"version_group"`
}

type Name struct {
	Name     string `json:"name"`
	Language `json:"language"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VerboseEffect struct {
	Effect      string `json:"effect"`
	ShortEffect string `json:"short_effect"`
	Language    `json:"language"`
}

type VersionEncounterDetail struct {
	Version          `json:"version"`
	MaxChance        int         `json:"max_chance"`
	EncounterDetails []Encounter `json:"encounter_details"`
}

type VersionGameIndex struct {
	GameIndex int `json:"game_index"`
	Version   `json:"version"`
}

type VersionGroupFlavorText struct {
	Text         string `json:"text"`
	Language     `json:"language"`
	VersionGroup `json:"version_group"`
}
