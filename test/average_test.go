package test

import (
	"github.com/horvatic/vaticwella/imagePixels"
	"image/color"
	"testing"
)

func TestAverage(test *testing.T) {
	colors := buildColors()

	color := imagePixels.Average(colors)

	r, g, b, _ := color.RGBA()
	if r>>8 != 1 && g>>8 != 2 && b>>8 != 3 {
		test.Errorf("Error colors not averaged r:%d, g:%d, b:%d", r>>8, g>>8, b>>8)
	}
}

func buildColors() []color.Color {
	var c []color.Color
	c = append(c, color.RGBA{1, 2, 3, 4})
	c = append(c, color.RGBA{1, 2, 3, 4})
	return c

}
