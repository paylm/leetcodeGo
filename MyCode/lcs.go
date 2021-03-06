package main

import (
	"fmt"
	"sort"
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
candiates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
**/
func combinationSum(candidates []int, target int) [][]int {
	//fmt.Printf("candidates:%v,target:%d\n", candidates, target)
	res := [][]int{}
	//sort3Parttion(candidates, 0, len(candidates)-1)
	sort.Ints(candidates)
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

/***
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
**/

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	//sort3Parttion(candidates, 0, len(candidates)-1)
	sort.Ints(candidates)
	for i := 0; i < len(candidates); i++ {
		if target == candidates[i] {
			res = append(res, []int{candidates[i]})
		}
		if candidates[i] > target {
			break
		}

		res = append(res, combinationx2(candidates, target-candidates[i], i+1, []int{candidates[i]})...)
	}
	//fmt.Printf("res:%v\n", res)
	return uniqueArr(res)
}

func combinationx2(candidates []int, target int, idx int, temp []int) [][]int {
	res := [][]int{}
	for i := idx; i < len(candidates); i++ {
		if target == candidates[i] {
			//fmt.Printf("found a sultion:%v\n", append(temp, candidates[i]))
			tgarr := make([]int, len(temp)+1)
			copy(tgarr, temp)
			tgarr[len(temp)] = candidates[i]
			res = append(res, tgarr)
			continue
		}
		if target < candidates[i] {
			break
		}
		res = append(res, combinationx2(candidates, target-candidates[i], i+1, append(temp, candidates[i]))...)
	}
	return res
}

func uniqueArr(src [][]int) [][]int {
	uqMap := make(map[string]int)
	res := [][]int{}
	for _, iarr := range src {
		key := ""
		for _, k := range iarr {
			key = fmt.Sprintf("%s%d", key, k)
		}
		_, ok := uqMap[key]
		if !ok {
			uqMap[key] = 1
			res = append(res, iarr)
		}
	}
	return res
}

/**
https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others/
在一个给定的数组nums中，总是存在一个最大元素 。

查找数组中的最大元素是否至少是数组中每个其他数字的两倍。

如果是，则返回最大元素的索引，否则返回-1。

示例 1:

输入: nums = [3, 6, 1, 0]
输出: 1
解释: 6是最大的整数, 对于数组中的其他整数,
6大于数组中其他元素的两倍。6的索引是1, 所以我们返回1.


示例 2:

输入: nums = [1, 2, 3, 4]
输出: -1
解释: 4没有超过3的两倍大, 所以我们返回 -1.

**/
func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	res := 1
	max := nums[0]
	max_i := 0
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			if max*2 <= nums[i] {
				res = 1
			} else {
				res = -1
			}
			max = nums[i]
			max_i = i
		} else if max < nums[i]*2 {
			res = -1
		}
	}
	if res == -1 {
		return res
	}

	return max_i
}

/***
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

**/
func maxSubArray(nums []int) int {
	res := make([]int, len(nums))
	res[0] = nums[0]
	if len(nums) < 2 {
		return res[0]
	}
	//res[1] = max(nums[0],nums[1])
	maxVal := res[0]
	for i := 1; i < len(nums); i++ {
		res[i] = max(res[i-1]+nums[i], nums[i])
		maxVal = max(maxVal, res[i])
	}

	//fmt.Printf("res:%v\n", res)
	return maxVal
}

/**
https://leetcode-cn.com/problems/3sum/
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
**/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	visit := make(map[string]bool)
	//fmt.Printf("after sort:%v\n", nums)
	for i := 0; i < len(nums); i++ {
		//防止重复计算
		if i > 1 && nums[i] == nums[i-1] {
			continue
		}

		s, l := i+1, len(nums)-1
		for ; s < l; s++ {
			k := nums[i] + nums[s]
			_, ok := visit[fmt.Sprintf("%d|%d", nums[i], nums[s])]
			if !ok {
				visit[fmt.Sprintf("%d|%d", nums[i], nums[s])] = true
			} else {
				continue
			}
			il := l
			for {
				if il <= s {
					break
				}
				if -k == nums[il] {
					res = append(res, []int{nums[i], nums[s], nums[il]})
					l = il
					break
				}
				il--
			}
		}
	}
	//fmt.Printf("res:%v,visit:%v\n", res, visit)
	return uniqueArr(res)
}

