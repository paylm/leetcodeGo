package main

type DisjSet []int

func NewDisjSet(size int) DisjSet {
	dj := make([]int, size)
	for i := 0; i < size; i++ {
		dj[i] = -1
	}
	return dj
}

//union by height
func UnionSet(dj DisjSet, root1 int, root2 int) {
	if dj[root1] > dj[root2] {
		dj[root1] = root2
	} else {
		if dj[root1] == dj[root2] {
			dj[root1]--
		}
		dj[root2] = root1
	}
}

func Find(dj DisjSet, x int) int {
	if dj[x] <= 0 {
		return x
	} else {
		return Find(dj, dj[x])
	}
}
