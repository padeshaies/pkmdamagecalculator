package helpers

import (
	"errors"
	"slices"
	"testing"

	"github.com/padeshaies/pkmdamagecalculator/types"
)

func TestGetPokemonFromShowdownSet(t *testing.T) {
	cases := []struct {
		name     string
		input    types.ShowdownPokemon
		expected struct {
			pokemon types.Pokemon
			err     error
		}
	}{
		{
			name: "happy path",
			input: types.ShowdownPokemon{
				Name:     "gholdengo",
				Item:     "life orb",
				Ability:  "good as gold",
				Level:    50,
				TeraType: "steel",
				IsTera:   false,
				Nature:   "timid",
				IVs: map[string]int{
					"hp":              31,
					"attack":          0,
					"defense":         31,
					"special-attack":  31,
					"special-defense": 31,
					"speed":           31,
				},
				EVs: map[string]int{
					"hp":              52,
					"attack":          0,
					"defense":         4,
					"special-attack":  196,
					"special-defense": 4,
					"speed":           252,
				},
				Boosts: map[string]int{
					"attack":          0,
					"defense":         0,
					"special-attack":  2,
					"special-defense": 0,
					"speed":           0,
				},
				CurrentHP: 169,
				Status:    "healthy",
				Move1:     "make it rain",
				Move2:     "shadow ball",
				Move3:     "nasty plot",
				Move4:     "protect",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{
					Name:    "Gholdengo",
					Type:    []types.Type{types.Steel, types.Ghost},
					Item:    "Life Orb",
					Ability: "Good as Gold",
					Level:   50,
					Tera:    types.Steel,
					IsTera:  false,
					Stats: map[string]int{
						"hp":              169,
						"attack":          58,
						"defense":         116,
						"special-attack":  356,
						"special-defense": 112,
						"speed":           149,
					},
					CurrentHP: 169,
					Status:    types.Healthy,
					Moves: []types.Move{
						{Name: "Make It Rain", Type: types.Steel, Power: 120, DamageClass: "special", Target: "all-opponents", CriticalHit: false},
						{Name: "Shadow Ball", Type: types.Ghost, Power: 80, DamageClass: "special", Target: "selected-pokemon", CriticalHit: false},
						{Name: "Nasty Plot", Type: types.Dark, Power: 0, DamageClass: "status", Target: "user", CriticalHit: false},
						{Name: "Protect", Type: types.Normal, Power: 0, DamageClass: "status", Target: "user", CriticalHit: false},
					},
				},
				err: nil,
			},
		},
		{
			name: "shedinja",
			input: types.ShowdownPokemon{
				Name: "shedinja",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{
					Name:  "Shedinja",
					Type:  []types.Type{types.Bug, types.Ghost},
					Tera:  types.Bug,
					Level: 1,
					Stats: map[string]int{
						"hp":              1,
						"attack":          7,
						"defense":         6,
						"special-attack":  5,
						"special-defense": 5,
						"speed":           6,
					},
					CurrentHP: 1,
					Status:    types.Healthy,
					Moves:     []types.Move{},
				},
				err: nil,
			},
		},
		{
			name: "invalid name",
			input: types.ShowdownPokemon{
				Name: "invalid name",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{},
				err:     errors.New("invalid name"),
			},
		},
		{
			name: "invalid item",
			input: types.ShowdownPokemon{
				Name: "gholdengo",
				Item: "invalid item",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{},
				err:     errors.New("invalid item"),
			},
		},
		{
			name: "invalid ability",
			input: types.ShowdownPokemon{
				Name:    "gholdengo",
				Ability: "invalid ability",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{},
				err:     errors.New("invalid ability"),
			},
		},
		{
			name: "invalid tera type",
			input: types.ShowdownPokemon{
				Name:     "gholdengo",
				TeraType: "invalid type",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{},
				err:     errors.New("invalid tera type: invalid type"),
			},
		},
		{
			name: "invalid status",
			input: types.ShowdownPokemon{
				Name:   "gholdengo",
				Status: "invalid status",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{},
				err:     errors.New("invalid status: invalid status"),
			},
		},
		{
			name: "invalid nature",
			input: types.ShowdownPokemon{
				Name:   "gholdengo",
				Nature: "invalid nature",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{},
				err:     errors.New("invalid nature: invalid nature"),
			},
		},
		{
			name: "invalid move",
			input: types.ShowdownPokemon{
				Name:  "gholdengo",
				Move1: "invalid move",
			},
			expected: struct {
				pokemon types.Pokemon
				err     error
			}{
				pokemon: types.Pokemon{},
				err:     errors.New("invalid move: invalid move"),
			},
		},
	}

	for _, c := range cases {
		pokemon, err := ImportPokemonFromShowdownSet(c.input)
		if c.expected.err != nil {
			if err == nil {
				t.Errorf("[case: %s] Expected error %v, got nil", c.name, c.expected.err)
			}
		} else if err != nil {
			t.Errorf("[case: %s] Error getting pokemon from showdown set: %v", c.name, err)
		}

		// Compare Name
		if pokemon.Name != c.expected.pokemon.Name {
			t.Errorf("[case: %s] Expected pokemon name to be %s, got %s", c.name, c.expected.pokemon.Name, pokemon.Name)
		}

		// Compare Types
		if !slices.Equal(pokemon.Type, c.expected.pokemon.Type) {
			t.Errorf("[case: %s] Expected pokemon type to be %s, got %s", c.name, c.expected.pokemon.Type, pokemon.Type)
		}

		// Compare Item
		if pokemon.Item != c.expected.pokemon.Item {
			t.Errorf("[case: %s] Expected pokemon item to be %s, got %s", c.name, c.expected.pokemon.Item, pokemon.Item)
		}

		// Compare Ability
		if pokemon.Ability != c.expected.pokemon.Ability {
			t.Errorf("[case: %s] Expected pokemon ability to be %s, got %s", c.name, c.expected.pokemon.Ability, pokemon.Ability)
		}

		// Compare Level
		if pokemon.Level != c.expected.pokemon.Level {
			t.Errorf("[case: %s] Expected pokemon level to be %d, got %d", c.name, c.expected.pokemon.Level, pokemon.Level)
		}

		// Compare Tera
		if pokemon.Tera != c.expected.pokemon.Tera {
			t.Errorf("[case: %s] Expected pokemon tera to be %s, got %s", c.name, c.expected.pokemon.Tera, pokemon.Tera)
		}

		// Compare IsTera
		if pokemon.IsTera != c.expected.pokemon.IsTera {
			t.Errorf("[case: %s] Expected pokemon istera to be %v, got %v", c.name, c.expected.pokemon.IsTera, pokemon.IsTera)
		}

		// Compare Stats
		for stat, value := range c.expected.pokemon.Stats {
			if pokemon.Stats[stat] != value {
				t.Errorf("[case: %s] Expected pokemon stat for %s to be %d, got %d", c.name, stat, value, pokemon.Stats[stat])
			}
		}

		// Compare CurrentHP
		if pokemon.CurrentHP != c.expected.pokemon.CurrentHP {
			t.Errorf("[case: %s] Expected pokemon current HP to be %d, got %d", c.name, c.expected.pokemon.CurrentHP, pokemon.CurrentHP)
		}

		// Compare Status
		if pokemon.Status != c.expected.pokemon.Status {
			t.Errorf("[case: %s] Expected pokemon status to be %s, got %s", c.name, c.expected.pokemon.Status, pokemon.Status)
		}

		// Compare Moves
		if len(pokemon.Moves) != len(c.expected.pokemon.Moves) {
			t.Errorf("[case: %s] Expected %d moves, got %d", c.name, len(c.expected.pokemon.Moves), len(pokemon.Moves))
		}
		for i, move := range c.expected.pokemon.Moves {
			if i >= len(pokemon.Moves) {
				break
			}
			if pokemon.Moves[i].Name != move.Name {
				t.Errorf("[case: %s] Expected move %d to be %s, got %s", c.name, i+1, move.Name, pokemon.Moves[i].Name)
			}
			if pokemon.Moves[i].Type != move.Type {
				t.Errorf("[case: %s] Expected move %d type to be %s, got %s", c.name, i+1, move.Type, pokemon.Moves[i].Type)
			}
			if pokemon.Moves[i].Power != move.Power {
				t.Errorf("[case: %s] Expected move %d power to be %d, got %d", c.name, i+1, move.Power, pokemon.Moves[i].Power)
			}
			if pokemon.Moves[i].DamageClass != move.DamageClass {
				t.Errorf("[case: %s] Expected move %d damage class to be %s, got %s", c.name, i+1, move.DamageClass, pokemon.Moves[i].DamageClass)
			}
			if pokemon.Moves[i].Target != move.Target {
				t.Errorf("[case: %s] Expected move %d target to be %s, got %s", c.name, i+1, move.Target, pokemon.Moves[i].Target)
			}
			if pokemon.Moves[i].CriticalHit != move.CriticalHit {
				t.Errorf("[case: %s] Expected move %d critical hit to be %v, got %v", c.name, i+1, move.CriticalHit, pokemon.Moves[i].CriticalHit)
			}
		}
	}
}
