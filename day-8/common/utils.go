package common

import "math"

// sort the distance data in reverse order (smallest distance last)
func Comp(a, b DistanceData) int {
	threshold := 0.0000001
	d := math.Abs(a.Distance - b.Distance)

	if d < threshold {
		return 0
	} else if a.Distance > b.Distance {
		return -1
	} else {
		return 1
	}
}
