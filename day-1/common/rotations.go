package common

import "strconv"

// Used to extract rotation details from a rotation string
// e.g. "R11" returns 'R' 11
func ExtractRotation(rotation string) (byte, int) {
	var direction = rotation[0]
	clicks, err := strconv.Atoi(rotation[1:])

	if err != nil {
		panic(err)
	}

	return direction, clicks
}

// Perform a rotation, and handle negative results after modular addition
func DoRotation(direction byte, currentRotation, clicks int) int {
	var newRotation int

	if direction == 'L' {
		newRotation = (currentRotation - clicks) % 100
	} else {
		newRotation = (currentRotation + clicks) % 100
	}

	if newRotation < 0 {
		return 100 + newRotation
	}

	return newRotation
}
