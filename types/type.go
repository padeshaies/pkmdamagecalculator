package types

type Type string

const (
	Normal   Type = "normal"
	Fighting Type = "fighting"
	Flying   Type = "flying"
	Poison   Type = "poison"
	Ground   Type = "ground"
	Rock     Type = "rock"
	Bug      Type = "bug"
	Ghost    Type = "ghost"
	Steel    Type = "steel"
	Fire     Type = "fire"
	Water    Type = "water"
	Grass    Type = "grass"
	Electric Type = "electric"
	Psychic  Type = "psychic"
	Ice      Type = "ice"
	Dragon   Type = "dragon"
	Dark     Type = "dark"
	Fairy    Type = "fairy"

	Stellar Type = "stellar"
)

var TypeChart = map[Type]map[Type]float64{
	Normal: {
		Rock:     0.5,
		Ghost:    0,
		Steel:    0.5,
		Normal:   1,
		Fire:     1,
		Water:    1,
		Grass:    1,
		Electric: 1,
		Ice:      1,
		Fighting: 1,
		Poison:   1,
		Ground:   1,
		Flying:   1,
		Psychic:  1,
		Bug:      1,
		Dragon:   1,
		Dark:     1,
		Fairy:    1,
	},
	Fire: {
		Fire:     0.5,
		Water:    0.5,
		Grass:    2,
		Ice:      2,
		Bug:      2,
		Rock:     0.5,
		Dragon:   0.5,
		Steel:    2,
		Normal:   1,
		Electric: 1,
		Fighting: 1,
		Poison:   1,
		Ground:   1,
		Flying:   1,
		Psychic:  1,
		Ghost:    1,
		Dark:     1,
		Fairy:    1,
	},
	Water: {
		Fire:     2,
		Water:    0.5,
		Grass:    0.5,
		Ground:   2,
		Rock:     2,
		Dragon:   0.5,
		Normal:   1,
		Electric: 1,
		Ice:      1,
		Fighting: 1,
		Poison:   1,
		Flying:   1,
		Psychic:  1,
		Bug:      1,
		Ghost:    1,
		Dark:     1,
		Steel:    1,
		Fairy:    1,
	},
	Grass: {
		Fire:     0.5,
		Water:    2,
		Grass:    0.5,
		Poison:   0.5,
		Ground:   2,
		Flying:   0.5,
		Bug:      0.5,
		Rock:     2,
		Dragon:   0.5,
		Steel:    0.5,
		Normal:   1,
		Electric: 1,
		Ice:      1,
		Fighting: 1,
		Psychic:  1,
		Ghost:    1,
		Dark:     1,
		Fairy:    1,
	},
	Electric: {
		Water:    2,
		Grass:    0.5,
		Electric: 0.5,
		Ground:   0,
		Flying:   2,
		Dragon:   0.5,
		Normal:   1,
		Fire:     1,
		Ice:      1,
		Fighting: 1,
		Poison:   1,
		Bug:      1,
		Rock:     1,
		Psychic:  1,
		Ghost:    1,
		Dark:     1,
		Steel:    1,
		Fairy:    1,
	},
	Ice: {
		Fire:     0.5,
		Water:    0.5,
		Grass:    2,
		Ice:      0.5,
		Ground:   2,
		Flying:   2,
		Dragon:   2,
		Steel:    0.5,
		Normal:   1,
		Electric: 1,
		Fighting: 1,
		Poison:   1,
		Bug:      1,
		Rock:     1,
		Psychic:  1,
		Ghost:    1,
		Dark:     1,
		Fairy:    1,
	},
	Fighting: {
		Normal:   2,
		Ice:      2,
		Poison:   0.5,
		Flying:   0.5,
		Psychic:  0.5,
		Bug:      0.5,
		Rock:     2,
		Ghost:    0,
		Dark:     2,
		Steel:    2,
		Fairy:    0.5,
		Fire:     1,
		Water:    1,
		Grass:    1,
		Electric: 1,
		Ground:   1,
		Dragon:   1,
	},
	Poison: {
		Grass:    2,
		Poison:   0.5,
		Ground:   0.5,
		Rock:     0.5,
		Ghost:    0.5,
		Steel:    0,
		Fairy:    2,
		Normal:   1,
		Fire:     1,
		Water:    1,
		Electric: 1,
		Ice:      1,
		Flying:   1,
		Psychic:  1,
		Bug:      1,
		Dragon:   1,
		Dark:     1,
	},
	Ground: {
		Fire:     2,
		Electric: 2,
		Grass:    0.5,
		Poison:   2,
		Flying:   0,
		Bug:      0.5,
		Rock:     2,
		Steel:    2,
		Normal:   1,
		Water:    1,
		Ice:      1,
		Fighting: 1,
		Psychic:  1,
		Ghost:    1,
		Dragon:   1,
		Dark:     1,
		Fairy:    1,
	},
	Flying: {
		Electric: 0.5,
		Grass:    2,
		Fighting: 2,
		Bug:      2,
		Rock:     0.5,
		Steel:    0.5,
		Normal:   1,
		Fire:     1,
		Water:    1,
		Ice:      1,
		Poison:   1,
		Ground:   1,
		Psychic:  1,
		Ghost:    1,
		Dragon:   1,
		Dark:     1,
		Fairy:    1,
	},
	Psychic: {
		Fighting: 2,
		Poison:   2,
		Psychic:  0.5,
		Dark:     0,
		Steel:    0.5,
		Normal:   1,
		Fire:     1,
		Water:    1,
		Grass:    1,
		Electric: 1,
		Ice:      1,
		Ground:   1,
		Flying:   1,
		Bug:      1,
		Rock:     1,
		Ghost:    1,
		Dragon:   1,
		Fairy:    1,
	},
	Bug: {
		Fire:     0.5,
		Grass:    2,
		Fighting: 0.5,
		Flying:   0.5,
		Poison:   0.5,
		Ghost:    0.5,
		Steel:    0.5,
		Fairy:    0.5,
		Psychic:  2,
		Dark:     2,
		Normal:   1,
		Water:    1,
		Electric: 1,
		Ice:      1,
		Ground:   1,
		Rock:     1,
		Dragon:   1,
	},
	Rock: {
		Fire:     2,
		Ice:      2,
		Fighting: 0.5,
		Ground:   0.5,
		Flying:   2,
		Bug:      2,
		Steel:    0.5,
		Normal:   1,
		Water:    1,
		Grass:    1,
		Electric: 1,
		Poison:   1,
		Psychic:  1,
		Ghost:    1,
		Dragon:   1,
		Dark:     1,
		Fairy:    1,
	},
	Ghost: {
		Normal:   0,
		Psychic:  2,
		Ghost:    2,
		Dark:     0.5,
		Fire:     1,
		Water:    1,
		Grass:    1,
		Electric: 1,
		Ice:      1,
		Fighting: 1,
		Poison:   1,
		Ground:   1,
		Flying:   1,
		Bug:      1,
		Rock:     1,
		Dragon:   1,
		Steel:    1,
		Fairy:    1,
	},
	Dragon: {
		Dragon:   2,
		Steel:    0.5,
		Fairy:    0,
		Normal:   1,
		Fire:     1,
		Water:    1,
		Grass:    1,
		Electric: 1,
		Ice:      1,
		Fighting: 1,
		Poison:   1,
		Ground:   1,
		Flying:   1,
		Psychic:  1,
		Bug:      1,
		Rock:     1,
		Ghost:    1,
		Dark:     1,
	},
	Dark: {
		Fighting: 0.5,
		Psychic:  2,
		Ghost:    2,
		Dark:     0.5,
		Fairy:    0.5,
		Normal:   1,
		Fire:     1,
		Water:    1,
		Grass:    1,
		Electric: 1,
		Ice:      1,
		Poison:   1,
		Ground:   1,
		Flying:   1,
		Bug:      1,
		Rock:     1,
		Dragon:   1,
		Steel:    1,
	},
	Steel: {
		Fire:     0.5,
		Water:    0.5,
		Electric: 0.5,
		Ice:      2,
		Rock:     2,
		Steel:    0.5,
		Fairy:    2,
		Normal:   1,
		Grass:    1,
		Fighting: 1,
		Poison:   1,
		Ground:   1,
		Flying:   1,
		Psychic:  1,
		Bug:      1,
		Ghost:    1,
		Dragon:   1,
		Dark:     1,
	},
	Fairy: {
		Fire:     0.5,
		Fighting: 2,
		Poison:   0.5,
		Dragon:   2,
		Dark:     2,
		Steel:    0.5,
		Normal:   1,
		Water:    1,
		Grass:    1,
		Electric: 1,
		Ice:      1,
		Ground:   1,
		Flying:   1,
		Psychic:  1,
		Bug:      1,
		Rock:     1,
		Ghost:    1,
	},
}
