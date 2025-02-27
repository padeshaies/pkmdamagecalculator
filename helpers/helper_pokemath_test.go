package helpers

import (
	"testing"
)

// Pokemon's way of applying damage modifiers
// (Applying the modifier M to the damage value D means multiplying D by M; then if the decimal part is ≤0.5, round the result down, otherwise round it up.)
func TestApplyMultiplier(t *testing.T) {
	cases := []struct {
		input struct {
			damage   int
			modifier int
		}
		expected int
	}{
		// 1.5 should be rounded down to 1
		{
			input: struct {
				damage   int
				modifier int
			}{damage: 1, modifier: 0x1800},
			expected: 1,
		},
		// 1.5000000001 should be rounded up to 2
		{
			input: struct {
				damage   int
				modifier int
			}{damage: 1, modifier: 0x1801},
			expected: 2,
		},
	}

	for _, c := range cases {
		got := ApplyMultiplier(c.input.damage, c.input.modifier)
		if got != c.expected {
			t.Errorf("ApplyMultiplier(%v, %v) = %v; expected %v", c.input.damage, c.input.modifier, got, c.expected)
		}
	}
}

// ChainMultipliers chains the modifiers together
func TestChainMultipliers(t *testing.T) {
	cases := []struct {
		input struct {
			modifiers []int
		}
		expected int
	}{
		{
			input: struct {
				modifiers []int
			}{modifiers: []int{0x1000, 0x1800}},
			expected: 0x1800, // 6144 = 1.5 * 4096
		},
		{
			input: struct {
				modifiers []int
			}{modifiers: []int{0x1800, 0x14CC}},
			expected: 0x1F32, // 7986 = 1.5 * 1.3 * 4096 (rounded down)
		},
	}

	for _, c := range cases {
		got := ChainMultipliers(c.input.modifiers...)
		if got != c.expected {
			t.Errorf("ChainMultipliers(%v) = %v; expected %v", c.input.modifiers, got, c.expected)
		}
	}
}
