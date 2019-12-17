package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: sinusoidal easing in - accelerating from zero velocity
func NewEaseInSine(duration time.Duration, startValue, endValue float64, interactionFunc func(value, percentToComplete float64, arguments []interface{}), doneFunc func(value float64), arguments ...interface{}) *tween.Tween {
	t := &tween.Tween{
		Arguments:   arguments,
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInSine,
		Interaction: interactionFunc,
		Done:        doneFunc,
	}
	t.Start()

	return t
}
