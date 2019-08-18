package imageProcessing

import (
	"image"

	"github.com/horvatic/vaticwella/src/imagePixels"
)

func medianFilterRunInParallel(source image.Image, searchArea int) image.Image {
	img := image.NewRGBA(source.Bounds())
	bound := img.Bounds()
	done := make(chan bool, (bound.Max.Y - bound.Min.Y - 1))
	for py := bound.Min.Y; py < bound.Max.Y; py++ {
		go func(cpy int) {
			for px := bound.Min.X; px < bound.Max.X; px++ {
				img.Set(px, cpy, imagePixels.GetColorsMedian((imagePixels.Around(source, px, cpy, searchArea))))
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

func medianFilterNoParallel(source image.Image, searchArea int) image.Image {
	img := image.NewRGBA(source.Bounds())
	bound := img.Bounds()
	for py := bound.Min.Y; py < bound.Max.Y; py++ {
		for px := bound.Min.X; px < bound.Max.X; px++ {
			img.Set(px, py, imagePixels.GetColorsMedian((imagePixels.Around(source, px, py, searchArea))))
		}
	}
	return img
}

func MedianFilter(source image.Image, searchArea int, runInParallel bool) image.Image {
	if runInParallel {
		return medianFilterRunInParallel(source, searchArea)
	}
	return medianFilterNoParallel(source, searchArea)
}
