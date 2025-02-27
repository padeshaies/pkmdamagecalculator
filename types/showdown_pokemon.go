package types

type ShowdownPokemon struct {
	Name     string         `json:"name"`
	Item     string         `json:"item"`
	Ability  string         `json:"ability"`
	Level    int            `json:"level"`
	TeraType Type           `json:"tera_type"`
	IsTera   bool           `json:"is_tera"`
	Nature   string         `json:"nature"`
	EVs      map[string]int `json:"evs"`
	IVs      map[string]int `json:"ivs"`
	Boosts   map[string]int `json:"boosts"`
	Move1    string         `json:"move1"`
	Move2    string         `json:"move2"`
	Move3    string         `json:"move3"`
	Move4    string         `json:"move4"`
}
