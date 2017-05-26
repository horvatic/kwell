package imageProcessing

import (
	"github.com/horvatic/vaticwella/imagePixels"
	"image"
)

func SuperSampling(source image.Image, searchArea int) image.Image {
	img := image.NewRGBA(source.Bounds())
	for py := 0; py < img.Rect.Dy(); py++ {
		for px := 0; px < img.Rect.Dx(); px++ {
			img.Set(px, py, imagePixels.Average(imagePixels.Around(source, px, py, searchArea)))
		}
	}
	return img
}
