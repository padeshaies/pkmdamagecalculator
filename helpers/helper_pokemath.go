package helpers

import (
	"math"
)

const (
	Modifier0_5x  = 0x800
	Modifier0_75x = 0xC00
	Modifier1x    = 0x1000
	Modifier1_3x  = 0x14CD
	Modifier1_5x  = 0x1800
	Modifier2x    = 0x2000
	Modifier2_25x = 0x2400
)

// Pokemon's way of applying damage modifiers
// (Applying the modifier M to the damage value D means multiplying D by M; then if the decimal part is ≤0.5, round the result down, otherwise round it up.)
func ApplyMultiplier(value int, modifier int) int {
	return pokeRound(float64(value) * float64(modifier) / Modifier1x)
}

// ChainMultipliers chains the modifiers together
func ChainMultipliers(modifiers ...int) int {
	var m = Modifier1x
	for _, modifier := range modifiers {
		if modifier != Modifier1x {
			m = int(m * modifier / Modifier1x)
		}
	}
	return m
}

// pokeRound rounds a float to the nearest integer, but if the decimal is 0.5 or less, it rounds down, otherwise it rounds up
func pokeRound(value float64) int {
	_, decimal := math.Modf(value)
	if decimal <= 0.5 {
		return int(value)
	}
	return int(value) + 1
}
