package bf

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"

	"big-battle-bits/noisemap"

	noise "github.com/aquilax/go-perlin"
)

var (
	Ocean = color.RGBA{0, 68, 128, 255}
	Land  = color.RGBA{0, 155, 0, 255}
)

type BattleGround struct {
	width   int
	height  int
	terrain image.Image
	armies  Armies
}

func (bg BattleGround) Width() int {
	return bg.width
}

func (bg BattleGround) Height() int {
	return bg.height
}

func centerDist(x, y, width, height int) float64 {
	xDif := float64(x)/float64(width) - 0.5
	yDif := float64(y)/float64(height) - 0.5
	return math.Sqrt(xDif*xDif + yDif*yDif)
}

func (bg BattleGround) forceColor(x, y int) color.Color {
	team, ok := bg.armies[image.Point{x, y}]
	if !ok {
		return nil
	}
	return team.Color
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
		Armies{},
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
	armies := Armies{}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if compareColor(Ocean, img.At(i, j)) {
				terrain.Set(i, j, Ocean)
			} else if compareColor(Land, img.At(i, j)) {
				terrain.Set(i, j, Land)
			} else {
				r, g, b, a := img.At(i, j).RGBA()
				armies[image.Point{i, j}] = Team{color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}}
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

func (bg BattleGround) Add(team Team, x, y int) error {
	return bg.AddAtPoint(team, image.Point{x, y})
}

func (bg BattleGround) AddAtPoint(team Team, p image.Point) error {
	if p.X > bg.width || p.Y > bg.height {
		return fmt.Errorf("Error, tried to add team out of bounds: %v to %v",
			p,
			bg.terrain.Bounds(),
		)
	}
	if _, ok := bg.armies[p]; ok {
		return fmt.Errorf("Error adding team to %v, already owned by another team")
	}
	if bg.IsOcean(p) {
		return fmt.Errorf("Error adding team to %v, it's in the water", p)
	}
	bg.armies[p] = team
	return nil
}

func (bg BattleGround) IsOcean(p image.Point) bool {
	return bg.terrain.At(p.X, p.Y) == Ocean
}

func (bg BattleGround) Output(writer io.Writer) error {
	encoder := &png.Encoder{png.BestSpeed, nil}
	return encoder.Encode(writer, bg)
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
