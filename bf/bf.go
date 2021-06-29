package bf

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"

	"bbb/noisemap"

	noise "github.com/aquilax/go-perlin"
)

type BattleGround struct {
	width  int
	height int
	nm     noisemap.NoiseMap
	armies map[image.Point]Force
}

var Ocean = color.RGBA{0, 68, 128, 255}
var Land = color.RGBA{0, 155, 0, 255}

func centerDist(x, y, width, height int) float64 {
	xDif := float64(x)/float64(width) - 0.5
	yDif := float64(y)/float64(height) - 0.5
	return math.Sqrt(xDif*xDif + yDif*yDif)
}

func (bg BattleGround) terrainColor(x, y int) color.Color {
	dist := centerDist(x, y, bg.width, bg.height)
	n := bg.nm[image.Point{x, y}]
	if n > dist*0.8+0.3 { // Please don't touch these magic numbers
		return Land
	}
	return nil
}

func (bg BattleGround) forceColor(x, y int) color.Color {
	for point, force := range bg.armies {
		if point.X == x && point.Y == y {
			return force.Color
		}
	}
	return nil
}

func NewBattleGround(width, height, seed int) BattleGround {
	return BattleGround{
		width,
		height,
		noisemap.NewNoiseMap(width, height, noise.NewPerlin(1, 2, 3, int64(seed))),
		nil,
	}
}

func (bg BattleGround) Output(writer io.Writer) error {
	return png.Encode(writer, bg)
}

// image.Image interface
func (bg BattleGround) ColorModel() color.Model {
	return color.RGBAModel
}

func (bg BattleGround) Bounds() image.Rectangle {
	return image.Rect(0, 0, bg.width, bg.height)
}

func (bg BattleGround) At(x, y int) color.Color {
	if fc := bg.forceColor(x, y); fc != nil {
		return fc
	}

	if tc := bg.terrainColor(x, y); tc != nil {
		return tc
	}

	return Ocean
}
