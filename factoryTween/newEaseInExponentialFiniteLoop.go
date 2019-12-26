package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: exponential easing in - accelerating from zero velocity
func NewEaseInExponentialFiniteLoop(duration time.Duration, loop int, startValue, endValue float64, interactionFunc func(value, percentToComplete float64, arguments []interface{}), onStartFunc func(value float64), onEndFunc func(value float64), arguments ...interface{}) *tween.Tween {
	t := &tween.Tween{
		OnStart:     onStartFunc,
		OnEnd:       onEndFunc,
		Arguments:   arguments,
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInExponential,
		Interaction: interactionFunc,
		Repeat:      loop,
	}
	t.Start()

	return t
}
