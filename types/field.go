package types

type Field struct {
	Weather Weather
	Terrain Terrain

	Side1 Side
	Side2 Side
}

type Side struct {
	HelpingHand bool
}

type Terrain string

const (
	ElectricTerrain Terrain = "electric"
	GrassyTerrain   Terrain = "grassy"
	MistyTerrain    Terrain = "misty"
	PsychicTerrain  Terrain = "psychic"
)

type Weather string

const (
	Sun  Weather = "sun"
	Rain Weather = "rain"
	Sand Weather = "sand"
	Snow Weather = "snow"
)
