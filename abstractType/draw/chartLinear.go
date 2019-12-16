package draw

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
)

type ChartLinear struct {
	Platform iotmaker_platform_IDraw.IDraw

	X      int
	Y      int
	Width  int
	Height int

	Density  interface{}
	IDensity iotmaker_platform_coordinate.IDensity
}

func (el *ChartLinear) Create() {

	// axe x
	el.lineTo(el.X, el.Y, el.X, el.Y+el.Height, 1)

	// axe y
	el.lineTo(el.X, el.Y+el.Height, el.X+el.Width, el.Y+el.Height, 1)

	for i := el.X; i <= el.X+el.Width; i += 10 {
		el.lineTo(i, el.Y, i, el.Y+el.Height+4, 0.1)
	}

	for i := el.Y; i <= el.Y+el.Height; i += 10 {
		el.lineTo(el.X-4, i, el.X+el.Width, i, 0.1)
	}
}

func (el *ChartLinear) lineTo(x1, y1, x2, y2 int, lineWidth interface{}) {
	el.Platform.BeginPath()
	el.Platform.SetLineWidth(lineWidth)
	el.Platform.MoveTo(x1, y1)
	el.Platform.LineTo(x2, y2)
	el.Platform.ClosePath(x1, y1)
	el.Platform.Stroke()
}

func (el *ChartLinear) Begin(x, y int) {
	el.Platform.BeginPath()
	el.Platform.SetLineWidth(1)
	el.Platform.MoveTo(x, y)
}

func (el *ChartLinear) Point(x, y int) {
	el.Platform.LineTo(x, y)
	el.Platform.Stroke()
}

func (el *ChartLinear) End() {
	el.Platform.Stroke()
}
