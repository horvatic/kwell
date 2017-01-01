package imagePixels

import (
	"image/color"
)

func Average(colors [9]color.Color) color.Color {
	var r, g, b, n uint32
	for _, c := range colors {
		cr, cg, cb, _ := c.RGBA()
		if cr == 0 && cg == 0 && cb == 0 {
			continue
		}
		r += cr
		g += cg
		b += cb
		n++
	}
	if n == 0 {
		return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
	}
	return color.RGBA{uint8(r / n >> 8), uint8(g / n >> 8), uint8(b / n >> 8), 255}
}
