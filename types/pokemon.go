package types

type Pokemon struct {
	Name    string         `json:"name"`
	Type    []Type         `json:"type"`
	Item    string         `json:"item"`
	Ability string         `json:"ability"`
	Level   int            `json:"level"`
	Tera    Type           `json:"tera"`
	IsTera  bool           `json:"is_tera"`
	EVs     map[string]int `json:"evs"`
	IVs     map[string]int `json:"ivs"`
	Nature  string         `json:"nature"`
	Moves   []Move         `json:"moves"`

	CurrentHP int    `json:"current_hp"`
	Status    string `json:"status"`

	FinalStats map[string]int
}
