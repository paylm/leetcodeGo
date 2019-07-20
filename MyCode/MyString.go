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
