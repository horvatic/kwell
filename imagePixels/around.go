package imagePixels

import (
	"image"
	"image/color"
)

func Around(source image.Image, px, py, searchArea int) []color.Color {
	var colors []color.Color
	for x := -searchArea; x <= searchArea; x++ {
		for y := -searchArea; y <= searchArea; y++ {
			colors = append(colors, source.At(px+x, py+y))
		}
	}
	return colors
}
