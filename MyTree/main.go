package main

import "fmt"

func test1() {

	fmt.Println()
	var q Queue
	q = NewMyQueue()
	q.put(1)
	q.put(2)
	q.put(3)
	q.put(7)
	for {
		if q.empty() {
			break
		}
		fmt.Println(q.pop())
	}

	arr := []int{8, 4, 10, 3, 5, 1, 7}
	var root *BTNode
	for _, v := range arr {
		if root == nil {
			root = NewBTNode(v)
		} else {
			root.InsertBTNode(NewBTNode(v))
		}
	}
}

func testP(p *int) {

	fmt.Println("testP:", *p)
	*p++
}

func main() {
	fmt.Println("vim-go")
	a := []int{0, 1, 2, 3, 4}
	//删除第i个元素
	i := 2
	a = append(a[:i], a[i+1:]...)

	r := NewBTNode(10)
	r.Left = NewBTNode(30)
	r.Right = NewBTNode(15)
	r.Left.Left = NewBTNode(20)
	r.Left.Right = NewBTNode(25)
	r.Right.Right = NewBTNode(5)
	r.Right.Left = NewBTNode(1)
	PreOrder(r)
	arrayToBinaryTree(r)
	PreOrder(r)
	showBST(r)
	fmt.Println(minValBTNode(r))
	delBTNode(r, 15)
	showBST(r)
	PreOrder(r)
	//	p := 2
	//	fmt.Println(&p)
	//	testP(&p)
	//	fmt.Println(p)
	fmt.Println("link to BST")
	l1 := NewLinkNode(1)
	l1.Next = NewLinkNode(3)
	l1.Next.Next = NewLinkNode(5)
	l1.Next.Next.Next = NewLinkNode(7)
	l1.Next.Next.Next.Next = NewLinkNode(9)
	l1.Next.Next.Next.Next.Next = NewLinkNode(11)
	l1.Next.Next.Next.Next.Next.Next = NewLinkNode(13)
	l1.show()
	r2 := linkToBST(l1)
	showBST(r2)
}
