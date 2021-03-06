package factoryGradient

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/gradient"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/point"
)

// en: Make a new filter linear gradient with fill for use with the canvas elements.
//
// Linear gradient creates a gradient along one imaginary line connecting two given coordinates, starting at (x0, y0)
// point and ending at (x1, y1) point and fill the image with colors in list
//
// pt_bt: Monta um novo filtro linear gradient para ser usado com os elementos do canvas.
//
// O gradiente linear cria um gradiente ao longo de uma linha imaginária com duas coordenadas, começando no ponto 0
// (x0, y0) e terminando no ponto 1 (x1, y1), pintando a linha com as cores da lista.
//    coordinateP0: Coordinate of the start point of the gradient. Please, use a NewPoint() function to set a point.
//    coordinateP1: Coordinate of the end point of the gradient. Please, use a NewPoint() function to set a point.
//    colorList:    Color position list. Please, use a NewColorPosition() and NewColorList() functions to set a color
//                  list. Example: NewColorList(NewColorPosition(), NewColorPosition(), ...)
func NewGradientLinearToFillBasicElement(colorList []gradient.ColorStop) iotmaker_platform_IDraw.IFilterGradientInterface {
	return &gradient.Gradient{
		Type:         gradient.KLinearFill,
		CoordinateP0: point.Point{},
		CoordinateP1: point.Point{},
		ColorList:    colorList,
	}
}
