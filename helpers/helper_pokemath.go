package helpers

import (
	"math"
)

// Pokemon's way of applying damage modifiers
// (Applying the modifier M to the damage value D means multiplying D by M; then if the decimal part is ≤0.5, round the result down, otherwise round it up.)
func ApplyMultiplier(damage int, modifier int) int {
	return pokeRound(float64(damage) * float64(modifier) / 0x1000)
}

// ChainMultipliers chains the modifiers together
func ChainMultipliers(modifiers ...int) int {
	var m = 0x1000
	for _, modifier := range modifiers {
		if modifier != 0x1000 {
			m = int(m * modifier / 0x1000)
		}
	}
	return m
}

// pokeRound rounds a float to the nearest integer, but if the decimal is 0.5 or less, it rounds down, otherwise it rounds up
func pokeRound(damage float64) int {
	_, decimal := math.Modf(damage)
	if decimal <= 0.5 {
		return int(damage)
	}
	return int(damage) + 1
}
