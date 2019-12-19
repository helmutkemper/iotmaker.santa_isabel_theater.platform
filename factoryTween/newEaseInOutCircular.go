package factoryTween

import (
	"github.com/helmutkemper/iotmaker.platform/tween"
	"time"
)

// en: circular easing in/out - acceleration until halfway, then deceleration
func NewEaseInOutCircular(duration time.Duration, startValue, endValue float64, interactionFunc func(value, percentToComplete float64, arguments []interface{}), onStartFunc func(value float64), onEndFunc func(value float64), arguments ...interface{}) *tween.Tween {
	t := &tween.Tween{
		OnStart:     onStartFunc,
		OnEnd:       onEndFunc,
		Arguments:   arguments,
		Duration:    duration,
		StartValue:  startValue,
		EndValue:    endValue,
		Func:        tween.KEaseInOutCircular,
		Interaction: interactionFunc,
	}
	t.Start()

	return t
}
