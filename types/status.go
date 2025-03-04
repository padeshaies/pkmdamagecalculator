package types

import "slices"

type Status string

const (
	Healthy       Status = "healthy"
	Poisoned      Status = "poisoned"
	BadlyPoisoned Status = "badly poisoned"
	Burned        Status = "burned"
	Paralyzed     Status = "paralyzed"
	Asleep        Status = "asleep"
	Frozen        Status = "frozen"
)

func (s Status) IsValid() bool {
	return slices.Contains([]Status{
		Healthy, Poisoned, BadlyPoisoned, Burned, Paralyzed, Asleep, Frozen,
	}, s)
}
