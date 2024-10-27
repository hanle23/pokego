package models

type Machine struct {
	ID           int          `json:"id"`
	Item         Item         `json:"item"`
	Move         Move         `json:"move"`
	VersionGroup VersionGroup `json:"version_group"`
}

type Machines struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}
