package main

import "fmt"

/**https://leetcode-cn.com/problems/number-of-islands/
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1

**/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	nums := make([]int, len(grid)*len(grid[0])+1)
	//fmt.Printf("input grid:%v,nums:%v\n", grid, nums)
	for i := 0; i < len(grid); i++ {
		l := len(grid[i])
		for j := 0; j < l; j++ {
			n := i*l + j + 1
			//fmt.Printf("i:%d j:%d n:%d\n", i, j, n)
			if j > 0 && grid[i][j] == 1 {
				if grid[i][j-1] == 1 {
					jn := i*l + j
					UnionSet(nums, jn, n)
					continue
				}
			}
			if i > 0 && grid[i][j] == 1 {
				if grid[i-1][j] == 1 {
					in := (i-1)*l + j + 1
					UnionSet(nums, in, n)
					continue
				}
			}
			if grid[i][j] == 1 {
				nums[n] = n
			}
		}
	}
	resMap := make(map[int]int)
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if _, ok := resMap[nums[i]]; !ok {
			res++
			resMap[nums[i]] = 1
		}
	}
	fmt.Printf("nums:%v\n", nums)
	return res
}

func UnionSet(r []int, i1 int, i2 int) {
	//fmt.Printf("UnionSet i1:%d,i2:%d\n", i1, i2)
	ri1 := FindRoot(r, i1)
	r[i2] = ri1
}

func FindRoot(r []int, i int) int {
	if r[i] == i {
		return i
	}
	if r[i] <= 0 {
		return i
	} else {
		return FindRoot(r, r[i])
	}
}
