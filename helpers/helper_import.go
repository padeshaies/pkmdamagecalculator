package helpers

import (
	"fmt"
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/padeshaies/pkmdamagecalculator/types"
)

func ImportPokemonFromShowdownSet(showdownSet types.ShowdownPokemon) (types.Pokemon, error) {
	apiInfo, err := pokeapi.Pokemon(showdownSet.Name)
	if err != nil {
		return types.Pokemon{}, err
	}

	// Get the English name of the pokemon
	var prettyName string
	apiSpecies, _ := pokeapi.PokemonSpecies(apiInfo.Species.Name)
	for _, name := range apiSpecies.Names {
		if name.Language.Name == "en" {
			prettyName = name.Name
			break
		}
	}

	// Get the English name of the item
	var prettyItem string
	apiItem, err := pokeapi.Item(strings.ReplaceAll(strings.ToLower(showdownSet.Item), " ", "-"))
	if err != nil {
		return types.Pokemon{}, err
	}
	for _, name := range apiItem.Names {
		if name.Language.Name == "en" {
			prettyItem = name.Name
			break
		}
	}

	// Get the English name of the ability
	var prettyAbility string
	apiAbility, err := pokeapi.Ability(strings.ReplaceAll(strings.ToLower(showdownSet.Ability), " ", "-"))
	if err != nil {
		return types.Pokemon{}, err
	}
	for _, name := range apiAbility.Names {
		if name.Language.Name == "en" {
			prettyAbility = name.Name
			break
		}
	}

	// Initialize Pokemon
	pokemon := types.Pokemon{
		Name:    prettyName,
		Item:    prettyItem,
		Ability: prettyAbility,
		Level:   max(1, showdownSet.Level), // Level 1 is the minimum level for a pokemon
		IsTera:  showdownSet.IsTera,
		Stats:   make(map[string]int, 6),
	}

	// Add Types
	for _, apiType := range apiInfo.Types {
		pokemon.Type = append(pokemon.Type, types.Type(apiType.Type.Name))
	}

	// Add Tera Type
	if showdownSet.TeraType == "" {
		pokemon.Tera = pokemon.Type[0]
	} else {
		pokemon.Tera = types.Type(strings.ToLower(showdownSet.TeraType))
		if !pokemon.Tera.IsValid() {
			return types.Pokemon{}, fmt.Errorf("invalid tera type: %s", showdownSet.TeraType)
		}
	}

	// Add Status
	if showdownSet.Status == "" {
		pokemon.Status = types.Healthy
	} else {
		pokemon.Status = types.Status(strings.ToLower(showdownSet.Status))
		if !pokemon.Status.IsValid() {
			return types.Pokemon{}, fmt.Errorf("invalid status: %s", showdownSet.Status)
		}
	}

	// Calculcate HP
	if pokemon.Name == "Shedinja" {
		pokemon.Stats["hp"] = 1
	} else {
		pokemon.Stats["hp"] = (2*apiInfo.Stats[0].BaseStat+safeGetIV(showdownSet, "hp")+safeGetEV(showdownSet, "hp")/4)*pokemon.Level/100 + pokemon.Level + 10
	}
	if showdownSet.CurrentHP > 0 {
		pokemon.CurrentHP = showdownSet.CurrentHP
	} else {
		pokemon.CurrentHP = pokemon.Stats["hp"]
	}

	// Calculate Stats
	if showdownSet.Nature == "" {
		showdownSet.Nature = "hardy"
	}
	apiNature, err := pokeapi.Nature(strings.ToLower(showdownSet.Nature))
	if err != nil {
		return types.Pokemon{}, err
	}
	for i, stat := range []string{"attack", "defense", "special-attack", "special-defense", "speed"} {
		pokemon.Stats[stat] = (2*apiInfo.Stats[i+1].BaseStat+safeGetIV(showdownSet, stat)+safeGetEV(showdownSet, stat)/4)*pokemon.Level/100 + 5

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

		boost := safeGetBoost(showdownSet, stat)
		if boost != 0 {
			// If boost is positive, multiply by (2 + boost) / 2 and if negative, multiply by 2 / (2 + boost)
			// source: https://bulbapedia.bulbagarden.net/wiki/Stat_modifier#Stage_multipliers
			pokemon.Stats[stat] = int(float64(max(2, 2+boost)) / float64(max(2, 2+(-1*boost))) * float64(pokemon.Stats[stat]))
		}
	}

	// Add Moves
	for _, move := range []string{showdownSet.Move1, showdownSet.Move2, showdownSet.Move3, showdownSet.Move4} {
		if move != "" {
			apiMove, err := pokeapi.Move(strings.ReplaceAll(strings.ToLower(move), " ", "-"))
			if err != nil {
				return types.Pokemon{}, err
			}

			// Get the English name of the move
			var moveName string
			for _, name := range apiMove.Names {
				if name.Language.Name == "en" {
					moveName = name.Name
					break
				}
			}

			pokemon.Moves = append(pokemon.Moves, types.Move{
				Name:        moveName,
				Power:       apiMove.Power,
				Type:        types.Type(apiMove.Type.Name),
				DamageClass: apiMove.DamageClass.Name,
				Target:      apiMove.Target.Name,
				CriticalHit: false, // TODO: Add Critical Hit
			})
		}
	}

	return pokemon, nil
}

// By default, EVs are 0
func safeGetEV(showdownSet types.ShowdownPokemon, stat string) int {
	ev, ok := showdownSet.EVs[stat]
	if !ok {
		return 0
	}
	return ev
}

// By default, IVs are 31
func safeGetIV(showdownSet types.ShowdownPokemon, stat string) int {
	iv, ok := showdownSet.IVs[stat]
	if !ok {
		return 31
	}
	return iv
}

// By default, Boosts are 0
func safeGetBoost(showdownSet types.ShowdownPokemon, stat string) int {
	boost, ok := showdownSet.Boosts[stat]
	if !ok {
		return 0
	}
	return boost
}
