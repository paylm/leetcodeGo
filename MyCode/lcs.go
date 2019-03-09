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
	var lcs [][]int
	for i := 0; i < len(arr1)+1; i++ {
		temp := make([]int, len(arr2)+1)
		lcs = append(lcs, temp)
	}
	for i := 0; i < len(arr1)+1; i++ {
		for j := 0; j < len(arr2)+1; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if arr1[i-1] == arr2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}
	return lcs[len(arr1)][len(arr2)]
}

// time T(m*n) , space O(m*n)
func commonSubArr(str1, str2 string) []string {
	//fmt.Printf("commonSubArr for :%s ,%s\n", str1, str2)
	arr1 := strings.Split(str1, "")
	arr2 := strings.Split(str2, "")
	var lcs [][]int
	for i := 0; i < len(arr1)+1; i++ {
		temp := make([]int, len(arr2)+1)
		lcs = append(lcs, temp)
	}
	for i := 0; i < len(arr1)+1; i++ {
		for j := 0; j < len(arr2)+1; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if arr1[i-1] == arr2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}
	//	fmt.Println("----lcs-----")
	//	for i := 0; i < len(lcs); i++ {
	//		fmt.Println(lcs[i])
	//		if i > 0 && lcs[i][len(arr2)] > lcs[i-1][len(arr2)] {
	//			//fmt.Println(arr1[i-1])
	//			res = append(res, arr1[i-1])
	//		}
	//	}
	//print_LCS(arr1, arr2, lcs, len(arr1), len(arr2)) //correct
	//fmt.Println()
	//fmt.Printf("commsize:%d\n", lcs[len(arr1)][len(arr2)])
	res := make([]string, lcs[len(arr1)][len(arr2)])
	rec_LCS(arr1, arr2, lcs, len(arr1), len(arr2), res)
	return res
}

//递归打输最长子序
func print_LCS(str1 []string, str2 []string, lcs [][]int, i int, j int) {
	if i == 0 || j == 0 {
		return
	}
	if lcs[i-1][j] > lcs[i][j-1] {
		//to up
		//fmt.Printf("->%s", str1[i-1])
		print_LCS(str1, str2, lcs, i-1, j)
	} else if lcs[i-1][j] < lcs[i][j-1] {
		//to left
		//fmt.Printf("->%s", str2[j-1])
		print_LCS(str1, str2, lcs, i, j-1)
	} else {
		print_LCS(str1, str2, lcs, i-1, j-1)
		if lcs[i][j] > lcs[i-1][j-1] {
			fmt.Printf("->%s", str2[j-1])
		}
	}
}

//返回最长子序
func rec_LCS(str1 []string, str2 []string, lcs [][]int, i int, j int, res []string) {
	if i == 0 || j == 0 {
		return
	}
	if lcs[i-1][j] > lcs[i][j-1] {
		//to up
		//fmt.Printf("->%s", str1[i-1])
		rec_LCS(str1, str2, lcs, i-1, j, res)
	} else if lcs[i-1][j] < lcs[i][j-1] {
		//to left
		//fmt.Printf("->%s", str2[j-1])
		rec_LCS(str1, str2, lcs, i, j-1, res)
	} else {
		rec_LCS(str1, str2, lcs, i-1, j-1, res)
		if lcs[i][j] > lcs[i-1][j-1] {
			//fmt.Printf("->%s(%d)", str2[j-1], j-1)
			res[lcs[i][j]-1] = str2[j-1]
		}
	}
}

// time O(m+n) , space O(m*n) nofix
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

//从数组找到互不相邻，总数加起来最大的和
func opt_arr(arr []int) int {
	opt := make([]int, len(arr))
	if len(arr) == 0 {
		return 0
	}
	opt[0] = arr[0]
	if len(arr) < 2 {
		return opt[0]
	}
	opt[1] = max(arr[1], arr[0])
	for i := 2; i < len(arr); i++ {
		opt[i] = max(opt[i-1], opt[i-2]+arr[i])
	}
	//fmt.Printf("opt max:%v\n", opt)
	return opt[len(arr)-1]
}

func fib(i int) []int {
	arr := make([]int, i)
	arr[0] = 1
	if i < 2 {
		return arr
	}
	arr[1] = 1
	if i < 3 {
		return arr
	}
	for j := 2; j < i; j++ {
		arr[j] = arr[j-1] + arr[j-2]
	}
	return arr
}

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
**/
func combinationSum(candidates []int, target int) [][]int {
	//fmt.Printf("candidates:%v,target:%d\n", candidates, target)
	res := [][]int{}
	sort3Parttion(candidates, 0, len(candidates)-1)
	for i := 0; i < len(candidates); i++ {
		if target == candidates[i] {
			//fmt.Printf("found a soution: %v i:%d target:%d\n", candidates, i, target)
			res = append(res, []int{candidates[i]})
			continue
		}
		if target < candidates[i] {
			break
		}
		//fmt.Printf("i:%d\n", i)
		res = append(res, combinationx(candidates, target-candidates[i], i, []int{candidates[i]})...)
	}
	//fmt.Println(res)
	return res
}

func combinationx(candidates []int, target int, idx int, temp []int) [][]int {

	res := [][]int{}
	for i := idx; i < len(candidates); i++ {
		if target == candidates[i] {
			//	fmt.Printf("found a soution: %v i:%d target:%d ,idx:%d,temp:%v\n", candidates, i, target, idx, append(temp, candidates[i]))
			//改用copy复制形式，以防后面引用对象修改原始答案
			tgarr := make([]int, len(temp)+1)
			copy(tgarr, temp)
			tgarr[len(temp)] = candidates[i]
			res = append(res, tgarr)
			continue
		}
		if target < candidates[i] {
			break
		}
		res = append(res, combinationx(candidates, target-candidates[i], i, append(temp, candidates[i]))...)
	}
	return res
}