func threeSum1(nums []int) [][]int {
	res := [][]int{}
	visit := make(map[string]bool)
	sort.Ints(nums)
	//fmt.Println("nums:", nums)
	for i := 0; i < len(nums); i++ {
		//防止重复计算
		if i > 1 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			_, ok := visit[fmt.Sprintf("%d|%d", nums[i], nums[j])]
			if !ok {
				visit[fmt.Sprintf("%d|%d", nums[i], nums[j])] = true
			} else {
				continue
			}
			e := res_threeSum(nums, j+1, nums[i]+nums[j], []int{nums[i], nums[j]})
			if e != nil {
				res = append(res, e)
			}
		}
	}
	//fmt.Printf("res:%v,visit:%v\n", res, visit)
	//return uniqueArr(res)
	fmt.Printf("res:%v\n", res)
	return res
}

func res_threeSum(nums []int, idx int, target int, temp []int) []int {
	for i := len(nums) - 1; i >= idx; i-- {
		if nums[i]+target == 0 {
			//fmt.Printf("found a result:%v\n", target)
			return append(temp, nums[i])
		}
	}
	return nil
}

/**
https://leetcode-cn.com/problems/longest-common-prefix/
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
**/
func longestCommonPrefix(strs []string) string {

	if len(strs) < 1 {
		return ""
	}
	if len(strs) < 2 {
		return strs[0]
	}
	res := strings.Split(strs[0], "")
	for i := 1; i < len(strs); i++ {
		res = comStr(res, strings.Split(strs[i], ""))
		if len(res) == 0 {
			return ""
		}
	}
	return strings.Join(res, "")
}

//str1
func comStr(str1 []string, str2 []string) []string {
	var ks1, ks2 []string
	if len(str1) < len(str2) {
		ks1 = str1
		ks2 = str2
	} else {
		ks1 = str2
		ks2 = str1
	}
	for i := 0; i < len(ks1); i++ {
		//fmt.Printf("compare %s <=>%s \n", ks1[i], ks2[i])
		if strings.Compare(ks1[i], ks2[i]) != 0 {
			//fmt.Printf("ks1:%v,i:%d\n", ks1[:i], i)
			return ks1[:i]
		}
	}
	return ks1
}

/***
https://leetcode-cn.com/problems/climbing-stairs/
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。
**/
func climbStairs1(n int) int {
	//此方法递归入stack 太低率，有重复计算的情况
	if n == 0 {
		//fmt.Printf("found a path\n")
		return 1
	}

	k := 0

	if n >= 1 {
		k = k + climbStairs1(n-1)
	}

	if n >= 2 {
		k = k + climbStairs1(n-2)
	}
	return k
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 2
	for i := 2; i < n; i++ {
		arr[i] = arr[i-2] + arr[i-1]
	}
	return arr[n-1]
}

/**
https://leetcode-cn.com/problems/integer-break/
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。
**/
func integerBreak(n int) int {
	//由数学计算公式推度出,拆分为 3  4  2  这些数值时乘积最大，优先获分为3 ， 不能拆分出1
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	ret := 1
	for {
		if n > 4 {
			ret = ret * 3
			n = n - 3
		} else {
			ret = ret * n
			break
		}
	}

	return ret
}

func arrX(nums []int) [][]int {
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
		combinArr(nums, 0, i)
	}
	return res
}

func combinArr(nums []int, start int, skip int) []int {
	if start >= len(nums) {
		fmt.Println()
		return nil
	}
	if start != skip {
		fmt.Printf("%d ", nums[start])
	}
	start++
	combinArr(nums, start, skip)
	return nil
}

/**
https://leetcode-cn.com/problems/permutations-ii/
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
**/
func permuteUnique(nums []int) [][]int {
	ret := [][]int{}
	for i := 0; i < len(nums); i++ {
		xpert(nums, i+1, i+1)
	}
	return ret
}

func xpert(nums []int, start int, skip int) []int {

	for i := start; i < len(nums); i++ {
		if i == skip {
			continue
		}
		//xpert(nums, i)
	}
	return nil
}

/**
https://leetcode-cn.com/problems/house-robber-iii/
计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

示例 1:

输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

输出: 7
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
示例 2:

输入: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

输出: 9
解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.

链接：https://leetcode-cn.com/problems/house-robber-iii
**/
//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(Val int) *TreeNode {
	n := new(TreeNode)
	n.Val = Val
	return n
}

