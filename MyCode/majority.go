package main

/**
求众数
数组出现次数大于为 n/2
**/
func majorityElement1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	majority := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			majority = nums[i]
			count++
			continue
		}
		if nums[i] == majority {
			count++
		} else {
			count--
		}
	}
	return majority
}

/**
https://leetcode-cn.com/problems/majority-element-ii/
求众数
给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。

示例 1:

输入: [3,2,3]
输出: [3]
示例 2:

输入: [1,1,1,3,3,2,2,2]
输出: [1,2]

还没完成
**/
func majorityElement(nums []int) []int {
	//用摩尔投票法
	ret := []int{}
	var majority1, majority2 int
	count1, count2 := 0, 0
	//投票处理后剩下的数字majority1 majority2
	for i := 0; i < len(nums); i++ {
		if nums[i] == majority1 {
			count1++
		} else if nums[i] == majority2 {
			count2++
		} else if count1 == 0 {
			majority1 = nums[i]
			count1++
		} else if count2 == 0 {
			majority2 = nums[i]
			count2++
		} else {
			count1--
			count2--
		}
	}
	//fmt.Printf("majority1:%d,count1:%d,majority2:%d,count2:%d\n", majority1, count1, majority2, count2)
	count1 = 0
	count2 = 0
	//重新计数拿出元素的出现次数
	for i := 0; i < len(nums); i++ {
		if nums[i] == majority1 {
			count1++
		} else if nums[i] == majority2 {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		ret = append(ret, majority1)
	}
	if count2 > len(nums)/3 {
		ret = append(ret, majority2)
	}
	return ret
}
