package imagePixelsTest

import (
	"image/color"
	"testing"

	"github.com/horvatic/vaticwella/src/imagePixels"
)

func TestAverage(test *testing.T) {
	colors := buildColors()

	color := imagePixels.Average(colors)

	r, g, b, _ := color.RGBA()
	if r>>8 != 1 || g>>8 != 2 || b>>8 != 3 {
		test.Errorf("Error colors not averaged r:%d, g:%d, b:%d", r>>8, g>>8, b>>8)
	}
}

func TestAverageAllZero(test *testing.T) {
	colors := buildAllZeroColors()

	color := imagePixels.Average(colors)

	r, g, b, _ := color.RGBA()
	if r != 0 || g != 0 || b != 0 {
		test.Errorf("Error colors not averaged r:%d, g:%d, b:%d", r>>8, g>>8, b>>8)
	}
}

func buildColors() []color.Color {
	var c []color.Color
	c = append(c, color.RGBA{1, 2, 3, 4})
	c = append(c, color.RGBA{1, 2, 3, 4})
	return c

}

func buildAllZeroColors() []color.Color {
	var c []color.Color
	c = append(c, color.RGBA{0, 0, 0, 0})
	c = append(c, color.RGBA{0, 0, 0, 0})
	return c

}
