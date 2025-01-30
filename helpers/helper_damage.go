package helpers

import (
	"math"
	"slices"

	"github.com/padeshaies/pkmdamagecalculator/types"
)

func CalculateDamage(attacker types.Pokemon, defender types.Pokemon, move types.Move) []int {
	// Calculate the damage dealt by the attacker to the defender using the given move
	damage := make([]int, 16)

	// Base damage calculation
	var a, d int
	switch move.DamageClass {
	case "physical":
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
		baseDamage = pokeRound(baseDamage, 0.75)
	default:
		return damage // Not targeting a Pokemon
	}

	// TODO: Apply the weather modifier

	// Apply the critical hit modifier
	if move.CriticalHit && defender.Ability != "Battle Armor" && defender.Ability != "Shell Armor" {
		baseDamage = pokeRound(baseDamage, 1.5)
	}

	// Store the STAB modifier
	stab := 1.0
	if slices.Contains(attacker.Type, move.Type) {
		// STAB bonus is 1.5x
		stab = 1.5
	}

	// Store the type effectiveness
	typeEffectiveness := typeChart[move.Type][defender.Type[0]]
	if len(defender.Type) > 1 {
		typeEffectiveness = typeEffectiveness * typeChart[move.Type][defender.Type[1]]
	}
	if defender.Ability == "Wonder Guard" && typeEffectiveness < 2 {
		typeEffectiveness = 0
	}
	if typeEffectiveness == 0 {
		return damage // Defender is immune to the move
	}

	// Is Burned
	isBurned := attacker.Status == "burn"
	//if attacker.Status == "burn" && move.DamageClass == "physical" && attacker.Ability != "Guts" {
	//	baseDamage = int16(float64(baseDamage) * 0.5)
	//}

	// add stuff to damage (ie: abilities, items, etc)

	// Add randomization and chain the modifiers ('cause turns out that switching between floors and rounds is a pain 🙃)
	for i, rand := range []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0} {
		tempDamage := int(baseDamage * (100 - rand) / 100)
		tempDamage = pokeRound(tempDamage, stab)
		tempDamage = int(float64(tempDamage) * typeEffectiveness)
		if isBurned {
			tempDamage = int(float64(tempDamage) * 0.5)
		}
		// TODO apply final modifier
		damage[i] = max(1, tempDamage)
	}
	return damage
}

