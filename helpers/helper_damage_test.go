package helpers

import (
	"testing"

	"github.com/padeshaies/pkmdamagecalculator/types"
)

func TestCalculateDamage(t *testing.T) {
	cases := []struct {
		input struct {
			attacker types.Pokemon
			defender types.Pokemon
			move     types.Move
		}
		expected []int
	}{
		// Bulbapedia example (https://bulbapedia.bulbagarden.net/wiki/Damage#Example)
		{
			input: struct {
				attacker types.Pokemon
				defender types.Pokemon
				move     types.Move
			}{
				attacker: types.Pokemon{
					Name:  "glaceon",
					Type:  []string{"ice"},
					Level: 75,
					FinalStats: map[string]int{
						"attack": 123,
					},
				},
				defender: types.Pokemon{
					Name: "garchomp",
					Type: []string{"dragon", "ground"},
					FinalStats: map[string]int{
						"defense": 163,
					},
				},
				move: types.Move{
					Name:        "ice fang",
					Type:        "ice",
					Power:       65,
					DamageClass: "physical",
					Target:      "selected-pokemon",
				},
			},
			expected: []int{168, 168, 168, 172, 172, 172, 180, 180, 180, 184, 184, 184, 192, 192, 192, 196},
		},
		// Absurdly high damage (to be update as we add more modifiers)
		{
			input: struct {
				attacker types.Pokemon
				defender types.Pokemon
				move     types.Move
			}{
				attacker: types.Pokemon{
					Name:  "miraidon",
					Type:  []string{"electric", "dragon"},
					Level: 100,
					FinalStats: map[string]int{
						"special-attack": 405,
					},
				},
				defender: types.Pokemon{
					Name: "wingull",
					Type: []string{"water", "flying"},
					FinalStats: map[string]int{
						"special-defense": 5,
					},
				},
				move: types.Move{
					Name:        "thunder",
					Type:        "electric",
					Power:       110,
					DamageClass: "special",
					Target:      "selected-pokemon",
				},
			},
			expected: []int{38176, 38620, 39072, 39520, 39972, 40420, 40872, 41320, 41764, 42216, 42664, 43116, 43564, 44016, 44464, 44916},
		},
	}

	for _, c := range cases {
		got := CalculateDamage(c.input.attacker, c.input.defender, c.input.move)

		for i := 0; i < len(got); i++ {
			if got[i] != c.expected[i] {
				t.Errorf("calculateDamage(%v, %v, %v)", c.input.attacker.Name, c.input.defender.Name, c.input.move.Name)
				t.Errorf("got      %v", got)
				t.Errorf("expected %v", c.expected)
				break
			}
		}
	}
}
