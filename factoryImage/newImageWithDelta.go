package factoryImage

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/image"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
)

func NewImageWithDelta(id string, stage iStage.IStage, platform, scratchPad iotmaker_platform_IDraw.IDraw, img interface{}, x, y, width, height, xDelta, yDelta int, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *image.Image {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.SetInt(x)
	x = densityCalc.Int()

	densityCalc.SetInt(y)
	y = densityCalc.Int()

	densityCalc.SetInt(width)
	width = densityCalc.Int()

	densityCalc.SetInt(height)
	height = densityCalc.Int()

	densityCalc.SetInt(xDelta)
	xDelta = densityCalc.Int()

	densityCalc.SetInt(yDelta)
	yDelta = densityCalc.Int()

	ret := &image.Image{
		Sprite: basic.Sprite{
			Id:         id,
			Stage:      stage,
			Platform:   platform,
			ScratchPad: scratchPad,
			Dimensions: dimensions.Dimensions{
				X:      x,
				Y:      y,
				Width:  width,
				Height: height,
			},
			OutBoxDimensions: dimensions.Dimensions{
				X:      x,
				Y:      y,
				Width:  width,
				Height: height,
			},

			MovieDeltaX: int(xDelta),
			MovieDeltaY: int(yDelta),
		},
		Img: img,
	}
	ret.Crete()

	return ret
}
