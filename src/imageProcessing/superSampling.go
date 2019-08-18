package imageProcessing

import (
	"image"

	"github.com/horvatic/vaticwella/src/imagePixels"
)

func superSamplingRunInParallel(source image.Image, searchArea int) image.Image {
	img := image.NewRGBA(source.Bounds())
	bound := img.Bounds()
	done := make(chan bool, (bound.Max.Y - bound.Min.Y - 1))
	for py := bound.Min.Y; py < bound.Max.Y; py++ {
		go func(cpy int) {
			for px := bound.Min.X; px < bound.Max.X; px++ {
				img.Set(px, cpy, imagePixels.Average(imagePixels.Around(source, px, cpy, searchArea)))
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

func superSamplingNoParallel(source image.Image, searchArea int) image.Image {
	img := image.NewRGBA(source.Bounds())
	bound := img.Bounds()
	for py := bound.Min.Y; py < bound.Max.Y; py++ {
		for px := bound.Min.X; px < bound.Max.X; px++ {
			img.Set(px, py, imagePixels.Average(imagePixels.Around(source, px, py, searchArea)))
		}
	}
	return img
}

func SuperSampling(source image.Image, searchArea int, runInParallel bool) image.Image {
	if runInParallel {
		return superSamplingRunInParallel(source, searchArea)
	}
	return superSamplingNoParallel(source, searchArea)
}
