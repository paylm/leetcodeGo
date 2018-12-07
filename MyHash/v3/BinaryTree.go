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
