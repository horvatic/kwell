package imageProcessing

import (
	"image"
	"image/color"

	"github.com/horvatic/vaticwella/src/imagePixels"
)

func binaryFilterNoParallel(source image.Image, searchArea int) image.Image {
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

func binaryFilterRunInParallel(source image.Image, searchArea int) image.Image {
	img := image.NewRGBA(source.Bounds())
	bound := img.Bounds()
	done := make(chan bool, (bound.Max.Y - bound.Min.Y - 1))
	for py := bound.Min.Y; py < bound.Max.Y; py++ {
		go func(cpy int) {
			for px := bound.Min.X; px < bound.Max.X; px++ {
				avgcolor := imagePixels.Average(imagePixels.Around(source, px, cpy, searchArea))
				ar, ag, ab, _ := avgcolor.RGBA()
				sr, sg, sb, _ := source.At(px, cpy).RGBA()
				avgA := (ar + ag + ab) / 3
				avgS := (sr + sg + sb) / 3
				if avgS >= avgA {
					img.Set(px, cpy, color.RGBA{255, 255, 255, 255})
				} else {
					img.Set(px, cpy, color.RGBA{0, 0, 0, 255})
				}
			}
			done <- true
		}(py)
	}

	for py := bound.Min.Y; py < bound.Max.Y; py++ {
		<-done
	}
	close(done)
	return img
}

func BinaryFilter(source image.Image, searchArea int, runInParallel bool) image.Image {
	if runInParallel {
		return binaryFilterRunInParallel(source, searchArea)
	}
	return binaryFilterNoParallel(source, searchArea)
}
