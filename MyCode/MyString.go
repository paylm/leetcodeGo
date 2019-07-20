package main

import (
	"fmt"
	"strconv"
	"strings"
)

/***
https://leetcode-cn.com/problems/restore-ip-addresses/
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
**/
func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	data := strings.Split(s, "")
	res := recurseIPAdrr(0, 0, data, []string{})
	return res
}

func recurseIPAdrr(z int, s int, data []string, src []string) []string {
	if z > 4 || s > len(data) {
		return nil
	}
	if z == 4 && s == len(data) {
		return []string{strings.Join(src, ".")}
	}
	var res []string
	n := ""
	for i := s; i < len(data); i++ {
		n = fmt.Sprintf("%s%s", n, data[i])
		k, _ := strconv.Atoi(n)
		if k > 255 || (k < (len(n)-1)*10) {
			break
		} else {
			//可用
			if rs := recurseIPAdrr(z+1, i+1, data, append(src, n)); rs != nil {
				if res == nil {
					res = rs
				} else {
					res = append(res, rs...)
				}
			}
		}
	}
	return res
}

/**
https://leetcode-cn.com/problems/summary-ranges/<Paste>

给定一个无重复元素的有序整数数组，返回数组区间范围的汇总。

示例 1:

输入: [0,1,2,4,5,7]
输出: ["0->2","4->5","7"]
解释: 0,1,2 可组成一个连续的区间; 4,5 可组成一个连续的区间。
**/
func summaryRanges(nums []int) []string {
	if len(nums) < 1 {
		return nil
	}
	s1, s2 := 0, 1
	res := []string{fmt.Sprintf("%d", nums[s1])}
	for i := s2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			//连续区间
			s2 = i
			res[len(res)-1] = fmt.Sprintf("%d->%d", nums[s1], nums[s2])
		} else {
			//fmt.Printf("other s1:%d,s2:%d,zone:%d->%d\n", s1, s2, nums[s1], nums[i-1])
			s1 = i
			s2 = s1
			res = append(res, fmt.Sprintf("%d", nums[i]))
		}
	}
	fmt.Printf("res :%v,nums:%v\n", res, nums)
	return res
}
