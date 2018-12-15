package main

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func NewLinkNode(v int) *LinkNode {
	n := new(LinkNode)
	n.Val = v
	return n
}

func countLinkNode(l *LinkNode) int {
	n := 0
	c := l
	for {
		if c == nil {
			return n
		}
		n++
		c = c.Next
	}
}

func (l *LinkNode) show() {
	c := l
	for {
		if c == nil {
			break
		}
		fmt.Printf("%d->", c.Val)
		c = c.Next
	}
	fmt.Println()
}
