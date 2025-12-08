package common

type DSU struct {
	Rank   []int
	Parent []int
}
type DSUCompatible interface {
	Initialize(n int)
	Find(i int) int
	Union(i, j int)
	GetComponentSizes() map[int]int
}

func (d *DSU) Initialize(n int) {
	d.Rank = make([]int, n)
	parent := make([]int, n)

	for i := range n {
		parent[i] = i
	}

	d.Parent = parent
}

func (d *DSU) Find(i int) int {
	root := i

	for d.Parent[root] != root {
		root = d.Parent[root]
	}
	current := i

	for d.Parent[current] != root {
		next := d.Parent[current]
		d.Parent[current] = root
		current = next
	}

	return root
}

func (d *DSU) Union(i, j int) {
	p1 := d.Find(i)
	p2 := d.Find(j)
	r1 := d.Rank[p1]
	r2 := d.Rank[p2]

	if r1 < r2 {
		d.Parent[p1] = p2
	} else if r1 > r2 {
		d.Parent[p2] = p1
	} else {
		d.Parent[p1] = p2
		d.Rank[p2]++
	}
}

func (d *DSU) GetComponentSizes() map[int]int {
	groupSizes := map[int]int{}

	for _, parent := range d.Parent {
		p := d.Find(parent)
		_, ok := groupSizes[p]

		if !ok {
			groupSizes[p] = 1
		} else {
			groupSizes[p]++
		}
	}

	return groupSizes
}
