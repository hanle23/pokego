package models

type ContestTypes struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

type ContestType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BerryFlavor `json:"berry_flavor"`
	Names       []ContestName `json:"names"`
}

type ContestName struct {
	Name     string `json:"name"`
	Color    string `json:"color"`
	Language `json:"language"`
}

type ContestEffects struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

type ContestEffect struct {
	ID                int          `json:"id"`
	Appeal            int          `json:"appeal"`
	Jam               int          `json:"jam"`
	EffectEntries     []Effect     `json:"effect_entries"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
}

type SuperContestEffects struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

type SuperContestEffect struct {
	ID                int          `json:"id"`
	Appeal            int          `json:"appeal"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
	Moves             []Move       `json:"moves"`
}