var typeChart = map[string]map[string]float64{
	"normal": {
		"rock":     0.5,
		"ghost":    0,
		"steel":    0.5,
		"normal":   1,
		"fire":     1,
		"water":    1,
		"grass":    1,
		"electric": 1,
		"ice":      1,
		"fighting": 1,
		"poison":   1,
		"ground":   1,
		"flying":   1,
		"psychic":  1,
		"bug":      1,
		"dragon":   1,
		"dark":     1,
		"fairy":    1,
	},
	"fire": {
		"fire":     0.5,
		"water":    0.5,
		"grass":    2,
		"ice":      2,
		"bug":      2,
		"rock":     0.5,
		"dragon":   0.5,
		"steel":    2,
		"normal":   1,
		"electric": 1,
		"fighting": 1,
		"poison":   1,
		"ground":   1,
		"flying":   1,
		"psychic":  1,
		"ghost":    1,
		"dark":     1,
		"fairy":    1,
	},
	"water": {
		"fire":     2,
		"water":    0.5,
		"grass":    0.5,
		"ground":   2,
		"rock":     2,
		"dragon":   0.5,
		"normal":   1,
		"electric": 1,
		"ice":      1,
		"fighting": 1,
		"poison":   1,
		"flying":   1,
		"psychic":  1,
		"bug":      1,
		"ghost":    1,
		"dark":     1,
		"steel":    1,
		"fairy":    1,
	},
	"grass": {
		"fire":     0.5,
		"water":    2,
		"grass":    0.5,
		"poison":   0.5,
		"ground":   2,
		"flying":   0.5,
		"bug":      0.5,
		"rock":     2,
		"dragon":   0.5,
		"steel":    0.5,
		"normal":   1,
		"electric": 1,
		"ice":      1,
		"fighting": 1,
		"psychic":  1,
		"ghost":    1,
		"dark":     1,
		"fairy":    1,
	},
	"electric": {
		"water":    2,
		"grass":    0.5,
		"electric": 0.5,
		"ground":   0,
		"flying":   2,
		"dragon":   0.5,
		"normal":   1,
		"fire":     1,
		"ice":      1,
		"fighting": 1,
		"poison":   1,
		"bug":      1,
		"rock":     1,
		"psychic":  1,
		"ghost":    1,
		"dark":     1,
		"steel":    1,
		"fairy":    1,
	},
	"ice": {
		"fire":     0.5,
		"water":    0.5,
		"grass":    2,
		"ice":      0.5,
		"ground":   2,
		"flying":   2,
		"dragon":   2,
		"steel":    0.5,
		"normal":   1,
		"electric": 1,
		"fighting": 1,
		"poison":   1,
		"bug":      1,
		"rock":     1,
		"psychic":  1,
		"ghost":    1,
		"dark":     1,
		"fairy":    1,
	},
	"fighting": {
		"normal":   2,
		"ice":      2,
		"poison":   0.5,
		"flying":   0.5,
		"psychic":  0.5,
		"bug":      0.5,
		"rock":     2,
		"ghost":    0,
		"dark":     2,
		"steel":    2,
		"fairy":    0.5,
		"fire":     1,
		"water":    1,
		"grass":    1,
		"electric": 1,
		"ground":   1,
		"dragon":   1,
	},
	"poison": {
		"grass":    2,
		"poison":   0.5,
		"ground":   0.5,
		"rock":     0.5,
		"ghost":    0.5,
		"steel":    0,
		"fairy":    2,
		"normal":   1,
		"fire":     1,
		"water":    1,
		"electric": 1,
		"ice":      1,
		"flying":   1,
		"psychic":  1,
		"bug":      1,
		"dragon":   1,
		"dark":     1,
	},
	"ground": {
		"fire":     2,
		"electric": 2,
		"grass":    0.5,
		"poison":   2,
		"flying":   0,
		"bug":      0.5,
		"rock":     2,
		"steel":    2,
		"normal":   1,
		"water":    1,
		"ice":      1,
		"fighting": 1,
		"psychic":  1,
		"ghost":    1,
		"dragon":   1,
		"dark":     1,
		"fairy":    1,
	},
	"flying": {
		"electric": 0.5,
		"grass":    2,
		"fighting": 2,
		"bug":      2,
		"rock":     0.5,
		"steel":    0.5,
		"normal":   1,
		"fire":     1,
		"water":    1,
		"ice":      1,
		"poison":   1,
		"ground":   1,
		"psychic":  1,
		"ghost":    1,
		"dragon":   1,
		"dark":     1,
		"fairy":    1,
	},
	"psychic": {
		"fighting": 2,
		"poison":   2,
		"psychic":  0.5,
		"dark":     0,
		"steel":    0.5,
		"normal":   1,
		"fire":     1,
		"water":    1,
		"grass":    1,
		"electric": 1,
		"ice":      1,
		"ground":   1,
		"flying":   1,
		"bug":      1,
		"rock":     1,
		"ghost":    1,
		"dragon":   1,
		"fairy":    1,
	},
	"bug": {
		"fire":     0.5,
		"grass":    2,
		"fighting": 0.5,
		"flying":   0.5,
		"poison":   0.5,
		"ghost":    0.5,
		"steel":    0.5,
		"fairy":    0.5,
		"psychic":  2,
		"dark":     2,
		"normal":   1,
		"water":    1,
		"electric": 1,
		"ice":      1,
		"ground":   1,
		"rock":     1,
		"dragon":   1,
	},
	"rock": {
		"fire":     2,
		"ice":      2,
		"fighting": 0.5,
		"ground":   0.5,
		"flying":   2,
		"bug":      2,
		"steel":    0.5,
		"normal":   1,
		"water":    1,
		"grass":    1,
		"electric": 1,
		"poison":   1,
		"psychic":  1,
		"ghost":    1,
		"dragon":   1,
		"dark":     1,
		"fairy":    1,
	},
	"ghost": {
		"normal":   0,
		"psychic":  2,
		"ghost":    2,
		"dark":     0.5,
		"fire":     1,
		"water":    1,
		"grass":    1,
		"electric": 1,
		"ice":      1,
		"fighting": 1,
		"poison":   1,
		"ground":   1,
		"flying":   1,
		"bug":      1,
		"rock":     1,
		"dragon":   1,
		"steel":    1,
		"fairy":    1,
	},
	"dragon": {
		"dragon":   2,
		"steel":    0.5,
		"fairy":    0,
		"normal":   1,
		"fire":     1,
		"water":    1,
		"grass":    1,
		"electric": 1,
		"ice":      1,
		"fighting": 1,
		"poison":   1,
		"ground":   1,
		"flying":   1,
		"psychic":  1,
		"bug":      1,
		"rock":     1,
		"ghost":    1,
		"dark":     1,
	},
	"dark": {
		"fighting": 0.5,
		"psychic":  2,
		"ghost":    2,
		"dark":     0.5,
		"fairy":    0.5,
		"normal":   1,
		"fire":     1,
		"water":    1,
		"grass":    1,
		"electric": 1,
		"ice":      1,
		"poison":   1,
		"ground":   1,
		"flying":   1,
		"bug":      1,
		"rock":     1,
		"dragon":   1,
		"steel":    1,
	},
	"steel": {
		"fire":     0.5,
		"water":    0.5,
		"electric": 0.5,
		"ice":      2,
		"rock":     2,
		"steel":    0.5,
		"fairy":    2,
		"normal":   1,
		"grass":    1,
		"fighting": 1,
		"poison":   1,
		"ground":   1,
		"flying":   1,
		"psychic":  1,
		"bug":      1,
		"ghost":    1,
		"dragon":   1,
		"dark":     1,
	},
	"fairy": {
		"fire":     0.5,
		"fighting": 2,
		"poison":   0.5,
		"dragon":   2,
		"dark":     2,
		"steel":    0.5,
		"normal":   1,
		"water":    1,
		"grass":    1,
		"electric": 1,
		"ice":      1,
		"ground":   1,
		"flying":   1,
		"psychic":  1,
		"bug":      1,
		"rock":     1,
		"ghost":    1,
	},
}

// Pokemon's way of applying damage modifiers
func pokeRound(damage int, modifier float64) int {
	modifiedDamage := float64(damage) * modifier
	_, decimal := math.Modf(modifiedDamage)
	if decimal <= 0.5 {
		return int(modifiedDamage)
	} else {
		return int(modifiedDamage) + 1
	}
}
