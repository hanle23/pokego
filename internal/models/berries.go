package models

type Berries struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

type Berry struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	GrowthTime       int    `json:"growth_time"`
	MaxHarvest       int    `json:"max_harvest"`
	NaturalGiftPower int    `json:"natural_gift_power"`
	Size             int    `json:"size"`
	Smoothness       int    `json:"smoothness"`
	SoilDryness      int    `json:"soil_dryness"`
	Firmness         struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"firmness"`
	Flavors []struct {
		Flavor struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"flavor"`
		Potency int `json:"potency"`
	} `json:"flavors"`
	Item struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"item"`
	NaturalGiftType struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"natural_gift_type"`
}

type BerryFirmness struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Berries []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"berries"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}

type BerryFlavor struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Berries []struct {
		Berry struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"berry"`
		Potency int `json:"potency"`
	} `json:"berries"`
	ContestType struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"contest_type"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}

type BerryFlavorMap struct {
	Potency int `json:"potency"`
	Berry   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"berry"`
}
