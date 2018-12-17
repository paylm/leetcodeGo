package main

import "fmt"

type BTNode struct {
	Val   int
	Left  *BTNode
	Right *BTNode
}

func NewBTNode(k int) *BTNode {
	n := new(BTNode)
	n.Val = k
	return n
}

func (rootBTNode *BTNode) Gt(n *BTNode) bool {
	if rootBTNode.Val > n.Val {
		return true
	}
	return false
}

func (rootBTNode *BTNode) InsertBTNode(n *BTNode) {

	if rootBTNode == nil {
		return
	}
	defer func() {
		fmt.Printf("insert %d ok \n", n.Val)
	}()
	if rootBTNode.Gt(n) {
		// root > b
		if rootBTNode.Left == nil {
			rootBTNode.Left = n
			return
		}
		rootBTNode.Left.InsertBTNode(n)

	} else {
		if rootBTNode.Right == nil {
			rootBTNode.Right = n
			return
		}
		rootBTNode.Right.InsertBTNode(n)
	}

}

func PreOrder(rootBTNode *BTNode) {

	if rootBTNode == nil {
		//fmt.Println("is a Empty BTNode")
		return
	}

	PreOrder(rootBTNode.Left)
	fmt.Println(rootBTNode.Val)
	PreOrder(rootBTNode.Right)
}
func InOrder(rootBTNode *BTNode) {

	if rootBTNode == nil {
		//fmt.Println("is a Empty BTNode")
		return
	}

	fmt.Println(rootBTNode.Val)
	InOrder(rootBTNode.Left)
	InOrder(rootBTNode.Right)
}

func PostOrder(rootBTNode *BTNode) {

	if rootBTNode == nil {
		//fmt.Println("is a Empty BTNode")
		return
	}
	PostOrder(rootBTNode.Right)
	fmt.Println(rootBTNode.Val)
	PostOrder(rootBTNode.Left)
}

/**
load the tree to map
        15
	  /    \
   15       25
  /  \    /  \
5     10 20   30

to:
5
15
15 10 20
25
30

**/
func loadBST(rootBTNode *BTNode, s map[int][]int, i int) {
	if rootBTNode == nil {
		return
	}
	v, ok := s[i]
	if !ok {
		s[i] = []int{rootBTNode.Val}
	} else {
		v = append(v, rootBTNode.Val)
		s[i] = v
	}
	loadBST(rootBTNode.Left, s, i-1)
	loadBST(rootBTNode.Right, s, i+1)
}

func showBST(rootBTNode *BTNode) {
	ts := make(map[int][]int)
	loadBST(rootBTNode, ts, 0)
	fmt.Println(ts)
}

func binaryTreeToArray(rootBTNode *BTNode) []int {
	if rootBTNode == nil {
		return nil
	}
	arr := []int{rootBTNode.Val}
	l := binaryTreeToArray(rootBTNode.Left)
	r := binaryTreeToArray(rootBTNode.Right)
	arr = append(arr, l...)
	arr = append(arr, r...)
	return arr
}

func arrayToBST(rootBTNode *BTNode, arr []int, i *int) {
	if rootBTNode == nil {
		return
	}

	arrayToBST(rootBTNode.Left, arr, i)
	//fmt.Printf("old val:%d,new val:%d,i:%d\n", rootBTNode.Val, arr[*i], *i)
	rootBTNode.Val = arr[*i]
	*i++
	arrayToBST(rootBTNode.Right, arr, i)

}

func arrayToBinaryTree(rootBTNode *BTNode) {
	if rootBTNode == nil {
		return
	}
	arr := binaryTreeToArray(rootBTNode)

	i := 0
	//fmt.Println("before quickSort:", arr)
	quickSort(arr, 0, len(arr)-1)
	//fmt.Println("after quickSort:", arr)
	arrayToBST(rootBTNode, arr, &i)
}

func quickSort(a []int, start int, end int) {
	if len(a) == 0 || start == end {
		return
	}

	k := a[start]
	i, j := start, end
	for {
		if i == j {
			break
		}
		for {
			if j > i && a[j] > k {
				j--
			} else {
				break
			}
		}

		for {
			if j > i && a[i] < k {
				i++
			} else {
				break
			}
		}
		temp := a[i]
		a[i] = a[j]
		a[j] = temp
	}

	a[i] = k
	//fmt.Println(a)
	quickSort(a, start, i)
	quickSort(a, j+1, end)
}

func minValBTNode(rootBTNode *BTNode) *BTNode {

	if rootBTNode.Left == nil {
		return rootBTNode
	} else {
		return minValBTNode(rootBTNode.Left)
	}
}

/**
Given a binary search tree and a key, this func deletes the key and return the new root
**/
func delBTNode(rootBTNode *BTNode, k int) *BTNode {

	if rootBTNode == nil {
		return nil
	}

	if rootBTNode.Val > k {
		rootBTNode.Left = delBTNode(rootBTNode.Left, k)
	} else if rootBTNode.Val < k {
		rootBTNode.Right = delBTNode(rootBTNode.Right, k)
	} else {

		//fmt.Println("del BTNode:", rootBTNode.Val)
		//if key is samve sa root's key,then this is the node to be delete
		if rootBTNode.Left == nil {
			temp := rootBTNode.Right
			rootBTNode = nil
			return temp
		} else if rootBTNode.Right == nil {
			temp := rootBTNode.Left
			rootBTNode = nil
			return temp
		}
		//node has two children : Get the inorder succuessor
		temp := minValBTNode(rootBTNode.Right)
		fmt.Println("del BTNode with two sun :", temp)
		rootBTNode.Val = temp.Val
		rootBTNode.Right = delBTNode(rootBTNode.Right, temp.Val)
	}

	return rootBTNode
}

//未完成
func linkToBSTRrecur(lp **LinkNode, n int) *BTNode {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			fmt.Printf("n:%d,err:%v\n", n, err)
	//		}
	//	}()
	if n <= 0 || (*lp) == nil {
		return nil
	}

	left := linkToBSTRrecur(lp, n/2)
	root := NewBTNode((*lp).Val)
	fmt.Printf("v:%T, v :%d \n", *lp, (*lp).Val)
	root.Left = left
	*lp = (*lp).Next
	fmt.Println("*lp:", *lp, "root:", root, "afer left:", root, "right:", (*lp).Val)
	fmt.Println("left n:", n, " right n:", n-n/2+1)
	right := linkToBSTRrecur(lp, n-n/2+1)
	fmt.Println("root:", root, "left:", left, "right:", right)
	root.Right = right
	return root
}

func linkToBST(l *LinkNode) *BTNode {

	if l == nil {
		return nil
	}
	n := countLinkNode(l)
	fmt.Printf("link len:=%d\n", n)
	fmt.Println("head:", l)
	root := linkToBSTRrecur(&l, n)
	return root
}
