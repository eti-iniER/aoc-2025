package common

import "math"

func FindDistance(j1, j2 JunctionBox) float64 {
	xDist := math.Pow(float64(j1.X-j2.X), 2)
	yDist := math.Pow(float64(j1.Y-j2.Y), 2)
	zDist := math.Pow(float64(j1.Z-j2.Z), 2)
	return math.Sqrt(xDist + yDist + zDist)
}
