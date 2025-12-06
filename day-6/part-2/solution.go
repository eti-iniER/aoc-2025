package main

import (
	"aoc-2025/day-6/common"
	"aoc-2025/utils"
	"fmt"
	"strings"
)

func extractProblemFromColumnGroup(input []string, columnStart, columnEnd int) common.Problem {
	operator := strings.Trim(input[len(input)-1][columnStart:columnEnd], " ")[0]
	values := []int{}
	current := ""

	for j := columnStart; j < columnEnd; j++ {
		for i := range len(input) - 1 {
			if input[i][j] != ' ' {
				current += string(input[i][j])
			}
		}

		if len(current) > 0 {
			values = append(values, utils.ToInt(current))
		}
		current = ""
	}

	return common.Problem{
		Values:   values,
		Operator: operator,
	}
}

func extractProblems(input []string) []common.Problem {
	lastColumnStop := 0
	problems := []common.Problem{}

	for j := lastColumnStop; j < len(input[0]); j++ {
		entireColumnIsBlank := true

		for i := range len(input) {
			if input[i][j] != ' ' {
				entireColumnIsBlank = false
			}
		}

		if entireColumnIsBlank {
			problem := extractProblemFromColumnGroup(input, lastColumnStop, j+1)
			problems = append(problems, problem)
			lastColumnStop = j + 1
		}
	}

	lastProblem := extractProblemFromColumnGroup(input, lastColumnStop, len(input[0]))
	problems = append(problems, lastProblem)

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
