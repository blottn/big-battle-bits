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
