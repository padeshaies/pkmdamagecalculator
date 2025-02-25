package types

type Move struct {
	Name        string
	Type        Type
	DamageClass string
	Power       int
	Target      string
	CriticalHit bool
}
