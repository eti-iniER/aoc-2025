package main

import (
	"aoc-2025/day-1/common"
	"aoc-2025/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("../input.txt", false)
	currentRotation := 50
	var password int

	for _, rotation := range input {
		direction, clicks := common.ExtractRotation(rotation)
		currentRotation = common.DoRotation(direction, currentRotation, clicks)

		if currentRotation == 0 {
			password++
		}
	}

	fmt.Printf("The password is: %d", password)
}
