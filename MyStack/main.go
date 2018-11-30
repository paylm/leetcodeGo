package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type Stacker interface {
	pop() interface{}
	push(v int)
	show()
	empty()
}

//链表实现stack
type Stack struct {
	head     *Node
	lenght   int
	listsize int
}

func NewNode(v int) *Node {
	n := new(Node)
	n.Val = v
	return n
}

func addHeadNode(head *Node, v int) *Node {
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

func (s *Stack) push(v int) {
	newHead := addHeadNode(s.head, v)
	s.head = newHead
	if s.lenght == s.listsize {
		s.listsize++
	}
	s.lenght++
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

func main() {
	fmt.Println("vim-go")
	var s1 Stacker
	s1 = NewStack()
	s1.push(1)
	s1.push(3)
	s1.push(5)
	s1.push(6)

	fmt.Println("------stack info -----")
	s1.show()
	fmt.Println("---- out stack -----")
	fmt.Println(s1.pop())
	fmt.Println(s1.pop())
	fmt.Println(s1.pop())

	fmt.Println("------stack info -----")
	s1.show()
	s1.push(8)
	s1.push(5)
	s1.show()
	s1.empty()
	s1.show()
}
