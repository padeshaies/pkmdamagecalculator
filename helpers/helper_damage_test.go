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
					Name:  "Glaceon",
					Type:  []types.Type{types.Ice},
					Level: 75,
					Stats: map[string]int{
						"attack": 123,
					},
				},
				defender: types.Pokemon{
					Name: "Garchomp",
					Type: []types.Type{types.Dragon, types.Ground},
					Stats: map[string]int{
						"defense": 163,
					},
				},
				move: types.Move{
					Name:        "Ice Fang",
					Type:        types.Ice,
					Power:       65,
					DamageClass: "physical",
					Target:      "selected-pokemon",
				},
				field: types.Field{},
			},
			expected: []int{168, 168, 168, 172, 172, 172, 180, 180, 180, 184, 184, 184, 192, 192, 192, 196},
		},
		{
			name: "status move should not deal damage",
			input: struct {
				attacker types.Pokemon
				defender types.Pokemon
				move     types.Move
				field    types.Field
			}{
				attacker: types.Pokemon{
					Name: "Incineroar",
				},
				defender: types.Pokemon{
					Name: "Porygon2",
				},
				move: types.Move{
					Name:        "Taunt",
					Type:        types.Dark,
					Power:       0,
					DamageClass: "status",
					Target:      "selected-pokemon",
				},
				field: types.Field{},
			},
			expected: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "spread move should deal less damage",
			input: struct {
				attacker types.Pokemon
				defender types.Pokemon
				move     types.Move
				field    types.Field
			}{
				attacker: types.Pokemon{
					Name:  "Landorus-Therian",
					Type:  []types.Type{types.Ground, types.Flying},
					Level: 50,
					Stats: map[string]int{
						"attack": 216,
					},
				},
				defender: types.Pokemon{
					Name: "Heatran",
					Type: []types.Type{types.Fire, types.Steel},
					Stats: map[string]int{
						"defense": 126,
					},
				},
				move: types.Move{
					Name:        "Earthquake",
					Type:        types.Ground,
					Power:       100,
					DamageClass: "physical",
					Target:      "all-other-pokemon",
				},
				field: types.Field{},
			},
			expected: []int{292, 292, 300, 304, 304, 312, 312, 316, 316, 324, 328, 328, 336, 336, 340, 348},
		},
		// TODO: ADD WEATHER TESTS
		{
			name: "critical hit should deal more damage",
			input: struct {
				attacker types.Pokemon
				defender types.Pokemon
				move     types.Move
				field    types.Field
			}{
				attacker: types.Pokemon{
					Name:  "Urshifu-Single-Strike",
					Type:  []types.Type{types.Fighting, types.Dark},
					Level: 50,
					Stats: map[string]int{
						"attack": 182,
					},
				},
				defender: types.Pokemon{
					Name: "Flutter Mane",
					Type: []types.Type{types.Ghost, types.Fairy},
					Stats: map[string]int{
						"defense": 75,
					},
				},
				move: types.Move{
					Name:        "Wicked Blow",
					Type:        types.Dark,
					Power:       75,
					DamageClass: "physical",
					Target:      "selected-pokemon",
					CriticalHit: true,
				},
				field: types.Field{},
			},
			expected: []int{156, 157, 160, 162, 163, 165, 166, 169, 171, 172, 174, 177, 178, 180, 181, 184},
		},
		{
			name: "burn should deal less damage if physical",
			input: struct {
				attacker types.Pokemon
				defender types.Pokemon
				move     types.Move
				field    types.Field
			}{
				attacker: types.Pokemon{
					Name:   "Urshifu-Rapid-Strike",
					Level:  50,
					Status: types.Burned,
					Type:   []types.Type{types.Fighting, types.Water},
					Stats: map[string]int{
						"attack": 182,
					},
				},
				defender: types.Pokemon{
					Name: "Incineroar",
					Type: []types.Type{types.Fire, types.Dark},
					Stats: map[string]int{
						"defense": 113,
					},
				},
				move: types.Move{
					Name:        "Close Combat",
					Type:        types.Fighting,
					Power:       120,
					DamageClass: "physical",
					Target:      "selected-pokemon",
				},
				field: types.Field{},
			},
			expected: []int{109, 111, 112, 114, 115, 117, 118, 120, 120, 121, 123, 124, 126, 127, 129, 130},
		},
	}
	for _, c := range cases {
		damage := CalculateDamage(c.input.attacker, c.input.defender, c.input.move, c.input.field)
		if !reflect.DeepEqual(damage, c.expected) {
			t.Errorf("expected %v, got %v", c.expected, damage)
		}
	}
}
