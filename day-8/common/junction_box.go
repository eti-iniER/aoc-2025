package common

import "fmt"

type JunctionBox struct {
	X, Y, Z int
}

func (j JunctionBox) String() string {
	return fmt.Sprintf("Junction Box (%d, %d, %d)", j.X, j.Y, j.Z)
}
