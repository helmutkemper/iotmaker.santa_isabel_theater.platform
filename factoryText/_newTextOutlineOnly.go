package factoryText

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/text"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryInk"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewTextOutlineOnly(platform iotmaker_platform_IDraw.IDraw, shadow iotmaker_platform_IDraw.IFilterShadowInterface, gradient iotmaker_platform_IDraw.IFilterGradientInterface, color interface{}, label string, x, y float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) text.Text {
	densityCalc := iDensity
	densityCalc.SetDensityFactor(density)

	densityCalc.Set(x)
	x = densityCalc.Float64()

	densityCalc.Set(y)
	y = densityCalc.Float64()

	ik := ink.Ink{}
	ik = factoryInk.NewInk(ik, 0, color, shadow, gradient, density, iDensity)

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