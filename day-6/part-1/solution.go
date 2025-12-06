package main

import (
	"aoc-2025/day-6/common"
	"aoc-2025/utils"
	"fmt"
)

func extractValues[T int | byte](line string) []T {
	values := []T{}
	current := ""
	for _, c := range line {
		if c == ' ' {
			var finalValue T

			if len(current) > 0 {
				switch any(finalValue).(type) {
				case int:
					finalValue = T(utils.ToInt(current))
				case byte:
					finalValue = T(current[0])
				}

				values = append(values, finalValue)
			}
			current = ""
		} else {
			current += string(c)
		}
	}

	var finalValue T

	if len(current) > 0 {
		switch any(finalValue).(type) {
		case int:
			finalValue = T(utils.ToInt(current))
		case byte:
			finalValue = T(current[0])
		}

		values = append(values, finalValue)
	}

	return values
}

func extractProblems(input []string) []common.Problem {
	operators := extractValues[byte](input[len(input)-1])
	problemCount := len(operators)
	problems := []common.Problem{}
	values := [][]int{}

	for i := range len(input) - 1 {
		line := input[i]
		intValues := extractValues[int](line)
		values = append(values, intValues)
	}

	for i := range problemCount {
		problemValues := []int{}

		for j := 0; j < len(input)-1; j++ {
			problemValues = append(problemValues, values[j][i])
		}

		problem := common.Problem{
			Values:   problemValues,
			Operator: operators[i],
		}

		problems = append(problems, problem)
	}

	return problems
}

func main() {
	input := utils.ReadInput("../input.txt", false)
	problems := extractProblems(input)
	total := 0

	for _, p := range problems {
		value := common.SolveProblem(p)
		total += value
	}

	fmt.Printf("The grand total of all the problem solutions is %d\n", total)
}
