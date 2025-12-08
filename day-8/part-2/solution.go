package main

import (
	"aoc-2025/day-8/common"
	"aoc-2025/utils"
	"fmt"
	"slices"
)

func getLargestComponentSize(dsu *common.DSU) int {
	components := dsu.GetComponentSizes()
	sizes := []int{}

	for _, v := range components {
		sizes = append(sizes, v)
	}

	slices.Sort(sizes)

	return sizes[len(sizes)-1]
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
	ans := 0

	for len(data) > 0 {
		last := data[len(data)-1]

		data = data[:len(data)-1]

		i, j := last.IndexOfJ1, last.IndexOfJ2

		dsu.Union(i, j)
		largestComponentSize := getLargestComponentSize(&dsu)

		if largestComponentSize == n {
			ans = junctionBoxes[i].X * junctionBoxes[j].X
			break
		}

	}

	fmt.Printf("The product of the X coordinates of the last pair that needs to be connected is %d\n", ans)
}
