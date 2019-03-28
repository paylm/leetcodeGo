package main

import "fmt"

//https://leetcode-cn.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	nums := make([]int, len(grid)>>2)
	l := len(grid)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			n := i*l + j
			if j-1 > 0 && grid[i][j] == 1 && grid[i][j-1] == 1 {
				jn := i*l + j - 1
				UnionSet(nums, jn, n)
			}
			if i-1 > 0 && grid[i][j] == 1 && grid[i-1][j] == 1 {
				in := (i-1)*l + j
				UnionSet(nums, in, n)
			}
		}
	}
	fmt.Printf("nums:%v\n", nums)
	return 1
}

func UnionSet(r []int, i1 int, i2 int) {
	ri1 := FindRoot(r, i1)
	r[i2] = ri1
}

func FindRoot(r []int, i int) int {
	if r[i] == 0 {
		return i
	}
	return FindRoot(r, r[i])
}
