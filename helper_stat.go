package main

import (
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/padeshaies/pkmdamagecalculator/types"
)

func calculateStats(pokemon *types.Pokemon) error {
	apiInfo, err := pokeapi.Pokemon(pokemon.Name)
	if err != nil {
		return err
	}

	pokemon.Stats = make(map[string]int)

	if pokemon.Name == "Shedinja" {
		pokemon.Stats["hp"] = 1
	} else {
		pokemon.Stats["hp"] = (2*apiInfo.Stats[0].BaseStat+pokemon.IVs["hp"]+pokemon.EVs["hp"]/4)*pokemon.Level/100 + pokemon.Level + 10
	}

	if pokemon.Nature == "" {
		pokemon.Nature = "Hardy"
	}
	apiNature, err := pokeapi.Nature(strings.ToLower(pokemon.Nature))
	if err != nil {
		return err
	}

	for i, stat := range []string{"attack", "defense", "special-attack", "special-defense", "speed"} {
		pokemon.Stats[stat] = (2*apiInfo.Stats[i+1].BaseStat+pokemon.IVs[stat]+pokemon.EVs[stat]/4)*pokemon.Level/100 + 5

		if apiNature.IncreasedStat != nil {
			if apiNature.IncreasedStat.(map[string]interface{})["name"].(string) == stat {
				pokemon.Stats[stat] = int(float64(pokemon.Stats[stat]) * 1.1)
			}
		}

		if apiNature.DecreasedStat != nil {
			if apiNature.DecreasedStat.(map[string]interface{})["name"].(string) == stat {
				pokemon.Stats[stat] = int(float64(pokemon.Stats[stat]) * 0.9)
			}
		}
	}

	return nil
}
