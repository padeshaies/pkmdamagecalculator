package types

import (
	"slices"
)

type Move struct {
	Name        string
	Type        Type
	DamageClass string
	Power       int
	Target      string
	CriticalHit bool

	isBullet  *bool
	isSound   *bool
	isWind    *bool
	isSlicing *bool
}

// source: https://bulbapedia.bulbagarden.net/wiki/Ball_and_bomb_move#List_of_ball_and_bomb_moves
func (m Move) IsBullet() bool {
	if m.isBullet == nil {
		*m.isBullet = slices.Contains([]string{
			"Acid Spray",
			"Aura Sphere",
			"Barrage",
			"Beak Blast",
			"Bullet Seed",
			"Egg Bomb",
			"Electro Ball",
			"Energy Ball",
			"Focus Blast",
			"Gyro Ball",
			"Ice Ball",
			"Magnet Bomb",
			"Mist Ball",
			"Mud Bomb",
			"Octazooka",
			"Pollen Puff",
			"Pyro Ball",
			"Rock Blast",
			"Searing Shot",
			"Seed Bomb",
			"Shadow Ball",
			"Sludge Bomb",
			"Syrup Bomb",
			"Weather Ball",
			"Zap Cannon",
		}, m.Name)
	}
	return *m.isBullet
}

// source: https://bulbapedia.bulbagarden.net/wiki/Sound-based_move#List_of_sound-based_moves
func (m Move) IsSound() bool {
	if m.isSound == nil {
		*m.isSound = slices.Contains([]string{
			"Alluring Voice",
			"Boomburst",
			"Bug Buzz",
			"Chatter",
			"Clanging Scales",
			"Disarming Voice",
			"Echoed Voice",
			"Eerie Spell",
			"Hyper Voice",
			"Overdrive",
			"Psychic Noise",
			"Relic Song",
			"Round",
			"Snarl",
			"Snore",
			"Sparkling Aria",
			"Torch Song",
			"Uproar",
		}, m.Name)
	}
	return *m.isSound
}

// source: https://bulbapedia.bulbagarden.net/wiki/Wind_move#List_of_wind_moves
func (m Move) IsWind() bool {
	if m.isWind == nil {
		*m.isWind = slices.Contains([]string{
			"Aeroblast",
			"Air Cutter",
			"Bleakwind Storm",
			"Blizzard",
			"Fairy Wind",
			"Gust",
			"Heat Wave",
			"Hurricane",
			"Icy Wind",
			"Petal Blizzard",
			"Sandsear Storm",
			"Springtide Storm",
			"Twister",
			"Wildbolt Storm",
		}, m.Name)
	}
	return *m.isWind
}

// source: https://bulbapedia.bulbagarden.net/wiki/Slicing_move#List_of_slicing_moves
func (m Move) IsSlicing() bool {
	if m.isSlicing == nil {
		*m.isSlicing = slices.Contains([]string{
			"Aerial Ace",
			"Air Cutter",
			"Air Slash",
			"Aqua Cutter",
			"Behemoth Blade",
			"Bitter Blade",
			"Ceaseless Edge",
			"Cross Poison",
			"Cut",
			"Fury Cutter",
			"Kowtow Cleave",
			"Leaf Blade",
			"Mighty Cleave",
			"Night Slash",
			"Population Bomb",
			"Psypblade",
			"Psycho Cut",
			"Razor Leaf",
			"Razor Shell",
			"Sacred Sword",
			"Secret Sword",
			"Slash",
			"Solar Blade",
			"Stone Axe",
			"Tachyon Cutter",
			"X-Scissor",
		}, m.Name)
	}
	return *m.isSlicing
}
