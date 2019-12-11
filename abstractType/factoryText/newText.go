package factoryText

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
	"github.com/helmutkemper/iotmaker.platform/abstractType/text"
)

func NewText(
	platform iotmaker_platform_IDraw.IDraw,
	shadow iotmaker_platform_IDraw.IFilterShadowInterface,
	gradient iotmaker_platform_IDraw.IFilterGradientInterface,
	label string,
	x,
	y int,
	density interface{}, iDensity iotmaker_platform_coordinate.IDensity) text.Text {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Int()

	densityCalc.Set(y)
	y = densityCalc.Int()

	lineWidth := 1 //fixme!
	ik := genericTypes.Ink{}
	ik = genericTypes.NewInc(ik, lineWidth, shadow, gradient, density, iDensity)

	tx := text.Text{
		Platform: platform,
		Ink:      ik,
		Label:    label,
		Stroke:   true,
		X:        x,
		Y:        y,
	}

	tx.ConfigShadowPlatformAndFilter()
	tx.ConfigGradientPlatformAndFilter()
	tx.Create()

	return tx
}
