package helpers

import (
	"testing"

	"github.com/padeshaies/pkmdamagecalculator/types"
)

func TestCalculateStats(t *testing.T) {
	cases := []struct {
		input    types.Pokemon
		expected types.Pokemon
	}{
		{
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
			expected: types.Pokemon{
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
		},
		{
			input: types.Pokemon{
				Name:  "Shedinja",
				Level: 1,
				IVs: map[string]int{
					"hp":              31,
					"attack":          31,
					"defense":         31,
					"special-attack":  31,
					"special-defense": 31,
					"speed":           31,
				},
			},
			expected: types.Pokemon{
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
		},
	}

	for _, c := range cases {
		tempNature := c.input.Nature
		if tempNature == "" {
			tempNature = "(no nature)"
		}
		t.Log("Calculating stats for", tempNature, c.input.Name)

		err := CalculateStats(&c.input)
		if err != nil {
			t.Errorf("/!\\ calculateStats(%v) returned an error: %v", c.input, err)
		}

		if c.input.Nature != c.expected.Nature {
			t.Errorf("/!\\ calculateStats(%v) Nature = %s; want %s", c.input, c.input.Nature, c.expected.Nature)
		}
		t.Logf("- Expected Nature: %s, got %s\n", c.expected.Nature, c.input.Nature)

		for stat, value := range c.expected.FinalStats {
			if c.input.FinalStats[stat] != value {
				t.Errorf("/!\\ calculateStats(%v) %s = %v; want %v", c.input, stat, c.input.FinalStats[stat], value)
			}

			t.Logf("- Expected %s: %v, got %v\n", stat, value, c.input.FinalStats[stat])
		}
	}
}
