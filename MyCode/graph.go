package main

import "fmt"

/**
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
**/
func minPathSum(grid [][]int) int {
	xl := len(grid)
	yl := len(grid[0])
	//用图运算方法处理
	cost := make([]int, xl*yl)
	visit := make([]bool, xl*yl)

	for i, _ := range cost {
		cost[i] = 10000000
		visit[i] = false
	}
	var i int
	target := xl*yl - 1
	cost[0] = grid[0][0]
	for {
		i = findMinIndex(cost, visit) //换成最小堆能提升效率
		if visit[i] == true {
			break
		}
		if i == target {
			fmt.Printf("find the result %d, cost :%d\n", target, cost[i])
			return cost[i]
		}
		visit[i] = true
		x := i / yl
		y := i % yl
		//fmt.Printf("x:%d,y:%d,i:%d,cost:%d\n", x, y, i, cost[i])
		//right
		if y < yl-1 {
			k := x*yl + y + 1
			if !visit[k] && cost[i]+grid[x][y+1] < cost[k] {
				cost[k] = cost[i] + grid[x][y+1]
			}
		}
		//down
		if x < xl-1 {
			kd := (x+1)*yl + y
			if !visit[kd] && cost[i]+grid[x+1][y] < cost[kd] {
				cost[kd] = cost[i] + grid[x+1][y]
			}
		}
	}
	fmt.Printf("cost:%v,target:%d\n", cost, target)
	return 0
}

func findMinIndex(arr []int, visit []bool) int {
	min := 100000
	idx := 0
	for i, _ := range arr {
		if visit[i] == false {
			if arr[i] < min {
				min = arr[i]
				idx = i
			}
		}
	}
	return idx
}
