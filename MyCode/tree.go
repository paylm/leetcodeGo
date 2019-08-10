package Mycode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/***
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---

链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
**/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	res = recurseTree(root, res, 0)
	return res
}

func recurseTree(root *TreeNode, res []int, i int) []int {
	if root == nil {
		return res
	}

	if len(res) == i {
		res = append(res, -1)
		res[i] = root.Val
	}
	res = recurseTree(root.Right, res, i+1)
	res = recurseTree(root.Left, res, i+1)
	return res
}

func buildTree(res []int, k int) *TreeNode {
	if len(res) < k || res[k] == -1 {
		return nil
	}
	t := TreeNode{res[k], nil, nil}
	t.Left = buildTree(res, 2*k+1)
	t.Right = buildTree(res, 2*k+2)
	return &t

}
