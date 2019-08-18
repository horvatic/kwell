package imageProcessing

import (
	"image"

	"github.com/horvatic/vaticwella/src/imagePixels"
)

func MedianFilter(source image.Image, searchArea int) image.Image {
	img := image.NewRGBA(source.Bounds())
	bound := img.Bounds()
	for py := bound.Min.Y; py < bound.Max.Y; py++ {
		for px := bound.Min.X; px < bound.Max.X; px++ {
			img.Set(px, py, imagePixels.GetColorsMedian((imagePixels.Around(source, px, py, searchArea))))
		}
	}
	return img

}
