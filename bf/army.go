package bf

import (
	"image"
	"image/color"
)

type Army map[image.Point]Force

type Force struct {
	Color color.Color
}
