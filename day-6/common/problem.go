package common

import "fmt"

type Problem struct {
	Values   []int
	Operator byte
}

func (p Problem) String() string {
	return fmt.Sprintf("Problem: \n   Values: %v\n   Operator: %c", p.Values, p.Operator)
}
