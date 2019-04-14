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

/**
https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
**/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	s, e := 0, len(nums)
	var m int
	for {
		if s >= e {
			m = s
			break
		}
		if nums[s] == target {
			m = s
			break
		}
		m = (s + e) / 2
		if nums[m] == target {
			break
		}
		if nums[m] > target {
			e = m - 1
		} else if nums[m] < target {
			s = m + 1
		}

	}

	//fmt.Printf("s:%d,e:%dm = %d =>%d ,data =  %v\n", s, e, m, nums[m], nums)
	if m == len(nums) || nums[m] != target {
		return []int{-1, -1}
	}

	i, j, k := m, m, m
	res := []int{}
	for {
		if i-1 < 0 {
			break
		}
		if nums[i] != target {
			break
		}
		j = i
		i--
	}
	res = append(res, j)

	for {
		if k+1 > len(nums) {
			break
		}
		if nums[k] != target {
			break
		}
		k++
	}
	res = append(res, k-1)
	return res
}
