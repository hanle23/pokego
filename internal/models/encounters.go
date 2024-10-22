package models

type EncounterMethod struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Names []Name `json:"names"`
}

type EncounterMethods struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

type EncounterCondition struct {
	ID     int                       `json:"id"`
	Name   string                    `json:"name"`
	Names  []Name                    `json:"names"`
	Values []EncounterConditionValue `json:"values"`
}

type EncounterConditionValue struct {
	ID        int                `json:"id"`
	Name      string             `json:"name"`
	Condition EncounterCondition `json:"condition"`
	Names     []Name             `json:"names"`
}
