package imageProcessing

import (
	"image"
	"image/color"

	"github.com/horvatic/vaticwella/imagePixels"
)

func BinaryFilter(source image.Image, searchArea int) image.Image {
	img := image.NewRGBA(source.Bounds())
	bound := img.Bounds()
	for py := bound.Min.Y; py < bound.Max.Y; py++ {
		for px := bound.Min.X; px < bound.Max.X; px++ {
			avgcolor := imagePixels.Average(imagePixels.Around(source, px, py, searchArea))
			ar, ag, ab, _ := avgcolor.RGBA()
			sr, sg, sb, _ := source.At(px, py).RGBA()
			avgA := (ar + ag + ab) / 3
			avgS := (sr + sg + sb) / 3
			if avgS >= avgA {
				img.Set(px, py, color.RGBA{255, 255, 255, 255})
			} else {
				img.Set(px, py, color.RGBA{0, 0, 0, 255})
			}
		}
	}
	return img

}
