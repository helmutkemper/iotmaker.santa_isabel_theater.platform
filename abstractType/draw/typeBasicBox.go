package draw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/genericTypes"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/independentDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
	iotmaker_threadsafe "github.com/helmutkemper/iotmaker.threadsafe"
	"image/color"
	"reflect"
)

type BasicBox struct {
	Platform   iotmaker_platform_IDraw.IDraw
	ScratchPad iotmaker_platform_IDraw.IDraw
	Id         string
	Dimensions dimensions.Dimensions

	Ink ink.Ink

	imageDataMethod           genericTypes.ImageDataCaptureMethod
	imageDataComplete         interface{}
	imageDataAlphaChannel     map[int]map[int]uint8
	imageDataBooleanCollision map[int]map[int]bool

	prepareShadowFilterFunctionPointer   func(iotmaker_platform_IDraw.ICanvasShadow)
	prepareGradientFilterFunctionPointer func(iotmaker_platform_IDraw.ICanvasGradient)

	// see clearRectangle() and getCompleteImageData() - start
	OutBoxDimensions dimensions.Dimensions
	// see clearRectangle() and getCompleteImageData() - end

	alphaChannelSensibility  uint8
	enableDataImageCalculate bool
}

func (el *BasicBox) prepareShadowFilter() {
	if el.prepareShadowFilterFunctionPointer != nil {
		el.prepareShadowFilterFunctionPointer(el.Platform)
	}
}

func (el *BasicBox) prepareGradientFilter(platform iotmaker_platform_IDraw.IDraw) {
	if reflect.DeepEqual(el.Ink.Color, color.RGBA{}) != true {
		platform.SetFillStyle(el.Ink.Color)
	}

	if el.prepareGradientFilterFunctionPointer != nil {
		el.prepareGradientFilterFunctionPointer(platform)
	}
}

func (el *BasicBox) ConfigShadowPlatformAndFilter() {
	if el.Ink.Shadow == nil {
		return
	}

	el.setShadowFilter(el.Ink.Shadow.PrepareFilter)
}

func (el *BasicBox) ConfigGradientPlatformAndFilter() {
	if el.Ink.Gradient == nil {
		return
	}

	el.setGradientAndMountColorListFilter(el.Ink.Gradient.PrepareFilter)
}

func (el *BasicBox) setGradientAndMountColorListFilter(f func(iotmaker_platform_IDraw.ICanvasGradient)) {
	el.prepareGradientFilterFunctionPointer = f
}

func (el *BasicBox) setShadowFilter(f func(iotmaker_platform_IDraw.ICanvasShadow)) {
	el.prepareShadowFilterFunctionPointer = f
}

func (el *BasicBox) calculateCoordinates() {
	//  draw_1:
	//              border        border
	//             x1  x2         x3 x4
	//           l     a          b     c
	//        y1    +--|----------|--+    y1
	//  border      |                |      border
	//        y2 k ---              --- d y2
	//              |                |
	//              |                |
	//        y3 j ---              --- e y3
	//  border      |                |      border
	//        y4    +--|----------|--+    y4
	//           i     h          g     f
	//             x1  x2         x3 x4
	//              border        border

	// correct the line width size
	el.Dimensions.X -= el.Ink.LineWidth / 2
	el.Dimensions.Y -= el.Ink.LineWidth / 2
	el.Dimensions.Width -= el.Ink.LineWidth
	el.Dimensions.Height -= el.Ink.LineWidth

	// calculate outline from the box
	x := el.Dimensions.X - el.Ink.LineWidth/2
	y := el.Dimensions.Y - el.Ink.LineWidth/2
	width := el.Dimensions.Width + el.Ink.LineWidth
	height := el.Dimensions.Height + el.Ink.LineWidth

	dm := dimensions.NewContainerWidthXY("calculate", x, y, width, height)
	el.OutBoxDimensions = *dm
	//el.OutBoxDimensions.Set(x, y, width, height, 0)
}

