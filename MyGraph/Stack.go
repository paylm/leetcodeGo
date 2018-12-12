package main

import "errors"

type Stack interface {
	empty() bool
	Pop() (error, int)
	Push(int)
	Peek() (error, int)
}

type MyStack struct {
	head *Node
	size int
}

type Node struct {
	Val  int
	Next *Node
}

func NewNode(v int) *Node {
	n := new(Node)
	n.Val = v
	return n
}

func NewMyStack() *MyStack {
	s := new(MyStack)
	s.size = 0
	return s
}

func (s *MyStack) empty() bool {
	if s.head == nil {
		return true
	}
	return false
}

func (s *MyStack) Pop() (error, int) {
	if s.head == nil {
		return errors.New("Stack is empty"), 0
	}

	pVal := s.head.Val
	s.head = s.head.Next
	s.size--
	return nil, pVal
}

func (s *MyStack) Peek() (error, int) {

	if s.head == nil {
		return errors.New("Stack is empty"), 0
	}

	return nil, s.head.Val
}

func (s *MyStack) Push(v int) {
	n := NewNode(v)
	n.Next = s.head
	s.head = n
	s.size++
}
