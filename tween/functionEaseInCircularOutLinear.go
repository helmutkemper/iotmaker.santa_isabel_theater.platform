package tween

import "math"

// en: circular easing in/out - acceleration until halfway, then deceleration
var KEaseInCircularOutLinear = func(interactionCurrent, interactionTotal, startValue, delta float64) float64 {
	interactionCurrentToCalc := interactionCurrent / interactionTotal * 2
	if interactionCurrentToCalc < 1 {
		return -1*delta/2*(math.Sqrt(1-interactionCurrentToCalc*interactionCurrentToCalc)-1) + startValue
	}
	return delta*interactionCurrent/interactionTotal + startValue
}
