package MyDisjSet

import (
	"testing"
)

func TestUnionSet(t *testing.T) {
	dj := NewDisjSet(9)
	UnionSet(dj, 5, 7)
	UnionSet(dj, 5, 8)
	UnionSet(dj, 2, 3)

	if dj.S[5] != -2 {
		t.Errorf("test fail , dj.S[%d] should be %d", 5, -2)
	}
	if dj.S[2] != -2 {
		t.Errorf("test fail , dj.S[%d] should be %d", 2, -2)
	}
	//fmt.Println(dj.S)
	UnionSet(dj, 5, 2)
	if dj.S[5] != -3 || dj.S[2] != 5 {
		t.Errorf("test fail at  UnionSet(dj,5,2) ")
	}
	//fmt.Println(dj.S)
}

func TestUnionSetBySize(t *testing.T) {
	dj := NewDisjSet(12)
	UnionSetBySize(dj, 5, 7)
	UnionSetBySize(dj, 5, 8)
	UnionSetBySize(dj, 3, 5)
	UnionSetBySize(dj, 3, 2)
	UnionSetBySize(dj, 3, 1)
	UnionSetBySize(dj, 2, 1)
	UnionSetBySize(dj, 2, 0)
	UnionSetBySize(dj, 5, 9)
	//fmt.Println(dj)
}

func TestFind(t *testing.T) {
	dj := NewDisjSet(12)
	UnionSet(dj, 5, 7)
	UnionSet(dj, 5, 8)
	UnionSet(dj, 2, 3)
	UnionSet(dj, 2, 4)
	UnionSet(dj, 5, 2)
	UnionSet(dj, 9, 10)
	if Find(dj, 2) != Find(dj, 8) {
		t.Errorf("test fail , %d %d belong same set", 2, 8)
	}

	if Find(dj, 1) == Find(dj, 8) {
		t.Errorf("test fail , %d(%d) %d(%d) is not belong a same set", 1, Find(dj, 1), 8, Find(dj, 8))
	}
	//fmt.Println(dj.S)
}

func Test_FindRedundantConnection(t *testing.T) {
	data := []struct {
		input  [][]int
		output []int
	}{
		{input: [][]int{{1, 2}, {1, 3}, {2, 3}}, output: []int{2, 3}},
		{input: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, output: []int{1, 4}},
	}

	for _, ts := range data {
		res := findRedundantConnection(ts.input)
		if len(res) != len(ts.output) || res[0] != ts.output[0] || res[1] != ts.output[1] {
			t.Errorf("test fail at %v, return res:%v\n", ts, res)
		}
	}
}
