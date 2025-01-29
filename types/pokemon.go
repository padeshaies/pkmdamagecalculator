package types

type Pokemon struct {
	Name    string         `json:"name"`
	Item    string         `json:"item"`
	Ability string         `json:"ability"`
	Level   int            `json:"level"`
	Tera    string         `json:"tera"`
	EVs     map[string]int `json:"evs"`
	IVs     map[string]int `json:"ivs"`
	Nature  string         `json:"nature"`
	Moves   []string       `json:"moves"`

	Stats map[string]int
}

type Move struct {
	Name string `json:"name"`
}