func buildTreeByArr(arr []int, i int) *TreeNode {
	//left = 2p+1 ,right = 2p+2
	if i >= len(arr) || arr[i] < 0 {
		return nil
	}
	r := NewTreeNode(arr[i])
	r.Left = buildTreeByArr(arr, 2*i+1)
	r.Right = buildTreeByArr(arr, 2*i+2)
	return r
}

func inOrderTree(root *TreeNode) {
	if root == nil {
		fmt.Printf("nil\n")
		return
	}
	fmt.Printf("%d\n", root.Val)
	inOrderTree(root.Left)
	inOrderTree(root.Right)
}

//todo
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	res = 0
	rob_recurse(root, 0, 0, &res)
	return res
}

//assume root is not nil
func maxSub(root *TreeNode) int {
	if root.Left == nil && root.Right != nil {
		return root.Right.Val
	} else if root.Left != nil && root.Right == nil {
		return root.Left.Val
	} else if root.Left == nil && root.Right == nil {
		return 0
	} else if root.Left.Val > root.Right.Val {
		return root.Left.Val
	} else {
		return root.Right.Val
	}
}

func rob_recurse(root *TreeNode, prev1 int, prev2 int, res *int) {
	if root == nil {
		return
	}
	k := maxSub(root)

	if k+prev2 > root.Val+prev1 {
		// no select current node
		//fmt.Printf("no select %v\n", root)
		rob_recurse(root.Left, prev2, prev2, res)
		rob_recurse(root.Right, prev2, prev2, res)
	} else {
		// select current node
		fmt.Printf("select %v\n", root)
		*res = *res + root.Val
		rob_recurse(root.Left, prev2, prev1+root.Val, res)
		rob_recurse(root.Right, prev2, prev1+root.Val, res)
	}

	//return rob_recurse(root.Left, ) + rob_recurse(root.Right, )

}

func rob_recurse1(root *TreeNode, prev1 int, prev2 int) int {
	if root == nil {
		return 0
	}

	if root.Val+prev1 > prev2 {
		prev1 = prev2
		prev2 = root.Val + prev1
	} else {

	}
	if root.Left == nil && root.Right == nil {
		fmt.Printf("prev1:%d,prev2:%d at node %v\n", prev1, prev2, root)
	}

	return rob_recurse1(root.Left, prev1, prev2) + rob_recurse1(root.Right, prev1, prev2)
}

/**
https://leetcode-cn.com/problems/video-stitching/
你将会获得一系列视频片段，这些片段来自于一项持续时长为 T 秒的体育赛事。这些片段可能有所重叠，也可能长度不一。

视频片段 clips[i] 都用区间进行表示：开始于 clips[i][0] 并于 clips[i][1] 结束。我们甚至可以对这些片段自由地再剪辑，例如片段 [0, 7] 可以剪切成 [0, 1] + [1, 3] + [3, 7] 三部分。

我们需要将这些片段进行再剪辑，并将剪辑后的内容拼接成覆盖整个运动过程的片段（[0, T]）。返回所需片段的最小数目，如果无法完成该任务，则返回 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/video-stitching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
示例 1：

输入：clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
输出：3
解释：
我们选中 [0,2], [8,10], [1,9] 这三个片段。
然后，按下面的方案重制比赛片段：
将 [1,9] 再剪辑为 [1,2] + [2,8] + [8,9] 。
现在我们手上有 [0,2] + [2,8] + [8,10]，而这些涵盖了整场比赛 [0, 10]。

提示：

1 <= clips.length <= 100
0 <= clips[i][0], clips[i][1] <= 100
0 <= T <= 100
**/
func videoStitching(clips [][]int, T int) int {
	// 从0->n -> T 的区间里找最找长区域
	start := 0
	count := 0 // 寻找次数
	for {
		zone := findMaxZone(clips, start)
		count += 1
		if zone == nil {
			return -1
		}
		//fmt.Printf("zone is %v\n", zone)
		if zone[1] >= T {
			return count
		}
		start = zone[1]
	}
}

//寻找以<=n 以开头的最大区间
func findMaxZone(clips [][]int, start int) []int {
	end := start
	for i := 0; i < len(clips); i++ {
		if clips[i][0] <= start && clips[i][1] > end {
			end = clips[i][1]
		}
	}
	if end == start {
		return nil
	}
	return []int{start, end}
}
