package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewNode(k int) *Node {
	n := new(Node)
	n.Val = k
	return n
}

func (rootNode *Node) Gt(n *Node) bool {
	if rootNode.Val > n.Val {
		return true
	}
	return false
}

func (rootNode *Node) InsertNode(n *Node) {

	if rootNode == nil {
		return
	}
	defer func() {
		fmt.Printf("insert %d ok \n", n.Val)
	}()
	if rootNode.Gt(n) {
		// root > b
		if rootNode.Left == nil {
			rootNode.Left = n
			return
		}
		rootNode.Left.InsertNode(n)

	} else {
		if rootNode.Right == nil {
			rootNode.Right = n
			return
		}
		rootNode.Right.InsertNode(n)
	}

}

func PreOrder(rootNode *Node) {

	if rootNode == nil {
		//fmt.Println("is a Empty Node")
		return
	}

	PreOrder(rootNode.Left)
	fmt.Println(rootNode.Val)
	PreOrder(rootNode.Right)
}
func InOrder(rootNode *Node) {

	if rootNode == nil {
		//fmt.Println("is a Empty Node")
		return
	}

	fmt.Println(rootNode.Val)
	InOrder(rootNode.Left)
	InOrder(rootNode.Right)
}

func PostOrder(rootNode *Node) {

	if rootNode == nil {
		//fmt.Println("is a Empty Node")
		return
	}
	PostOrder(rootNode.Right)
	fmt.Println(rootNode.Val)
	PostOrder(rootNode.Left)
}

func binaryTreeToArray(rootNode *Node) []int {
	if rootNode == nil {
		return nil
	}
	arr := []int{rootNode.Val}
	l := binaryTreeToArray(rootNode.Left)
	r := binaryTreeToArray(rootNode.Right)
	arr = append(arr, l...)
	arr = append(arr, r...)
	return arr
}

func arrayToBST(rootNode *Node, arr []int, i *int) {
	if rootNode == nil {
		return
	}

	arrayToBST(rootNode.Left, arr, i)
	//fmt.Printf("old val:%d,new val:%d,i:%d\n", rootNode.Val, arr[*i], *i)
	rootNode.Val = arr[*i]
	*i++
	arrayToBST(rootNode.Right, arr, i)

}

func arrayToBinaryTree(rootNode *Node) {
	if rootNode == nil {
		return
	}
	arr := binaryTreeToArray(rootNode)

	i := 0
	fmt.Println("before quickSort:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("after quickSort:", arr)
	arrayToBST(rootNode, arr, &i)
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
