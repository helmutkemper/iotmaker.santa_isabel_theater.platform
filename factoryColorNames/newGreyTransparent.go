package factoryColorNames

import "image/color"

func NewGreyTransparent() color.RGBA {
	return color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0x00} // rgb(128, 128, 128)
}
