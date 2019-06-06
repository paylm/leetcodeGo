package main

import "testing"

func Test_MinPathSum(t *testing.T) {
	data := []struct {
		arr    [][]int
		result int
	}{
		{arr: [][]int{{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1}}, result: 7},
		{arr: [][]int{{1, 2, 5}, {3, 2, 1}}, result: 6},
	}

	for _, d := range data {
		res := minPathSum(d.arr)
		if res != d.result {
			t.Errorf("test fail at %v , return %d\n", d, res)
		}
	}
}
