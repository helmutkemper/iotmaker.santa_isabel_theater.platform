package factoryColorNames

import "image/color"

func NewSlategrayTransparent() color.RGBA {
	return color.RGBA{R: 0x70, G: 0x80, B: 0x90, A: 0x00} // rgb(112, 128, 144)
}
