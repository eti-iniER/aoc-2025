package common

import "aoc-2025/utils"

func FindArea(t1, t2 []int) int {
	return (utils.AbsInt(t1[0]-t2[0]) + 1) * (utils.AbsInt(t1[1]-t2[1]) + 1)
}
