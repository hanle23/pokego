package models

type EncounterMethod struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Names []Name `json:"names"`
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
