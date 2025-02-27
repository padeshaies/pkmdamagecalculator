package helpers

import (
	"errors"
	"testing"

	"github.com/padeshaies/pkmdamagecalculator/types"
)

func TestCalculateStats(t *testing.T) {
	cases := []struct {
		input    types.Pokemon
		expected struct {
			pokemon types.Pokemon
			err     error
		}
	}{
		{ // Regular case
			input: types.Pokemon{
				Name:   "Gholdengo",
				Nature: "Timid",
				Level:  50,
				EVs: map[string]int{
					"hp":              52,
					"attack":          0,
					"defense":         4,
					"special-attack":  196,
					"special-defense": 4,
					"speed":           252,
				},
				IVs: map[string]int{
					"hp":              31,
					"attack":          0,
					"defense":         31,
					"special-attack":  31,
					"special-defense": 31,
					"speed":           31,
				},
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{
					Name:   "Gholdengo",
					Nature: "Timid",
					FinalStats: map[string]int{
						"hp":              169,
						"attack":          58,
						"defense":         116,
						"special-attack":  178,
						"special-defense": 112,
						"speed":           149,
					},
				},
				err: nil,
			},
		},
		{ // Shedinja case
			input: types.Pokemon{
				Name:   "Shedinja",
				Nature: "Hardy",
				Level:  1,
				IVs: map[string]int{
					"hp":              31,
					"attack":          31,
					"defense":         31,
					"special-attack":  31,
					"special-defense": 31,
					"speed":           31,
				},
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{
					Name:   "Shedinja",
					Nature: "Hardy",
					FinalStats: map[string]int{
						"hp":              1,
						"attack":          7,
						"defense":         6,
						"special-attack":  5,
						"special-defense": 5,
						"speed":           6,
					},
				},
				err: nil,
			},
		},
		{ // No nature
			input: types.Pokemon{
				Name:  "Gholdengo",
				Level: 50,
				EVs: map[string]int{
					"hp":              52,
					"attack":          0,
					"defense":         4,
					"special-attack":  196,
					"special-defense": 4,
					"speed":           252,
				},
				IVs: map[string]int{
					"hp":              31,
					"attack":          0,
					"defense":         31,
					"special-attack":  31,
					"special-defense": 31,
					"speed":           31,
				},
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{
					Name:   "Gholdengo",
					Nature: "Hardy",
					FinalStats: map[string]int{
						"hp":              169,
						"attack":          65,
						"defense":         116,
						"special-attack":  178,
						"special-defense": 112,
						"speed":           136,
					},
				},
				err: nil,
			},
		},
		{ // Bad Pokemon name
			input: types.Pokemon{
				Name: "Bad Pokemon Name",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{},
				err:     errors.New("pokemon not found"),
			},
		},
		{ // Bad nature
			input: types.Pokemon{
				Name:   "Gholdengo",
				Nature: "Bad Nature",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{},
				err:     errors.New("nature not found"),
			},
		},
	}

	for _, c := range cases {
		tempNature := c.input.Nature
		if tempNature == "" {
			tempNature = "(no nature)"
		}
		t.Log("Calculating stats for", tempNature, c.input.Name)

		err := CalculateStats(&c.input)
		if c.expected.err != nil && err != nil {
			continue
		} else if c.expected.err != nil && err == nil {
			t.Errorf("/!\\ calculateStats(%v) should have returned an error: %v", c.input, c.expected.err)
		} else if c.expected.err == nil && err != nil {
			t.Errorf("/!\\ calculateStats(%v) returned an error: %v", c.input, err)
		}

		if c.input.Nature != c.expected.pokemon.Nature {
			t.Errorf("/!\\ calculateStats(%v) Nature = %s; want %s", c.input, c.input.Nature, c.expected.pokemon.Nature)
		}
		t.Logf("- Expected Nature: %s, got %s\n", c.expected.pokemon.Nature, c.input.Nature)

		for stat, value := range c.expected.pokemon.FinalStats {
			if c.input.FinalStats[stat] != value {
				t.Errorf("/!\\ calculateStats(%v) %s = %v; want %v", c.input, stat, c.input.FinalStats[stat], value)
			}

			t.Logf("- Expected %s: %v, got %v\n", stat, value, c.input.FinalStats[stat])
		}
	}
}
