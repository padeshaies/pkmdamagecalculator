package helpers

import (
	"math"
	"slices"

	"github.com/padeshaies/pkmdamagecalculator/types"
)

// CalculateDamage calculates the damage dealt by the attacker to the defender using the given move
// Source: https://www.smogon.com/bw/articles/bw_complete_damage_formula
func CalculateDamage(attacker types.Pokemon, defender types.Pokemon, move types.Move, field types.Field) []int {
	// Calculate the damage dealt by the attacker to the defender using the given move
	damage := make([]int, 16)

	// Increase the power of the move depending on the terrain
	if field.Terrain != "" {
		considerTerrain(field, &move, attacker, defender)
	}

	// Base damage calculation (ie: BaseDamage = ((((2 × Level) ÷ 5 + 2) * BasePower * [Sp]Atk) ÷ [Sp]Def) ÷ 50 + 2)
	var a, d int
	switch move.DamageClass {
	case "physical":
		// TODO: Refactor this in a function for Choice items and other abilities that modify the stats
		if attacker.Status == "burn" && attacker.Ability == "Guts" {
			a = int(float64(attacker.FinalStats["attack"]) * 1.5)
		} else {
			a = int(attacker.FinalStats["attack"])
		}
		d = int(defender.FinalStats["defense"])
	case "special":
		a = int(attacker.FinalStats["special-attack"])
		d = int(defender.FinalStats["special-defense"])
	default:
		return damage // Status moves don't deal damage
	}
	// 🤮 (to emulate the multiple floors they do in the games, we're forced to it like this...)
	baseDamage := int(int(int((2*attacker.Level)/5+2)*move.Power*a)/d/50 + 2)

	// Apply the multi-target modifier
	switch move.Target {
	case "selected-pokemon":
		break
	case "all-other-pokemon":
	case "all-opponents":
		baseDamage = applyMultiplier(baseDamage, 0.75)
	default:
		return damage // Not targeting a Pokemon
	}

	// Apply the weather modifier
	if field.Weather != "" {
		baseDamage = applyMultiplier(baseDamage, calculateWeatherModifier(field.Weather, move.Type))
	}

	// Apply the critical hit modifier
	if move.CriticalHit && defender.Ability != "Battle Armor" && defender.Ability != "Shell Armor" {
		baseDamage = applyMultiplier(baseDamage, 1.5)
	}

	// Store the STAB modifier
	stabModifier := calculateStab(attacker.Type, move.Type)

	// Store the type effectiveness
	typeEffectiveness := calculateTypeEffectiveness(move.Type, defender.Type, defender.Ability)
	if typeEffectiveness == 0 {
		return damage // Defender is immune to the move
	}

	// Store the user's burn status
	isBurned := attacker.Status == "burn" && move.DamageClass == "physical" && attacker.Ability != "Guts"

	// add stuff to damage (ie: abilities, items, etc)

	// Add randomization and chain the modifiers ('cause turns out that switching between floors and rounds is a pain 🙃)
	for i, rand := range []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0} {
		tempDamage := int(baseDamage * (100 - rand) / 100)

		// Apply STAB modifier
		tempDamage = applyMultiplier(tempDamage, stabModifier)

		// Apply with type effectiveness
		tempDamage = int(float64(tempDamage) * typeEffectiveness)

		// Apply with user's burn status
		if isBurned {
			tempDamage = int(float64(tempDamage) * 0.5)
		}

		// Make sure the damage is at least 1
		if tempDamage < 1 {
			tempDamage = 1
		}

		// TODO apply final modifier
		damage[i] = tempDamage
	}
	return damage
}

// Pokemon's way of applying damage modifiers
// (Applying the modifier M to the damage value D means multiplying D by M; then if the decimal part is ≤0.5, round the result down, otherwise round it up.)
func applyMultiplier(damage int, modifier float64) int {
	modifiedDamage := float64(damage) * modifier
	_, decimal := math.Modf(modifiedDamage)
	if decimal <= 0.5 {
		return int(modifiedDamage)
	} else {
		return int(modifiedDamage) + 1
	}
}

/*func chainMultipliers(damage int, modifiers ...float64) int {
	modifiedDamage := float64(damage)
	for _, modifier := range modifiers {
		modifiedDamage = modifiedDamage * modifier
	}
	return int(modifiedDamage)
}*/

// calculateWeatherModifier returns the weather-based damage modifier for a move
func calculateWeatherModifier(weather types.Weather, moveType types.Type) float64 {
	if weather == "" {
		return 1.0
	}

	switch weather {
	case types.Sun:
		if moveType == types.Fire {
			return 1.5
		}
		if moveType == types.Water {
			return 0.5
		}
	case types.Rain:
		if moveType == types.Fire {
			return 0.5
		}
		if moveType == types.Water {
			return 1.5
		}
	}
	return 1.0
}

func isGrounded(pokemon types.Pokemon, field types.Field) bool {
	return !(slices.Contains(pokemon.Type, types.Flying) || pokemon.Ability == "Levitate" || pokemon.Item == "Air Balloon") || (field.Gravity || pokemon.Item == "Iron Ball")
}

// TODO: consider moves that are affected by terrain (terrain pulse, expanding force, etc.)
func considerTerrain(field types.Field, move *types.Move, attacker types.Pokemon, defender types.Pokemon) {
	bpBoost := 1.0
	switch field.Terrain {
	case types.ElectricTerrain:
		if move.Type == types.Electric && isGrounded(attacker, field) {
			bpBoost = 1.3
		}
	case types.GrassyTerrain:
		if move.Type == types.Grass && isGrounded(attacker, field) {
			bpBoost = 1.3
		}
	case types.PsychicTerrain:
		if move.Type == types.Psychic && isGrounded(attacker, field) {
			bpBoost = 1.3
		}
	case types.MistyTerrain:
		if move.Type == types.Dragon && isGrounded(defender, field) {
			bpBoost = 0.5
		}
	}

	move.Power = int(float64(move.Power) * bpBoost)
}

func calculateTypeEffectiveness(moveType types.Type, defenderTypes []types.Type, defenderAbility string) float64 {
	effectiveness := types.TypeChart[moveType][defenderTypes[0]]
	if len(defenderTypes) > 1 {
		effectiveness = effectiveness * types.TypeChart[moveType][defenderTypes[1]]
	}
	if defenderAbility == "Wonder Guard" && effectiveness < 2 {
		effectiveness = 0
	}
	return effectiveness
}

func calculateStab(attackerTypes []types.Type, moveType types.Type) float64 {
	// TODO: Consider Adaptability and tera type
	if slices.Contains(attackerTypes, moveType) {
		return 1.5 // STAB bonus is 1.5x
	}
	return 1.0
}
