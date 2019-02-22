package main

import (
	"fmt"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func commonSubLen(str1, str2 string) int {
	//fmt.Printf("commonSubArr for :%s ,%s\n", str1, str2)
	arr1 := strings.Split(str1, "")
	arr2 := strings.Split(str2, "")
	cmLen := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i-1] == arr2[j-1] {
				cmLen = cmLen + 1
				break
			}
		}
	}
	return cmLen
}

// time T(m*n) , space O(m*n)
func commonSubArr(str1, str2 string) []string {
	//fmt.Printf("commonSubArr for :%s ,%s\n", str1, str2)
	arr1 := strings.Split(str1, "")
	arr2 := strings.Split(str2, "")
	var lcs [][]int
	res := []string{}
	for i := 0; i < len(arr1)+1; i++ {
		temp := make([]int, len(arr2)+1)
		lcs = append(lcs, temp)
	}
	for i := 0; i < len(arr1)+1; i++ {
		for j := 0; j < len(arr2)+1; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if arr1[i-1] == arr2[j-1] {
				lcs[i][j] = max(lcs[i][j], lcs[i-1][j-1]+1)
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}
	fmt.Println("----lcs-----")
	for i := 0; i < len(lcs); i++ {
		fmt.Println(lcs[i])
		if i > 0 && lcs[i][len(arr2)] > lcs[i-1][len(arr2)] {
			//fmt.Println(arr1[i-1])
			res = append(res, arr1[i-1])
		}
	}
	return res
}

// time O(m+n) , space O(m*n)
func commonSubArr1(str1, str2 string) []string {
	//fmt.Printf("commonSubArr for :%s ,%s\n", str1, str2)
	arr1 := strings.Split(str1, "")
	arr2 := strings.Split(str2, "")
	res := []string{}
	for i := 0; i < len(arr1); i++ {
		for j := i; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				res = append(res, arr1[i])
				break
			}
		}
	}
	return res
}
