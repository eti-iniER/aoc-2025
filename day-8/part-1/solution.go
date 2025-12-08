package main

import (
	"aoc-2025/day-8/common"
	"aoc-2025/utils"
	"fmt"
	"slices"
)

func multiplyLargestNCircuits(dsu *common.DSU, n int) int {
	components := dsu.GetComponentSizes()
	sizes := []int{}

	for _, v := range components {
		sizes = append(sizes, v)
	}

	slices.Sort(sizes)

	ans := 1
	count := 0

	for count < min(len(sizes), n) {
		ans *= sizes[len(sizes)-(count+1)]
		count += 1
	}

	return ans
}

func main() {
	input := utils.ReadInput("../input.txt", false)
	junctionBoxes := common.ExtractJunctionBoxes(input)
	data := []common.DistanceData{}
	n := len(junctionBoxes)

	dsu := common.DSU{
		Rank:   []int{},
		Parent: []int{},
	}

	dsu.Initialize(n)

	for i := range n {
		for j := i + 1; j < n; j++ {
			j1, j2 := junctionBoxes[i], junctionBoxes[j]
			distance := common.FindDistance(j1, j2)
			data = append(data, common.DistanceData{
				Distance:  distance,
				IndexOfJ1: i,
				IndexOfJ2: j,
			})
		}
	}

	slices.SortFunc(data, common.Comp)
	count := 0
	limit := 1000

	for len(data) > 0 {
		last := data[len(data)-1]

		data = data[:len(data)-1]

		i, j := last.IndexOfJ1, last.IndexOfJ2

		dsu.Union(i, j)
		count++

		if count == limit {
			break
		}

	}

	ans := multiplyLargestNCircuits(&dsu, 3)

	fmt.Printf("There are %d circuit(s) in total.\n", ans)
}
