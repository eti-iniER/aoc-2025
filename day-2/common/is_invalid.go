package common

import (
	"math"
	"strconv"
)

func IsInvalid(value int, allowMoreThanTwo bool) bool {
	valueAsString := strconv.Itoa(value)
	var prefix string
	var anyIsInvalid bool

	for _, v := range valueAsString {
		prefix += string(v)

		maxMultiples := 2
		if allowMoreThanTwo {
			maxMultiples = math.MaxInt64
		}

		multiples := len(valueAsString) / len(prefix)

		if len(valueAsString)%len(prefix) == 0 && multiples > 1 && multiples <= maxMultiples {
			var invalidCount int

			for offset := len(prefix); offset < len(valueAsString); offset += len(prefix) {
				invalidCount++

				for i := 0; i < len(prefix); i++ {
					if prefix[i] != valueAsString[offset+i] {
						invalidCount--
						break
					}
				}
			}

			if invalidCount+1 == multiples {
				anyIsInvalid = true
			}
		}
	}

	return anyIsInvalid
}
