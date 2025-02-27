package helpers

/*func TestCalculateDamage(t *testing.T) {
	cases := []struct {
		input struct {
			attacker types.Pokemon
			defender types.Pokemon
			move     types.Move
			field    types.Field
		}
		expected []int
	}{
		// Bulbapedia example (https://bulbapedia.bulbagarden.net/wiki/Damage#Example)
		{
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
					FinalStats: map[string]int{
						"attack": 123,
					},
				},
				defender: types.Pokemon{
					Name: "garchomp",
					Type: []types.Type{types.Dragon, types.Ground},
					FinalStats: map[string]int{
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
		// Using Miraidon as an example of multiple modifiers
		{
			input: struct {
				attacker types.Pokemon
				defender types.Pokemon
				move     types.Move
				field    types.Field
			}{
				attacker: types.Pokemon{
					Name:    "miraidon",
					Type:    []types.Type{types.Electric, types.Dragon},
					Level:   50,
					Ability: "hadron engine",
					Item:    "Choice Specs",
					FinalStats: map[string]int{
						"special-attack": 205,
					},
				},
				defender: types.Pokemon{
					Name: "urshifu-rapid-strike",
					Type: []types.Type{types.Fighting, types.Water},
					FinalStats: map[string]int{
						"special-defense": 80,
					},
				},
				move: types.Move{
					Name:        "electro drift",
					Type:        types.Electric,
					Power:       100,
					DamageClass: "special",
					Target:      "selected-pokemon",
				},
				field: types.Field{
					Terrain: types.ElectricTerrain,
				},
			},
			expected: []int{998, 1010, 1022, 1032, 1044, 1056, 1068, 1080, 1092, 1104, 1116, 1128, 1140, 1152, 1164, 1176},
		},
	}

	for _, c := range cases {
		got := []int{} //CalculateDamage(c.input.attacker, c.input.defender, c.input.move, c.input.field)

		for i := 0; i < len(got); i++ {
			if got[i] != c.expected[i] {
				t.Errorf("calculateDamage(%v, %v, %v)", c.input.attacker.Name, c.input.defender.Name, c.input.move.Name)
				t.Errorf("got      %v", got)
				t.Errorf("expected %v", c.expected)
				break
			}
		}
	}
}*/
