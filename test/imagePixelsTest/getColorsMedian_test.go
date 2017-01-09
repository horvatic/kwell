package imagePixelsTest

import (
	"github.com/horvatic/vaticwella/imagePixels"
	"image/color"
	"testing"
)

func TestGetColorsMedian(test *testing.T) {
	colors := buildMixedColors()

	color := imagePixels.GetColorsMedian(colors)

	r, g, b, _ := color.RGBA()
	if r>>8 != 2 || g>>8 != 5 || b>>8 != 6 {
		test.Errorf("Error did not find median color r:%d, g:%d, b:%d", r>>8, g>>8, b>>8)
	}
}

func buildMixedColors() []color.Color {
	var c []color.Color
	c = append(c, color.RGBA{1, 5, 7, 10})
	c = append(c, color.RGBA{2, 6, 2, 12})
	c = append(c, color.RGBA{3, 4, 6, 1})
	return c

}
