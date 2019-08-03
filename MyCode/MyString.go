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
https://leetcode-cn.com/problems/summary-ranges/

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

/*
把ip 地址转化为int , format like xxx.xx.xx.xx
利用位运算处理
*/
func ipToInt(ip string) int {
	ips := strings.Split(ip, ".")
	if len(ips) != 4 {
		fmt.Printf("%v is not legal ip addr\n", ip)
		return -1
	}
	fmt.Printf("ips :%v\n", ips)
	intIp := 0
	for _, i := range ips {
		intI, err := strconv.Atoi(i)
		if err != nil {
			fmt.Printf("convert %v throw err:%v\n", i, err)
			return -1
		} else if intI > 255 {
			fmt.Printf("convert %v throw err , mask should less than 257\n", i)
			return -1
		}

		intIp = (intIp << 8) + intI
	}
	fmt.Printf("ip :%v convert to %d\n", ip, intIp)
	return intIp
}

/*
把int 反转为化为ip addr , format like xxx.xx.xx.xx
利用位运算处理
*/
func intToIP(ip int) string {
	ips := []string{"", "", "", ""}
	mask := 2<<7 - 1
	for i := len(ips) - 1; i >= 0; i-- {
		ips[i] = fmt.Sprintf("%d", (ip & mask))
		ip >>= 8
	}
	return strings.Join(ips, ".")
}

/**
https://leetcode-cn.com/problems/utf-8-validation/
UTF-8 中的一个字符可能的长度为 1 到 4 字节，遵循以下的规则：

对于 1 字节的字符，字节的第一位设为0，后面7位为这个符号的unicode码。
对于 n 字节的字符 (n > 1)，第一个字节的前 n 位都设为1，第 n+1 位设为0，后面字节的前两位一律设为10。剩下的没有提及的二进制位，全部为这个符号的unicode码。
这是 UTF-8 编码的工作方式：

   Char. number range  |        UTF-8 octet sequence
      (hexadecimal)    |              (binary)
   --------------------+---------------------------------------------
   0000 0000-0000 007F | 0xxxxxxx
   0000 0080-0000 07FF | 110xxxxx 10xxxxxx
   0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
   0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
给定一个表示数据的整数数组，返回它是否为有效的 utf-8 编码。

注意:
输入是整数数组。只有每个整数的最低 8 个有效位用来存储数据。这意味着每个整数只表示 1 字节的数据。

示例 1:

data = [197, 130, 1], 表示 8 位的序列: 11000101 10000010 00000001.

返回 true 。
这是有效的 utf-8 编码，为一个2字节字符，跟着一个1字节字符。
**/
func validUtf8(data []int) bool {
	mask := 2 << 6 // 1 0 0 0 0 0 0 0
	hr := 0        // 是否为长度字符位
	for i := 0; i < len(data); i++ {
		if i == hr {
			n := validSize(data[i])
			if n == 0 || n+i > len(data) {
				return false
			}
			//set next hr
			hr = n + i
			fmt.Printf("%d =>size:%d,next:%d\n", data[i], n, hr)
		} else if data[i]&mask != mask {
			return false
		}
	}
	return true
}

//验证最高位长度，最大为4,最小1 ， 0 表示不合法
func validSize(i int) int {
	mask1 := 2 << 6
	mask := 3 << 6 // 1 0 0 0 0 0 0 0
	fmt.Printf("mask=%v,i=%d\n", mask, i)
	if i&(mask1) == 0 {
		return 1
	}
	j := 0
	for s := 2; s <= 4; s++ {
		if i&mask == mask {
			i <<= 1
			j = s
		} else {
			break
		}
	}
	if (i<<1)&mask1 == mask1 {
		//第 n+1 位设为0
		return 0
	}
	return j
}
