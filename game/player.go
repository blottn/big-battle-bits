package game

import (
	"big-battle-bits/bf"
	"image"
	"image/color"
)

type PlayerConfigs map[string]PlayerConfig

type PlayerConfig struct {
	Priority bf.Prioritiser
	Start    image.Point
	Color    color.Color
}

func NewDefaultPlayerConfig() PlayerConfig {
	return PlayerConfig{
		bf.Angle(0),
		image.Point{0, 0},
		color.Black,
	}
}
