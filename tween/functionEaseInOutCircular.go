package tween

import "math"

// en: circular easing in/out - acceleration until halfway, then deceleration
var KEaseInOutCircular = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrent = interactionCurrent / interactionTotal * 2
	if interactionCurrent < 1 {
		return -1*delta/2*(math.Sqrt(1-interactionCurrent*interactionCurrent)-1) + startValue
	}
	interactionCurrent -= 2
	return delta/2*(math.Sqrt(1-interactionCurrent*interactionCurrent)+1) + startValue
}