func (el *BasicBox) getCompleteImageData(platform iotmaker_platform_IDraw.IDraw) {
	//el.imageDataComplete = platform.GetImageData(int(el.OutBoxDimensions.X), int(el.OutBoxDimensions.Y), int(el.OutBoxDimensions.Width), int(el.OutBoxDimensions.Height), true)
}

func (el *BasicBox) getImageDataAlphaChannelOnly(platform iotmaker_platform_IDraw.IDraw) {
	el.imageDataAlphaChannel = platform.GetImageDataAlphaChannelOnly(int(el.OutBoxDimensions.X), int(el.OutBoxDimensions.Y), int(el.OutBoxDimensions.Width), int(el.OutBoxDimensions.Height))
}
func (el *BasicBox) getImageDataCollisionByAlphaChannelValue(platform iotmaker_platform_IDraw.IDraw) {
	el.imageDataBooleanCollision = platform.GetImageDataCollisionByAlphaChannelValue(int(el.OutBoxDimensions.X), int(el.OutBoxDimensions.Y), int(el.OutBoxDimensions.Width), int(el.OutBoxDimensions.Height), el.alphaChannelSensibility)
}

func (el *BasicBox) clearRectangle(platform iotmaker_platform_IDraw.IDraw) {
	platform.ClearRect(el.OutBoxDimensions.X, el.OutBoxDimensions.Y, el.OutBoxDimensions.Width, el.OutBoxDimensions.Height)
}

func (el *BasicBox) drawInvisible(platform iotmaker_platform_IDraw.IDraw) {
	platform.SetLineWidth(el.Ink.LineWidth)
	el.prepareGradientFilter(el.Platform)
	independentDraw.DrawBoxWithRoundedCornersIntoThePath(
		platform,
		float64(el.Dimensions.X),
		float64(el.Dimensions.Y),
		float64(el.Dimensions.Width),
		float64(el.Dimensions.Height),
		float64(el.Dimensions.Border),
	)
}

func (el *BasicBox) drawVisible() {
	el.drawInvisible(el.Platform)
	el.prepareGradientFilter(el.Platform)
	el.prepareShadowFilter()
	el.Platform.Stroke()
}

func (el *BasicBox) Create() {
	el.enableDataImageCalculate = true
	el.imageDataMethod = genericTypes.KImageDataCaptureMethodCompleteData

	el.calculateCoordinates()

	if el.enableDataImageCalculate == true {
		el.CalculateImageData()
	}

	el.drawVisible()

	if el.alphaChannelSensibility == 0 {
		el.alphaChannelSensibility = 25 // 10% de 255
	}

	el.Platform.ResetStrokeStyle()
	el.Platform.ResetFillStyle()
	el.Platform.ResetShadow()
	el.Platform.ResetLineWidth()
}

// CalculateImageData
// see SetEnableDataImageCalculate()
// see SetAlphaChannelSensibility()
// see SetImageDataCalculateMethod()
func (el *BasicBox) CalculateImageData() {
	switch el.imageDataMethod {
	case genericTypes.KImageDataCaptureMethodCompleteData:
		iotmaker_threadsafe.ScratchPad(
			el.ScratchPad,
			el.prepareGradientFilter,
			el.drawInvisible,
			el.getCompleteImageData,
			el.clearRectangle,
		)
	case genericTypes.KImageDataCaptureMethodAlphaChannelOnly:
		iotmaker_threadsafe.ScratchPad(
			el.ScratchPad,
			el.prepareGradientFilter,
			el.drawInvisible,
			el.getImageDataAlphaChannelOnly,
			el.clearRectangle,
		)
	case genericTypes.KImageDataCaptureMethodBooleanSensibility:
		iotmaker_threadsafe.ScratchPad(
			el.ScratchPad,
			el.prepareGradientFilter,
			el.drawInvisible,
			el.getImageDataCollisionByAlphaChannelValue,
			el.clearRectangle,
		)
	}

	el.ScratchPad.ResetStrokeStyle()
	el.ScratchPad.ResetFillStyle()
	el.ScratchPad.ResetShadow()
	el.ScratchPad.ClearRect(el.OutBoxDimensions.X, el.OutBoxDimensions.Y, el.OutBoxDimensions.Width, el.OutBoxDimensions.Height)
}

