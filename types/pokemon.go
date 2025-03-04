package types

type Pokemon struct {
	Name      string
	Type      []Type
	Item      string
	Ability   string
	Level     int
	Tera      Type
	IsTera    bool
	Stats     map[string]int
	CurrentHP int
	Status    Status
	Moves     []Move
}
