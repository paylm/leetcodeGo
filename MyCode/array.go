package main

import "fmt"

/**
https://leetcode-cn.com/problems/set-matrix-zeroes/

给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

输入:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]

一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个常数空间的解决方案吗？
**/
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	nums := []int{} //额外空间记录这些元素
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				k := i*n + j
				fmt.Printf("read matrix form %d:%d,k:%d\n", i, j, k)
				nums = append(nums, k)
			}
		}
	}
	for _, v := range nums {
		i, j := v/n, v%n
		fmt.Printf("i:%d,j:%d\n", i, j)
		for s := 0; s < n; s++ {
			matrix[i][s] = 0
		}
		for s := 0; s < m; s++ {
			matrix[s][j] = 0
		}
	}
}
