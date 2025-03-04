package helpers

import (
	"slices"

	"github.com/padeshaies/pkmdamagecalculator/types"
)

// CalculateDamage calculates the damage dealt by the attacker to the defender using the given move
// Source: https://www.smogon.com/bw/articles/bw_complete_damage_formula
func CalculateDamage(attacker types.Pokemon, defender types.Pokemon, move types.Move, field types.Field) []int {
	// Calculate the damage dealt by the attacker to the defender using the given move
	damage := make([]int, 16)

	// Status moves don't deal damage
	if move.DamageClass != "physical" && move.DamageClass != "special" {
		return damage
	}

	// 1. CALCULATE BASE DAMAGE (includes stats, bp, etc)
	baseDamage := getBaseDamage(attacker, defender, move, field)

	// 2. APPLY MULTI-TARGET MODIFIER
	switch move.Target {
	case "selected-pokemon":
		break
	case "all-other-pokemon":
	case "all-opponents":
		baseDamage = ApplyMultiplier(baseDamage, Modifier0_75x)
	default:
		return damage // Not targeting a Pokemon
	}

	// 3. APPLY WEATHER MODIFIER
	if field.Weather != "" {
		baseDamage = ApplyMultiplier(baseDamage, getWeatherModifier(field.Weather, move.Type))
	}

	// 4. APPLY CRITICAL HIT MODIFIER
	if move.CriticalHit && defender.Ability != "Battle Armor" && defender.Ability != "Shell Armor" {
		baseDamage = ApplyMultiplier(baseDamage, Modifier1_5x)
	}

	// STORE FUTURE MODIFIERS
	stabModifier := getStabModifier(attacker, move)
	typeEffectiveness := getTypeEffectiveness(&move, defender, field)
	isBurned := attacker.Status == "burn" && move.DamageClass == "physical" && attacker.Ability != "Guts"
	finalModifiers := []int{}

	// 5. ALTER WITH RANDOM FACTOR
	for i := 0; i < 16; i++ {
		tempDamage := int(float64(baseDamage) * (85.0 + float64(i)) / 100.0)

		// 6. APPLY STAB MODIFIER
		tempDamage = ApplyMultiplier(tempDamage, stabModifier)

		// 7. APPLY TYPE EFFECTIVENESS
		tempDamage = int(float64(tempDamage) * typeEffectiveness)

		// 8. ALTER WITH BURN STATUS
		if isBurned {
			tempDamage = int(float64(tempDamage) * 0.5)
		}

		// 9. CHAIN FINAL MODIFIERS
		finalModifier := ChainMultipliers(finalModifiers...)
		tempDamage = ApplyMultiplier(tempDamage, finalModifier)

		// 10. STORE DAMAGE
		damage[i] = tempDamage
	}

	return damage
}

func getBaseDamage(attacker types.Pokemon, defender types.Pokemon, move types.Move, field types.Field) int {
	// Calculate the base damage of the move
	basePower := getBasePower(move, attacker, defender, field)
	baseDamageMods := getBaseDamageModifiers(move /*attacker, defender,*/, field)
	basePower = ApplyMultiplier(basePower, ChainMultipliers(baseDamageMods...))

	var attack, defense int
	if move.DamageClass == "physical" {
		attack = attacker.Stats["attack"]
		defense = defender.Stats["defense"]
	} else {
		attack = attacker.Stats["special-attack"]
		defense = defender.Stats["special-defense"]
	}

	baseDamage := int(int(int((2*attacker.Level)/5+2)*basePower*attack)/defense/50 + 2)

	return baseDamage
}

func getBasePower(move types.Move, attacker types.Pokemon, defender types.Pokemon, field types.Field) int {
	var basePower int

	switch move.Name {

	// HP Based Moves
	case "Eruption":
	case "Water Spout":
	case "Dragon Energy":
		basePower = max(1, int(float64(attacker.CurrentHP)/float64(attacker.Stats["hp"])*150.0))

	// Default
	default:
		basePower = move.Power
	}

	return basePower
}

func getBaseDamageModifiers(move types.Move /*attacker types.Pokemon, defender types.Pokemon,*/, field types.Field) []int {
	var modifiers []int

	isAttackerGrounded, isDefenderGrounded := true, true // TODO: Implement this

	// Offensive Terrains
	if isAttackerGrounded {
		if (field.Terrain == types.ElectricTerrain && move.Type == types.Electric) ||
			(field.Terrain == types.GrassyTerrain && move.Type == types.Grass) ||
			(field.Terrain == types.PsychicTerrain && move.Type == types.Psychic) {
			modifiers = append(modifiers, Modifier1_3x)
		}
	}

	// Defensive Terrains
	if isDefenderGrounded {
		if (field.Terrain == types.MistyTerrain && move.Type == types.Dragon) ||
			(field.Terrain == types.GrassyTerrain && (move.Name == "Earthquake" || move.Name == "Buldoze")) {
			modifiers = append(modifiers, Modifier0_5x)
		}
	}

	return modifiers
}

