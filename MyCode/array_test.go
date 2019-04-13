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
