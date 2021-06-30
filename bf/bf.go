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
	width   int
	height  int
	terrain image.Image
	armies  Armies
}

var Ocean = color.RGBA{0, 68, 128, 255}
var Land = color.RGBA{0, 155, 0, 255}

func centerDist(x, y, width, height int) float64 {
	xDif := float64(x)/float64(width) - 0.5
	yDif := float64(y)/float64(height) - 0.5
	return math.Sqrt(xDif*xDif + yDif*yDif)
}

func (bg BattleGround) forceColor(x, y int) color.Color {
	for point, force := range bg.armies {
		if point.X == x && point.Y == y {
			return force.Color
		}
	}
	return nil
}

func getTerrain(x, y, w, h int, nm noisemap.NoiseMap) color.Color {
	dist := centerDist(x, y, w, h)
	n := nm[image.Point{x, y}]
	if n > dist*0.8+0.3 {
		return Land
	}
	return Ocean
}

func newTerrain(width, height int, nm noisemap.NoiseMap) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			img.Set(i, j, getTerrain(i, j, width, height, nm))
		}
	}
	return img
}

func NewBattleGround(width, height, seed int) BattleGround {
	nm := noisemap.NewNoiseMap(width, height, noise.NewPerlin(1, 2, 3, int64(seed)))
	return BattleGround{
		width,
		height,
		newTerrain(width, height, nm),
		nil,
	}
}

func compareColor(left, right color.Color) bool {
	left_r, left_g, left_b, left_a := left.RGBA()
	right_r, right_g, right_b, right_a := right.RGBA()
	return left_r == right_r &&
		left_g == right_g &&
		left_b == right_b &&
		left_a == right_a
}

func From(reader io.Reader) (*BattleGround, error) {
	// Collect armies
	img, err := png.Decode(reader)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	terrain := image.NewRGBA(bounds)
	forces := Armies{}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if compareColor(Ocean, img.At(i, j)) {
				terrain.Set(i, j, Ocean)
			} else if compareColor(Land, img.At(i, j)) {
				terrain.Set(i, j, Land)
			} else {
				col := img.At(i, j)
				points, ok := armies[col]
				if !ok {
					armies[col] = []image.Point{}
				}
				armies[col] = append(armies[col], image.Point(i, j))

			}
		}
	}

	return &BattleGround{
		width,
		height,
		terrain,
		armies,
	}, nil

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

	return bg.terrain.At(x, y)
}
