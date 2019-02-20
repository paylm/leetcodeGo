package main

type DisjSet struct {
	set  []int
	size int
}

func NewDisjSet(size int) *DisjSet {
	d := new(DisjSet)
	d.size = size
	d.set = make([]int, size)
	for i := 0; i < size; i++ {
		d.set[i] = -1 //default depeer is -1
	}
	return d
}

//assume root1,root2 are root
func (d *DisjSet) UnionSet(root1 int, root2 int) {
	if d.set[root1] > d.set[root2] {
		d.set[root1] = root2
	} else {
		if d.set[root1] == d.set[root2] && root1 < root2 {
			d.set[root2] = root1
			d.set[root1]--
		} else {
			d.set[root1] = root2
			d.set[root2]--
		}

	}
}

func (d *DisjSet) Find(e int) bool {

	return true
}
