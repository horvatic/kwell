package imagePixels

import (
	"image"
	"image/color"
)

func Around(source image.Image, px, py int) [9]color.Color {
	var colors [9]color.Color
	colors[0] = source.At(px-1, py-1)
	colors[1] = source.At(px, py-1)
	colors[2] = source.At(px+1, py-1)
	colors[3] = source.At(px-1, py)
	colors[4] = source.At(px, py)
	colors[5] = source.At(px+1, py)
	colors[6] = source.At(px-1, py+1)
	colors[7] = source.At(px, py+1)
	colors[8] = source.At(px+1, py+1)
	return colors
}
