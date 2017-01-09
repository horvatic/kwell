package imagePixels

func shift(color uint32) uint8 {
	return uint8(color >> 8)
}

type Colors []uint32

func (c Colors) Len() int {
	return len(c)
}

func (c Colors) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c Colors) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
