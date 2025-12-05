package utils

import "strconv"

func ToInt(value string) int {
	intValue, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return intValue
}
