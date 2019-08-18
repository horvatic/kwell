package imagePixels

func shift(color uint32) uint8 {
	return uint8(color >> 8)
}
