package factoryColorNames

import "image/color"

func NewLightgreyTransparent() color.RGBA {
	return color.RGBA{R: 0xd3, G: 0xd3, B: 0xd3, A: 0x00} // rgb(211, 211, 211)
}
