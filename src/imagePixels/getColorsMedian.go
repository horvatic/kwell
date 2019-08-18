package imagePixels

import (
	"image/color"
	"sort"
)

func GetColorsMedian(colors []color.Color) color.Color {
	var r, g, b, a Colors
	mid := len(colors) / 2
	for _, c := range colors {
		cr, cg, cb, ca := c.RGBA()
		r = append(r, cr)
		g = append(g, cg)
		b = append(b, cb)
		a = append(a, ca)
	}

	sort.Sort(r)
	sort.Sort(g)
	sort.Sort(b)
	sort.Sort(a)

	return color.RGBA{shift(r[mid]), shift(g[mid]), shift(b[mid]), shift(a[mid])}
}
