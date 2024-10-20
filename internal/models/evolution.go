package models

type EvolutionChain struct {
	ID              int       `json:"id"`
	BabyTriggerItem Item      `json:"baby_trigger_item"`
	Chain           ChainLink `json:"chain"`
}

type ChainLink struct {
	IsBaby           bool              `json:"is_baby"`
	Species          PokemonSpecies    `json:"species"`
	EvolutionDetails []EvolutionDetail `json:"evolution_details"`
	EvolvesTo        []ChainLink       `json:"evolves_to"`
}

type EvolutionDetail struct {
	Item                  Item             `json:"item"`
	Trigger               EvolutionTrigger `json:"trigger"`
	Gender                int              `json:"gender"`
	HeldItem              Item             `json:"held_item"`
	KnownMove             Move             `json:"known_move"`
	KnownMoveType         Type             `json:"known_move_type"`
	Location              Location         `json:"location"`
	MinLevel              int              `json:"min_level"`
	MinHappiness          int              `json:"min_happiness"`
	MinBeauty             int              `json:"min_beauty"`
	MinAffection          int              `json:"min_affection"`
	NeedsOverworldRain    bool             `json:"needs_overworld_rain"`
	PartySpecies          PokemonSpecies   `json:"party_species"`
	PartyType             Type             `json:"party_type"`
	RelativePhysicalStats int              `json:"relative_physical_stats"`
	TimeOfDay             string           `json:"time_of_day"`
	TradeSpecies          PokemonSpecies   `json:"trade_species"`
	TurnUpsideDown        bool             `json:"turn_upside_down"`
}

type EvolutionTrigger struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	Names          []Name           `json:"names"`
	PokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}
