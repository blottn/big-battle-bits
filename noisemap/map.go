package noisemap

import (
	"image"
	"image/color"
	"math"

	noise "github.com/aquilax/go-perlin"
)

type NoiseMap map[image.Point]float64

func NewNoiseMap(w, h int, noiseSrc *noise.Perlin) NoiseMap {
	nm := NoiseMap{}
	min := 99.0
	max := -99.0
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			v := noiseSrc.Noise2D(float64(i)/float64(w), float64(j)/float64(h))
			min = math.Min(v, min)
			max = math.Max(v, max)
			nm[image.Point{i, j}] = v
		}
	}

	for point, val := range nm {
		nm[point] = (val + math.Abs(min)) / (max - min)
	}
	return nm
}

func (nm NoiseMap) ColorModel() color.Model {
	return color.GrayModel
}

func (nm NoiseMap) Bounds() image.Rectangle {
	maxX := 0.0
	maxY := 0.0
	for p, _ := range nm {
		maxX = math.Max(maxX, float64(p.X))
		maxY = math.Max(maxY, float64(p.Y))
	}
	return image.Rect(0, 0, int(maxX), int(maxY))
}
func (nm NoiseMap) At(x, y int) color.Color {
	return color.Gray{uint8(nm[image.Point{x, y}] * 255)}
}
