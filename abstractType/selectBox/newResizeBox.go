package selectBox

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	mouseWebBrowser "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/draw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryDimensions"
	"image/color"
)

func NewResizeBoxFromBasicBox(basicBox *draw.BasicBox, offsetX, offsetY, width, height, border float64, density interface{}, iDensity iotmaker_platform_coordinate.IDensity) *ResizeBoxes {

	dm := dimensions.Dimensions{}
	dm = factoryDimensions.NewDimensions(dm, offsetX, offsetY, width, height, border, density, iDensity)

	rb := &ResizeBoxes{
		CornerFillColor:        color.RGBA{R: 0x00, G: 0x90, B: 0x90, A: 0xFF},
		CornerLineWidth:        0,
		ContourColor:           color.RGBA{R: 0xFF, G: 0x40, B: 0x40, A: 0xFF},
		ContourLineWidth:       3,
		MouseCornerA:           mouseWebBrowser.KCursorNWSeResize,
		MouseCornerB:           mouseWebBrowser.KCursorNSResize,
		MouseCornerC:           mouseWebBrowser.KCursorNESwResize,
		MouseCornerD:           mouseWebBrowser.KCursorEWResize,
		MouseCornerE:           mouseWebBrowser.KCursorNWSeResize,
		MouseCornerF:           mouseWebBrowser.KCursorNSResize,
		MouseCornerG:           mouseWebBrowser.KCursorNESwResize,
		MouseCornerH:           mouseWebBrowser.KCursorEWResize,
		MouseGeneric:           mouseWebBrowser.KCursorAuto,
		Platform:               basicBox.Platform,
		ScratchPad:             basicBox.ScratchPad,
		Dimensions:             dm,
		FatherOutBoxDimensions: basicBox.OutBoxDimensions,
	}
	rb.Create()

	return rb
}
