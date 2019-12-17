package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: quartic easing in/out - acceleration until halfway, then deceleration
func NewEaseInOutQuartic(duration time.Duration, startValue, endValue float64, interactionFunc func(value, percentToComplete float64, arguments []interface{}), doneFunc func(value float64), arguments ...interface{}) *tween.Tween {
	t := &tween.Tween{
		Arguments:   arguments,
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInOutQuartic,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
