package main

import (
	"big-battle-bits/noisemap"
	"image/png"
	"os"

	noise "github.com/aquilax/go-perlin"
)

func main() {
	n := noise.NewPerlin(1.1, 2, 3, 100)
	nm := noisemap.NewNoiseMap(999, 999, n)
	file, _ := os.OpenFile("noisemap.png", os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	png.Encode(file, nm)
}
