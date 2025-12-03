package common

import "strconv"

func MaxFromLine(line string, desiredLength int) int {
	startIndex := 0
	var value string

	for len(value) < desiredLength {
		endIndex := len(line) - (desiredLength - len(value))
		maxCharIndex := startIndex

		for i := startIndex; i <= endIndex; i++ {
			if line[i] > line[maxCharIndex] {
				maxCharIndex = i
			}
		}

		value += string(line[maxCharIndex])
		startIndex = maxCharIndex + 1
	}

	intValue, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return intValue
}
