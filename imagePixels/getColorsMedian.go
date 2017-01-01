package imagePixels

import (
	"image/color"
	"sort"
)

func GetColorsMedian(colors []color.Color) color.Color {
	var r, g, b, a []int
	mid := len(colors) / 2
	for _, c := range colors {
		cr, cg, cb, ca := c.RGBA()
		r = append(r, int(cr))
		g = append(g, int(cg))
		b = append(b, int(cb))
		a = append(a, int(ca))
	}

	sort.Ints(r)
	sort.Ints(g)
	sort.Ints(b)
	sort.Ints(a)

	return color.RGBA{shiftInt(r[mid]), shiftInt(g[mid]), shiftInt(b[mid]), shiftInt(a[mid])}
}
