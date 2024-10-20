package models

type Machine struct {
	ID           int          `json:"id"`
	Item         Item         `json:"item"`
	Move         Move         `json:"move"`
	VersionGroup VersionGroup `json:"version_group"`
}
