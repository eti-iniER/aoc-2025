package common

func SolveProblem(p Problem) int {
	var value int
	if p.Operator == '*' {
		value = 1
	} else {
		value = 0
	}

	for _, v := range p.Values {
		if p.Operator == '*' {
			value *= v
		} else {
			value += v
		}
	}

	return value
}
