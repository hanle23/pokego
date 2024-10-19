package models

type Ability struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	IsMainSeries      bool   `json:"is_main_series"`
	Generation        `json:"generation"`
	Names             []Name                `json:"names"`
	EffectEntries     []Effect              `json:"effect_entries"`
	EffectChanges     []AbilityEffectChange `json:"effect_changes"`
	FlavorTextEntries []AbilityFlavorText   `json:"flavor_text_entries"`
	Pokemon           []AbilityPokemon      `json:"pokemon"`
}

type AbilityEffectChange struct {
	EffectEntries []Effect `json:"effect_entries"`
	VersionGroup  `json:"version_group"`
}

type AbilityFlavorText struct {
	FlavorText   string `json:"flavor_text"`
	Language     `json:"language"`
	VersionGroup `json:"version_group"`
}

type AbilityPokemon struct {
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
	Pokemon  `json:"pokemon"`
}

type Characteristic struct {
	ID             int           `json:"id"`
	GeneModulo     int           `json:"gene_modulo"`
	PossibleValues []int         `json:"possible_values"`
	HighestStat    Stat          `json:"highest_stat"`
	Descriptions   []Description `json:"descriptions"`
}

type EggGroup struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	Names          []Name           `json:"names"`
	PokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}

type Gender struct {
	ID                    int                    `json:"id"`
	Name                  string                 `json:"name"`
	PokemonSpeciesDetails []PokemonSpeciesGender `json:"pokemon_species_details"`
	RequiredForEvolution  []PokemonSpecies       `json:"required_for_evolution"`
}

type PokemonSpeciesGender struct {
	Rate           int `json:"rate"`
	PokemonSpecies `json:"pokemon_species"`
}

type GrowthRate struct {
	ID             int                         `json:"id"`
	Name           string                      `json:"name"`
	Formula        string                      `json:"formula"`
	Descriptions   []Description               `json:"descriptions"`
	Levels         []GrowthRateExperienceLevel `json:"levels"`
	PokemonSpecies []PokemonSpecies            `json:"pokemon_species"`
}

type GrowthRateExperienceLevel struct {
	Level      int `json:"level"`
	Experience int `json:"experience"`
}

type Nature struct {
	ID                         int                         `json:"id"`
	Name                       string                      `json:"name"`
	DecreasedStat              Stat                        `json:"decreased_stat"`
	IncreasedStat              Stat                        `json:"increased_stat"`
	HatesFlavor                BerryFlavor                 `json:"hates_flavor"`
	LikesFlavor                BerryFlavor                 `json:"likes_flavor"`
	PokeathlonStatChanges      []NatureStatChange          `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []MoveBattleStylePreference `json:"move_battle_style_preferences"`
	Names                      []Name                      `json:"names"`
}

type NatureStatChange struct {
	MaxChange      int `json:"max_change"`
	PokeathlonStat `json:"pokeathlon_stat"`
}

type MoveBattleStylePreference struct {
	LowHpPreference  int `json:"low_hp_preference"`
	HighHpPreference int `json:"high_hp_preference"`
	MoveBattleStyle  `json:"move_battle_style"`
}

type PokeathlonStat struct {
	ID               int                            `json:"id"`
	Name             string                         `json:"name"`
	Names            []Name                         `json:"names"`
	AffectingNatures NaturePokeathlonStatAffectSets `json:"affecting_natures"`
}

type NaturePokeathlonStatAffectSets struct {
	Increase []NaturePokeathlonStatAffect `json:"increase"`
	Decrease []NaturePokeathlonStatAffect `json:"decrease"`
}

type NaturePokeathlonStatAffect struct {
	MaxChange int `json:"max_change"`
	Nature    `json:"nature"`
}

type Pokemon struct {
	ID                     int                `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  bool               `json:"order"`
	Weight                 int                `json:"weight"`
	Abilities              []PokemonAbility   `json:"abilities"`
	Forms                  []PokemonForm      `json:"forms"`
	GameIndices            []VersionGameIndex `json:"game_indices"`
	HeldItems              []PokemonHeldItem  `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []PokemonMove      `json:"moves"`
	PastTypes              []PokemonTypePast  `json:"past_types"`
	Sprites                PokemonSprites     `json:"sprites"`
	Cries                  PokemonCries       `json:"cries"`
	Species                PokemonSpecies     `json:"species"`
	Stats                  []PokemonStat      `json:"stats"`
	Types                  []PokemonType      `json:"types"`
}

type PokemonAbility struct {
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
	Ability  `json:"ability"`
}

type PokemonType struct {
	Slot int `json:"slot"`
	Type `json:"type"`
}

type PokemonFormType struct {
	Slot int `json:"slot"`
	Type `json:"type"`
}

type PokemonTypePast struct {
	Generation `json:"generation"`
	Types      []PokemonType `json:"type"`
}

type PokemonHeldItem struct {
	Item           `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonMove struct {
	Move                `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod `json:"move_learn_method"`
	VersionGroup    `json:"version_group"`
	LevelLearnedAt  int `json:"level_learned_at"`
}

type PokemonStat struct {
	Stat     `json:"stat"`
	Effort   int `json:"effort"`
	BaseStat int `json:"base_stat"`
}

type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type LocationAreaEncounter struct {
	LocationArea   `json:"location_area"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type PokemonColor struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	Names          []Name           `json:"names"`
	PokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}

type PokemonForm struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Order        int    `json:"order"`
	FormOrder    int    `json:"form_order"`
	IsDefault    bool   `json:"is_default"`
	IsBattleOnly bool   `json:"is_battle_only"`
	IsMega       bool   `json:"is_mega"`
	FormName     string `json:"form_name"`
	Pokemon      `json:"pokemon"`
	Types        []PokemonFormType  `json:"types"`
	Sprites      PokemonFormSprites `json:"sprites"`
	VersionGroup `json:"version_group"`
	Names        []Name `json:"names"`
	FormNames    []Name `json:"form_names"`
}

type PokemonFormSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
}

type PokemonHabitat struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	Names          []Name           `json:"names"`
	PokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}

type PokemonShape struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	AwesomeNames   []AwesomeName    `json:"awesome_names"`
	Names          []Name           `json:"names"`
	PokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}

type AwesomeName struct {
	AwesomeName string `json:"awesome_name"`
	Language    `json:"language"`
}

type PokemonSpecies struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Order                int    `json:"order"`
	GenderRate           int    `json:"gender_rate"`
	CaptureRate          int    `json:"capture_rate"`
	BaseHappiness        int    `json:"base_happiness"`
	IsBaby               bool   `json:"is_baby"`
	IsLegendary          bool   `json:"is_legendary"`
	IsMythical           bool   `json:"is_mythical"`
	HatchCounter         int    `json:"hatch_counter"`
	HasGenderDifferences bool   `json:"has_gender_differences"`
	FormsSwitchable      bool   `json:"forms_switchable"`
	GrowthRate           `json:"growth_rate"`
	PokedexNumbers       []PokemonSpeciesDexEntry `json:"pokedex_numbers"`
	EggGroups            []EggGroup               `json:"egg_groups"`
	Color                PokemonColor             `json:"color"`
	Shape                PokemonShape             `json:"shape"`
	EvolvesFromSpecies   *PokemonSpecies          `json:"evolves_from_species"`
	EvolutionChain       `json:"evolution_chain"`
	Habitat              PokemonHabitat `json:"habitat"`
	Generation           `json:"generation"`
	Names                []Name                  `json:"names"`
	PalParkEncounters    []PalParkEncounterArea  `json:"pal_park_encounters"`
	FlavorTextEntries    []FlavorText            `json:"flavor_text_entries"`
	FormDescriptions     []Description           `json:"form_descriptions"`
	Genera               []Genus                 `json:"genera"`
	Varieties            []PokemonSpeciesVariety `json:"varieties"`
}

type Genus struct {
	Genus    string `json:"genus"`
	Language `json:"language"`
}

type PokemonSpeciesDexEntry struct {
	EntryNumber int `json:"entry_number"`
	Pokedex     `json:"pokedex"`
}

type PalParkEncounterArea struct {
	BaseScore int         `json:"base_score"`
	Rate      int         `json:"rate"`
	Area      PalParkArea `json:"area"`
}

type PokemonSpeciesVariety struct {
	IsDefault bool `json:"is_default"`
	Pokemon   `json:"pokemon"`
}

type Stat struct {
	ID               int                  `json:"id"`
	Name             string               `json:"name"`
	GameIndex        int                  `json:"game_index"`
	IsBattleOnly     bool                 `json:"is_battle_only"`
	AffectingMoves   MoveStatAffectSets   `json:"affecting_moves"`
	AffectingNatures NatureStatAffectSets `json:"affecting_natures"`
	Characteristics  `json:"characteristics"`
	MoveDamageClass  `json:"move_damage_class"`
	Names            []Name `json:"names"`
}

type MoveStatAffectSets struct {
	Increase []MoveStatAffect `json:"increase"`
	Decrease []MoveStatAffect `json:"decrease"`
}

type MoveStatAffect struct {
	Change int `json:"change"`
	Move   `json:"move"`
}

type NatureStatAffectSets struct {
	Increase []Nature `json:"increase"`
	Decrease []Nature `json:"decrease"`
}

type Type struct {
	ID                  int                   `json:"id"`
	Name                string                `json:"name"`
	DamageRelations     TypeRelations         `json:"damage_relations"`
	PastDamageRelations []Type                `json:"past_damage_relations"`
	GameIndices         []GenerationGameIndex `json:"game_indices"`
	Generation          `json:"generation"`
	MoveDamageClass     `json:"move_damage_class"`
	Names               []Name        `json:"names"`
	Pokemon             []TypePokemon `json:"pokemon"`
	Moves               []Move        `json:"moves"`
}

type TypePokemon struct {
	Slot    int `json:"slot"`
	Pokemon `json:"pokemon"`
}

type TypeRelations struct {
	NoDamageTo       []Type `json:"no_damage_to"`
	HalfDamageTo     []Type `json:"half_damage_to"`
	DoubleDamageTo   []Type `json:"double_damage_to"`
	NoDamageFrom     []Type `json:"no_damage_from"`
	HalfDamageFrom   []Type `json:"half_damage_from"`
	DoubleDamageFrom []Type `json:"double_damage_from"`
}

type TypeRelationsPast struct {
	Generation      `json:"generation"`
	DamageRelations TypeRelations `json:"damage_relations"`
}
