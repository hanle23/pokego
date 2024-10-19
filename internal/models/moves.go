package models

type Move struct {
	ID                 int              `json:"id"`
	Name               string           `json:"name"`
	Accuracy           int              `json:"accuracy"`
	EffectChance       int              `json:"effect_chance"`
	PP                 int              `json:"pp"`
	Priority           int              `json:"priority"`
	Power              int              `json:"power"`
	ContestCombos      ContestComboSets `json:"contest_combos"`
	ContestType        `json:"contest_type"`
	ContestEffect      `json:"contest_effect"`
	DamageClass        MoveDamageClass       `json:"damage_class"`
	EffectEntries      []VerboseEffect       `json:"effect_entries"`
	EffectChanges      []AbilityEffectChange `json:"effect_changes"`
	LearnedByPokemon   []Pokemon             `json:"learned_by_pokemon"`
	FlavorTextEntries  []MoveFlavorText      `json:"flavor_text_entries"`
	Generation         `json:"generation"`
	Machines           []MachineVersionDetail `json:"machines"`
	Meta               MoveMetaData           `json:"meta"`
	Names              []Name                 `json:"names"`
	PastValues         []PastMoveStatValues   `json:"past_values"`
	StatChanges        []MoveStatChange       `json:"stat_changes"`
	SuperContestEffect `json:"super_contest_effect"`
	Target             MoveTarget `json:"target"`
	Type               `json:"type"`
}

type ContestComboSets struct {
	Normal ContestComboDetail `json:"normal"`
	Super  ContestComboDetail `json:"super"`
}

type ContestComboDetail struct {
	UseBefore []Move `json:"use_before"`
	UseAfter  []Move `json:"use_after"`
}

type MoveFlavorText struct {
	FlavorText   string `json:"flavor_text"`
	Language     `json:"language"`
	VersionGroup `json:"version_group"`
}

type MoveMetaData struct {
	Ailment        MoveAilment  `json:"ailment"`
	Category       MoveCategory `json:"category"`
	MinHits        int          `json:"min_hits"`
	MaxHits        int          `json:"max_hits"`
	MinTurns       int          `json:"min_turns"`
	MaxTurns       int          `json:"max_turns"`
	Drain          int          `json:"drain"`
	Healing        int          `json:"healing"`
	CritRate       int          `json:"crit_rate"`
	Ailment_chance int          `json:"ailment_chance"`
	FlinchChance   int          `json:"flinch_chance"`
	StatChance     int          `json:"stat_chance"`
}

type MoveStatChange struct {
	Change int `json:"change"`
	Stat   `json:"stat"`
}

type PastMoveStatValues struct {
	Accuracy      int             `json:"accuracy"`
	EffectChance  int             `json:"effect_chance"`
	Power         int             `json:"power"`
	PP            int             `json:"pp"`
	EffectEntries []VerboseEffect `json:"effect_entries"`
	Type          `json:"type"`
	VersionGroup  `json:"version_group"`
}

type MoveAilment struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Moves []Move `json:"moves"`
	Names []Name `json:"names"`
}

type MoveBattleStyle struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []Name `json:"names"`
}

type MoveCategory struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Moves        []Move        `json:"moves"`
	Descriptions []Description `json:"descriptions"`
}
