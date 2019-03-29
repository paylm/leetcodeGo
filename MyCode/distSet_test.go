package main

import "testing"

func Test_NumIslands(t *testing.T) {
	data := []struct {
		input [][]byte
		res   int
	}{
		{input: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, res: 1},
		{input: [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}, res: 3},
		{input: [][]byte{{}}, res: 0},
		{input: [][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}, res: 1},
	}
	for _, v := range data {
		r := numIslands(v.input)
		if r != v.res {
			t.Errorf("test fail at %v, return :%d\n", v, r)
		}
	}
}
