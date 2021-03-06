// +build js

package main

import (
  iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
  coordinateManager "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserDocument"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserStage"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/fontFamily"
  webBrowserMouse "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/mouse"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/colornames"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/selectBox"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryColor"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryFont"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryGradient"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryPoint"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryShadow"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryText"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
  "time"
)

var (
	density                                   = 3.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	bx2                                       = &draw.BasicBox{}
	stage                                     = canvas.Stage{}
	gradientFilter iotmakerPlatformIDraw.IFilterGradientInterface
	html           iotmakerPlatformIDraw.IHtml
)

func main() {

	//todo: canvasDrawImage()

	done := make(chan struct{}, 0)

	html = &html.Html{}
	browserDocument := factoryBrowserDocument.NewDocument()

	var colorShadow = colornames.BlackHalfTransparent
	var blur int = 10
	var offsetX int = 2
	var offsetY int = 2
	var shadowFilter = factoryShadow.NewShadowFilter(colorShadow, blur, offsetX, offsetY, density, densityManager)

	//mouse.AddFunctionPointer(bx1.GetAlphaChannel)

	stage = factoryBrowserStage.NewStage(
		browserDocument,
		"stage",
		300,
		300,
		density,
		densityManager,
	)

	colorWhite := factoryColor.NewColorPosition(colornames.Red, 0.5)
	colorBlack := factoryColor.NewColorPosition(colornames.Black, 1)
	colorList := factoryColor.NewColorList(colorBlack, colorWhite)

	coordinateP0 := factoryPoint.NewPoint(0, 0, density, densityManager)
	coordinateP1 := factoryPoint.NewPoint(120, 150, density, densityManager)
	gradientFilter = factoryGradient.NewGradientLinearToFillAndStroke(coordinateP0, coordinateP1, colorList)

	fontText := factoryFont.NewFont(15, "px", fontFamily.KVerdana, density, densityManager)
	factoryFont.SetGlobal(
		&stage.Canvas,
		fontText,
	)

	factoryText.NewText(
		&stage.Canvas,
		shadowFilter,
		gradientFilter,
		nil,
		"Olá Mundo!",
		25,
		20,
		density,
		densityManager,
	)

	fontText = factoryFont.NewFont(15, "px", fontFamily.KArial, density, densityManager)
	factoryText.NewTextWithFont(
		&stage.Canvas,
		nil,
		nil,
		colornames.Cadetblue,
		fontText,
		"Olá Mundo!",
		25,
		40,
		density,
		densityManager,
	)

	factoryGradient.ResetStylesGlobal(&stage.Canvas)
	factoryDraw.NewBasicBox(
		&stage.Canvas,
		&stage.ScratchPad,
		nil,
		"bbox_1",
		20,
		50,
		100,
		100,
		5,
		5,
		nil,
		nil,
		density,
		densityManager,
	)

	bx2 = factoryDraw.NewBasicBox(
		&stage.Canvas,
		&stage.ScratchPad,
		nil,
		"bbox_2",
		20+50,
		50+50,
		100,
		100,
		10,
		8,
		shadowFilter,
		gradientFilter,
		density,
		densityManager,
	)

	selectBox.NewResizeBoxFromBasicBox(bx2, -3, -3, 6, 6, 1, density, densityManager)

	factoryGradient.ResetStylesGlobal(&stage.Canvas)
	imgHtml := factoryImage.NewHtmlImage(
		html,
		browserDocument.SelfDocument,
		map[string]interface{}{
			"id":  "player",
			"src": "./player_big.png",
		},
		true,
		true,
	)

	i := factoryImage.NewMultipleSpritesImage(
		&stage.Canvas,
		imgHtml,
		48,
		60,
		0,
		7,
		80*time.Millisecond,
		45,
		70,
		48,
		60,
		density,
		densityManager,
	)
	i.Crete()

	//fmt.Printf("over: %v\n", bx.GetAlphaChannel(0, 100))

	/*abstractType.NewLineTo(
		&stage.Canvas,
		"line",
		1.0,
		0,
		0,
		300,
		300,
		1,
	)*/

	browserDocument.SetMouseMoveListener(webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	mouse.SetPlatform(mouse.KPlatformWebBrowser)
	mouse.AddFunctionPointer("bBox2", bx2.GetCollisionBySimpleBox, bateu)

	<-done
}

var lastCursor bool

func bateu(x, y int, collision bool) {
	if collision == false {
		webBrowserMouse.SetCursor(stage.SelfElement, mouse.KCursorAuto)
	} else {
		webBrowserMouse.SetCursor(stage.SelfElement, mouse.KCursorColResize)
	}

	lastCursor = collision

	//fmt.Printf("bateu: %v\n", collision)
}
