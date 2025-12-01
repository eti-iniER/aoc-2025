package main

import (
	"aoc-2025/day-1/common"
	"aoc-2025/utils"
	"fmt"
)

func divMod(clicks int) (int, int) {
	return clicks / 100, clicks % 100
}

func main() {
	input := utils.ReadInput("../input.txt", false)
	currentRotation := 50
	var password int

	for _, rotation := range input {
		direction, clicks := common.ExtractRotation(rotation)
		loops, effectiveRotation := divMod(clicks)
		previousRotation := currentRotation
		currentRotation = common.DoRotation(direction, currentRotation, effectiveRotation)

		// All the rotations that spin around the dial and are ignored
		password += loops

		if direction == 'L' {
			if effectiveRotation >= previousRotation && previousRotation != 0 {
				password++
			}
		} else {
			if effectiveRotation+previousRotation >= 100 {
				password++
			}
		}

	}

	fmt.Printf("The password is: %d", password)
}
