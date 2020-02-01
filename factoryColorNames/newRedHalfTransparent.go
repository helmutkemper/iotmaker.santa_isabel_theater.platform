package factoryColorNames

import "image/color"

func NewRedHalfTransparent() color.RGBA {
	return color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0x80} // rgb(255, 0, 0)
}