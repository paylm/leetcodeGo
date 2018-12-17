package main

import "fmt"

type AvlNode struct {
	Val    int
	Height int
	Left   *AvlNode
	Right  *AvlNode
}

func NewAvlNode(v int) *AvlNode {
	n := new(AvlNode)
	n.Val = v
	n.Height = 0
	return n
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Height(n *AvlNode) int {
	if n == nil {
		return 0
	} else {
		return n.Height
	}
}

func getBalance(n *AvlNode) int {
	if n == nil {
		return 0
	}

	return Height(n.Left) - Height(n.Right)
}

//插入数据，并带avl
func Insert(n *AvlNode, x int) *AvlNode {
	if n == nil {
		return NewAvlNode(x)
	}

	if x > n.Val {
		n.Right = Insert(n.Right, x)
	} else if x < n.Val {
		n.Left = Insert(n.Left, x)
	} else {
		// x aready exsits
		return n
	}

	n.Height = Max(Height(n.Left), Height(n.Right)) + 1
	balance := getBalance(n)

	// left left rotate
	if balance > 1 && x < n.Left.Val {
		return RightRotate(n)
	}

	// right right case
	if balance < -1 && x > n.Right.Val {
		return LeftRotate(n)
	}

	// right left case
	if balance > 1 && x > n.Left.Val {
		n.Left = LeftRotate(n.Left)
		return RightRotate(n)
	}

	//left right case
	if balance < -1 && x < n.Right.Val {
		n.Right = RightRotate(n.Right)
		return LeftRotate(n)
	}
	/* return the (unchanged) node pointer */
	return n
}

/**
k1 old root
k2 new root

   k1                 k2
     \              /    \
	  k2    ==>   k1     xxx
	    \
		 xx
**/
func LeftRotate(k1 *AvlNode) *AvlNode {
	fmt.Printf("LeftRotate %v\n", k1)
	k2 := k1.Right
	k1.Right = k2.Left
	k2.Left = k2

	k1.Height = Max(Height(k1.Left), Height(k1.Right)) + 1
	k2.Height = Max(Height(k2.Left), Height(k2.Right)) + 1

	return k2
}

/**
k1 old root
k2 new root
**/
func RightRotate(k1 *AvlNode) *AvlNode {
	fmt.Printf("RightRotate %v , left :%v , right:%v\n", k1, k1.Left, k1.Right)
	k2 := k1.Left
	k1.Left = k2.Right
	k2.Left = k1
	k1.Height = Max(Height(k1.Left), Height(k1.Right)) + 1
	k2.Height = Max(Height(k2.Left), Height(k2.Right)) + 1

	return k2
}

func AvlTraverse(n *AvlNode) {
	if n == nil {
		return
	}

	fmt.Printf("%d ", n.Val)
	AvlTraverse(n.Left)
	AvlTraverse(n.Right)
}
