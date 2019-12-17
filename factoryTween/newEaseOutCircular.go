package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: circular easing out - decelerating to zero velocity
func NewEaseOutCircular(duration time.Duration, startValue, endValue float64, interactionFunc func(value, percentToComplete float64, arguments []interface{}), doneFunc func(value float64), arguments ...interface{}) *tween.Tween {
	t := &tween.Tween{
		Arguments:   arguments,
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseOutCircular,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
