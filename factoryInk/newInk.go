package factoryInk

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
)

func NewInk(lineWidth int, fillColor interface{}) ink.Ink {
	var inkObj = ink.Ink{}

	densityCalc := global.Global.DensityManager
	densityCalc.SetInt(lineWidth)
	densityCalc.SetDensityFactor(global.Global.Density)

	inkObj.Set(densityCalc.Int(), fillColor, nil, nil)
	return inkObj
}
