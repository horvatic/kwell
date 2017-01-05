package test

import (
	"github.com/horvatic/vaticwella/imagePixels"
	"image"
	"image/color"
	"strconv"
	"testing"
)

func TestAround(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	imageSetUp(img)

	colors := imagePixels.Around(img, 5, 5)

	testColors(colors, t)
}

func testColors(colors []color.Color, t *testing.T) {
	expColor := counterUInt32()
	for _, c := range colors {
		cr, cg, cb, ca := c.RGBA()
		er := expColor()
		eg := expColor()
		eb := expColor()
		ea := expColor()

		checkColor(t, er, cr>>8)
		checkColor(t, eg, cg>>8)
		checkColor(t, eb, cb>>8)
		checkColor(t, ea, ca>>8)

	}
}

func imageSetUp(img *image.RGBA) {
	setColor := counterUInt8()
	img.Set(4, 4, color.RGBA{setColor(), setColor(), setColor(), setColor()})
	img.Set(4, 5, color.RGBA{setColor(), setColor(), setColor(), setColor()})
	img.Set(4, 6, color.RGBA{setColor(), setColor(), setColor(), setColor()})
	img.Set(5, 4, color.RGBA{setColor(), setColor(), setColor(), setColor()})
	img.Set(5, 5, color.RGBA{setColor(), setColor(), setColor(), setColor()})
	img.Set(5, 6, color.RGBA{setColor(), setColor(), setColor(), setColor()})
	img.Set(6, 4, color.RGBA{setColor(), setColor(), setColor(), setColor()})
	img.Set(6, 5, color.RGBA{setColor(), setColor(), setColor(), setColor()})
	img.Set(6, 6, color.RGBA{setColor(), setColor(), setColor(), setColor()})
}

func checkColor(t *testing.T, expectColor, currentColor uint32) {

	if expectColor != currentColor {
		t.Error("Expect color to be " + strconv.Itoa(int(expectColor)) + " color was " + strconv.Itoa(int(currentColor)))
	}
}

func counterUInt8() func() uint8 {
	var cnt uint8
	return func() uint8 {
		cnt++
		return cnt
	}
}

func counterUInt32() func() uint32 {
	var cnt uint32
	return func() uint32 {
		cnt++
		return cnt
	}
}
