package helpers

import (
	"reflect"
	"testing"

	"github.com/padeshaies/pkmdamagecalculator/types"
)

func TestCalculateDamage(t *testing.T) {
	cases := []struct {
		name  string
		input struct {
			attacker types.Pokemon
			defender types.Pokemon
			move     types.Move
			field    types.Field
		}
		expected []int
	}{
		{
			// source: https://bulbapedia.bulbagarden.net/wiki/Damage#Example
			name: "base bulbapedia example (glaceon vs garchomp)",
			input: struct {
				attacker types.Pokemon
				defender types.Pokemon
				move     types.Move
				field    types.Field
			}{
				attacker: types.Pokemon{
					Name:  "glaceon",
					Type:  []types.Type{types.Ice},
					Level: 75,
					Stats: map[string]int{
						"attack": 123,
					},
				},
				defender: types.Pokemon{
					Name: "garchomp",
					Type: []types.Type{types.Dragon, types.Ground},
					Stats: map[string]int{
						"defense": 163,
					},
				},
				move: types.Move{
					Name:        "ice fang",
					Type:        types.Ice,
					Power:       65,
					DamageClass: "physical",
					Target:      "selected-pokemon",
				},
				field: types.Field{},
			},
			expected: []int{168, 168, 168, 172, 172, 172, 180, 180, 180, 184, 184, 184, 192, 192, 192, 196},
		},
	}

	for _, c := range cases {
		damage := CalculateDamage(c.input.attacker, c.input.defender, c.input.move, c.input.field)
		if !reflect.DeepEqual(damage, c.expected) {
			t.Errorf("expected %v, got %v", c.expected, damage)
		}
	}
}
