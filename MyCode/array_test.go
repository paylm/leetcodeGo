package main

import "testing"

func Test_SetZeroes(t *testing.T) {
	data := []struct {
		input  [][]int
		output [][]int
	}{
		{input: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, output: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
	}
	for _, d := range data {
		setZeroes(d.input)
		for i := 0; i < len(d.input); i++ {
			for j := 0; j < len(d.input[i]); j++ {
				if d.input[i][j] != d.output[i][j] {
					t.Errorf("test fail at %v\n", d)
				}
			}
		}
	}
}

func Test_SearchRange(t *testing.T) {
	data := []struct {
		input  []int
		target int
		output []int
	}{
		{input: []int{5, 7, 7, 8, 8, 10}, target: 8, output: []int{3, 4}},
		{input: []int{5, 7, 7, 8, 8, 10}, target: 6, output: []int{-1, -1}},
		{input: []int{10}, target: 8, output: []int{-1, -1}},
		{input: []int{1}, target: 1, output: []int{0, 0}},
		{input: []int{1, 4}, target: 4, output: []int{1, 1}},
		{input: []int{4, 4}, target: 4, output: []int{0, 1}},
		{input: []int{2, 2}, target: 3, output: []int{-1, -1}},
		{input: []int{0, 1, 2, 3, 4, 4, 4}, target: 2, output: []int{2, 2}},
		{input: []int{1, 2, 3, 3, 3, 3, 4, 5, 9}, target: 3, output: []int{2, 5}},
	}
	for _, d := range data {
		res := searchRange(d.input, d.target)
		if d.output[0] != res[0] || d.output[1] != res[1] {
			t.Errorf("test fail at %v,return %v\n", d, res)
		}
	}
}
