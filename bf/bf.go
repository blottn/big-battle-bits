package bf

import (
	"image"
	"image/color"
	"image/png"

	"io"

	noise "github.com/aquilax/go-perlin"
)

type BattleGround struct {
	img         image.Image
	noiseSource *noise.Perlin
}

func centerDist(x, y, width, height int) float64 {
	halfWidth := float64(width) / 2.0
	halfHeight := float64(height) / 2.0
	xDif := float64(x) - halfWidth
	yDif := float64(y) - halfHeight
	return xDif*xDif + yDif*yDif
}

func NewBattleGround(width, height, seed int) BattleGround {
	perlin := noise.NewPerlin(2, 2, 3, int64(seed))
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			d := centerDist(i, j, width, height) / float64(width)
			n := perlin.Noise2D(float64(i)/float64(width), float64(j)/float64(height))
			if d*0.007-0.2 < n {
				img.Set(i, j, color.Black)
			} else {
				img.Set(i, j, color.White)
			}
		}
	}
	// TODO force add some islands
	return BattleGround{img, perlin}
}

func (bg BattleGround) Output(writer io.Writer) error {
	return png.Encode(writer, bg.img)
}
