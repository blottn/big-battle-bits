package bf

import (
	"image"
	"image/color"
)

type Armies map[Team][]image.Point

type Team color.Color

var (
	Black  = color.RGBA{0, 0, 0, 255}
	Pink   = color.RGBA{244, 3, 252, 255}
	Orange = color.RGBA{244, 3, 252, 255}
)
