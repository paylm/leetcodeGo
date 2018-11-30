package main

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

type Stacker interface {
	pop() interface{}
	push(v interface{})
	getDep() int
	show()
	empty()
}

//链表实现stack
type Stack struct {
	head     *Node
	lenght   int
	listsize int
}

func NewNode(v interface{}) *Node {
	n := new(Node)
	n.Val = v
	return n
}

func addHeadNode(head *Node, v interface{}) *Node {
	if head == nil {
		return NewNode(v)
	}

	newHead := NewNode(v)
	newHead.Next = head
	return newHead
}

func removeHeadNode(head *Node) (h *Node, n *Node) {
	if head == nil {
		return nil, nil
	}
	p := head.Next
	head.Next = nil //截断关联
	return p, head
}

func traverse(head *Node) {
	current := head
	for {
		if current == nil {
			break
		}
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func NewStack() *Stack {
	s := new(Stack)
	s.lenght = 0
	s.listsize = 0
	return s
}

func (s *Stack) pop() interface{} {
	newHead, n := removeHeadNode(s.head)
	if n == nil {
		return nil
	}
	s.head = newHead
	s.lenght--

	return n.Val
}

func (s *Stack) push(v interface{}) {
	newHead := addHeadNode(s.head, v)
	s.head = newHead
	if s.lenght == s.listsize {
		s.listsize++
	}
	s.lenght++
}

func (s *Stack) getDep() int {
	return s.lenght
}

func (s *Stack) show() {
	fmt.Printf("stack len:%d , listsize:%d \n", s.lenght, s.listsize)
	traverse(s.head)
}

func (s *Stack) empty() {
	s.lenght = 0
	s.listsize = 0
	s.head = nil
}
