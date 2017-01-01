package imagePixels

func shiftInt(color int) uint8 {
	return uint8(color >> 8)
}

func shiftUint32(color uint32) uint8 {
	return uint8(color >> 8)
}
