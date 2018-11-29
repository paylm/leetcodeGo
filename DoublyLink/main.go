package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func NewNode(v int) *Node {
	n := new(Node)
	n.Val = v
	return n
}

func (head *Node) addNode(n *Node) {
	if head == nil {
		return
	}
	p, last := head, head
	for {
		if p == nil {
			break
		}
		last = p
		p = p.Next
	}
	last.Next = n
	n.Prev = last
}

func printList(head *Node) {
	if head == nil {
		return
	}
	current, last := head, head
	fmt.Println("Traversal in forward direction")
	for {
		if current == nil {
			break
		}
		fmt.Printf("%d->", current.Val)
		last = current
		current = current.Next
	}
	fmt.Println()
	fmt.Println("Traversal in reverse  direction")
	current = last
	for {
		if current == nil {
			break
		}
		fmt.Printf("%d->", current.Val)
		last = current
		current = current.Prev
	}
	fmt.Println()
}

func (head *Node) delNode(v int) {
	if head == nil {
		return
	}
	current := head
	for {
		if current == nil {
			fmt.Printf("节点[%d]没找到\n", v)
			return
		}
		if current.Val == v {
			break
		}
		//fmt.Printf("%d->", current.Val)
		current = current.Next
	}

	PreDel := current.Prev
	if current.Next == nil {
		PreDel.Next = nil
		return
	}
	PostDel := current.Next
	PreDel.Next = PostDel
	PostDel.Prev = PreDel
}

func reverse(head *Node) *Node {
	if head == nil {
		return nil
	}
	current, last := head.Next, head
	last.Next = nil //置空,防止有环
	for {
		if current == nil {
			break
		}
		fmt.Println("current :", last.Val)
		temp := current
		current = current.Next
		last.Prev = temp
		temp.Next = last
		last = temp

	}
	last.Prev = nil //置空,防止有环
	fmt.Println("current :", last.Val)
	fmt.Println("reverse ok")
	return last
}

func main() {
	l1 := NewNode(1)
	l1.addNode(NewNode(3))
	l1.addNode(NewNode(5))
	l1.addNode(NewNode(6))
	l1.addNode(NewNode(7))
	printList(l1)
	//l1.delNode(5)  //中节点节删
	l1.delNode(7) //尾节点删除
	printList(l1)

	l2 := reverse(l1)
	fmt.Println(l2)
	printList(l2)
}