// getWeatherModifier returns the weather-based damage modifier for a move
func getWeatherModifier(weather types.Weather, moveType types.Type) int {
	switch weather {
	case types.Sun:
		if moveType == types.Fire {
			return Modifier1_5x
		}
		if moveType == types.Water {
			return Modifier0_5x
		}
	case types.Rain:
		if moveType == types.Fire {
			return Modifier0_5x
		}
		if moveType == types.Water {
			return Modifier1_5x
		}
	}
	return Modifier1x
}

func getTypeEffectiveness(move *types.Move, defender types.Pokemon, field types.Field) float64 {
	effectiveness := 1.0

	// If the move is Stellar type and the defender is a Tera type, always return super effective
	if defender.IsTera && move.Type == types.Stellar {
		return 2.0
	}

	// Check for immunities
	if (move.Type == types.Grass && defender.Ability == "Sap Sipper") ||
		(move.Type == types.Fire && slices.Contains([]string{"Well-Baked Body", "Flash Fire"}, defender.Ability)) ||
		(move.Type == types.Water && slices.Contains([]string{"Water Absord", "Dry Skin", "Storm Drain"}, defender.Ability)) ||
		(move.Type == types.Electric && slices.Contains([]string{"Lightning Rod", "Volt Absorb", "Motor Drive"}, defender.Ability)) ||
		(move.Type == types.Ground && ((!field.Gravity && defender.Item != "Iron Ball" && (defender.Ability == "Levitate" || defender.Item == "Air Balloon")) || defender.Ability == "Earth Eater")) ||
		(move.IsBullet() && defender.Ability == "Bulletproof") ||
		(move.IsSound() && defender.Ability == "Soundproof") ||
		(move.IsWind() && defender.Ability == "Wind Rider") {
		return 0
	}

	defenderTypes := defender.Type
	// If the defender is a Tera type, but not Stellar, override the defender's type to the Tera type
	if defender.IsTera && defender.Tera != types.Stellar {
		defenderTypes = []types.Type{defender.Tera}
	}

	// If the move is Ground, but the move ignores flying immunity, return the other types effectiveness
	if move.Type == types.Ground && (move.Name == "Thousand Arrows" || defender.Item == "Iron Ball" || field.Gravity) {

		for _, defenderType := range defenderTypes {
			if defenderType != types.Flying {
				effectiveness *= types.TypeChart[move.Type][defenderType]
			}
		}
	} else if move.Type == types.Ground && move.Name != "Thousand Arrows" && defender.Item != "Iron Ball" && !field.Gravity &&
		(defender.Ability == "Levitate" || defender.Item == "Air Balloon") {

		effectiveness = 0
	} else {
		// Calculate the effectiveness of the move
		for _, defenderType := range defenderTypes {
			effectiveness *= types.TypeChart[move.Type][defenderType]
		}
	}

	// Check for Wonder Guard
	if defender.Ability == "Wonder Guard" && effectiveness < 2 {
		effectiveness = 0
	}

	return effectiveness
}

func getStabModifier(attacker types.Pokemon, move types.Move) int {
	if attacker.IsTera {
		if attacker.Tera == types.Stellar {
			if slices.Contains(attacker.Type, move.Type) {
				// If STAB and Tera type is Stellar and attacker has Adaptability, return 2.25x
				if attacker.Ability == "Adaptability" {
					return Modifier2_25x
				}
				// If STAB and Tera type is Stellar, return 2x
				return Modifier2x
			}
			// If STAB and Tera type is Stellar, return 1.5x
			return Modifier1_5x
		}

		if move.Type == attacker.Tera {
			if slices.Contains(attacker.Type, move.Type) {
				// If STAB and Tera type is the same as the attacker's type, return 2x
				return Modifier2x
			}
			// If Tera is the same as the attacker's type but not STAB, return 1.5x
			return Modifier1_5x
		}
	}

	if slices.Contains(attacker.Type, move.Type) {
		// If STAB and attacker has Adaptability, return 2x
		if attacker.Ability == "Adaptability" {
			return Modifier2x
		}
		// If STAB, return 1.5x
		return Modifier1_5x
	}

	// else, no STAB then return 1x
	return Modifier1x
}