// SetImageDataCalculateMethod
// pt_br: Define o método usado para calcular os dados da imagem usado para detectar colisão
//     KImageDataCaptureMethodBooleanSensibility: Arquiva um mapa de booleanos no formato mapa[x][y]bool
//     KImageDataCaptureMethodCompleteData: Arquiva um mapa de RGBA no formato mapa[x][y]RGBA
//     KImageDataCaptureMethodAlphaChannelOnly: Arquiva um mapa de uint8 no formato mapa[x][y]uint8
//     Veja também SetAlphaChannelSensibility() e SetEnableDataImageCalculate()
func (el *BasicBox) SetImageDataCalculateMethod(value genericTypes.ImageDataCaptureMethod) {
	el.imageDataMethod = value
}

func (el *BasicBox) SetAlphaChannelSensibility(value uint8) {
	el.alphaChannelSensibility = value
}

func (el *BasicBox) SetEnableDataImageCalculate(value bool) {
	el.enableDataImageCalculate = value
}

func (el *BasicBox) GetCollisionByAlphaChannel(x, y int) bool {
	return el.imageDataBooleanCollision[x][y]
}

func (el *BasicBox) GetCollisionBox(xEvent, yEvent float64) bool {
	return float64(el.OutBoxDimensions.X) <= xEvent && float64(el.OutBoxDimensions.X+el.OutBoxDimensions.Width) >= xEvent &&
		float64(el.OutBoxDimensions.Y) <= yEvent && float64(el.OutBoxDimensions.Y+el.OutBoxDimensions.Height) >= yEvent
}

func (el *BasicBox) GetMouseOverFunctionsList() map[string][]mouse.PointerCollisionFunction {
	ret := make(map[string][]mouse.PointerCollisionFunction)
	ret[el.Id] = []mouse.PointerCollisionFunction{
		el.GetCollisionBox,
	}

	return ret
}

func (el *BasicBox) GetOutBoxDimensions() dimensions.Dimensions {
	return el.OutBoxDimensions
}

func (el *BasicBox) GetPixelAlphaChannel(x, y float64) uint8 {
	if el.enableDataImageCalculate == false {
		return 0
	}

	if x < float64(el.OutBoxDimensions.X) || x > float64(el.OutBoxDimensions.X+el.OutBoxDimensions.Width) {
		return 0
	}

	if y < float64(el.OutBoxDimensions.Y) || y > float64(el.OutBoxDimensions.Y+el.OutBoxDimensions.Height) {
		return 0
	}

	return el.Platform.GetImageDataAlphaChannelByCoordinate(
		el.imageDataComplete,
		int(x)-el.OutBoxDimensions.X,
		int(y)-el.OutBoxDimensions.Y,
		int(el.OutBoxDimensions.Width),
	)
}

func (el *BasicBox) GetPixelColor(x, y float64) color.RGBA {
	if el.enableDataImageCalculate == false {
		return color.RGBA{}
	}

	if x < float64(el.OutBoxDimensions.X) || x > float64(el.OutBoxDimensions.X+el.OutBoxDimensions.Width) {
		return color.RGBA{}
	}

	if y < float64(el.OutBoxDimensions.Y) || y > float64(el.OutBoxDimensions.Y+el.OutBoxDimensions.Height) {
		return color.RGBA{}
	}

	return el.Platform.GetImageDataPixelByCoordinate(
		el.imageDataComplete,
		int(x)-el.OutBoxDimensions.X,
		int(y)-el.OutBoxDimensions.Y,
		int(el.OutBoxDimensions.Width),
	)
}

func (el *BasicBox) GetPlatform() iotmaker_platform_IDraw.IDraw {
	return el.Platform
}

func (el *BasicBox) GetScratchPad() iotmaker_platform_IDraw.IDraw {
	return el.ScratchPad
}

func (el *BasicBox) GetId() string {
	return el.Id
}

func (el *BasicBox) GetDimensions() dimensions.Dimensions {
	return el.Dimensions
}

func (el *BasicBox) GetInk() ink.Ink {
	return el.Ink
}
