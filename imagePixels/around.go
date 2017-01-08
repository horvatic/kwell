package imagePixels

import (
	"image"
	"image/color"
)

func Around(source image.Image, px, py int) []color.Color {
	var colors []color.Color
	colors = append(colors, source.At(px-1, py-1))
	colors = append(colors, source.At(px-1, py))
	colors = append(colors, source.At(px-1, py+1))
	colors = append(colors, source.At(px, py-1))
	colors = append(colors, source.At(px, py))
	colors = append(colors, source.At(px, py+1))
	colors = append(colors, source.At(px+1, py-1))
	colors = append(colors, source.At(px+1, py))
	colors = append(colors, source.At(px+1, py+1))
	return colors
}
