package main

import (
	"fmt"
)

//refere:https://www.geeksforgeeks.org/circular-singly-linked-list-insertion/

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func NewLinkNode(v int) *LinkNode {
	n := new(LinkNode)
	n.Val = v
	return n
}

func traverse(last *LinkNode) {
	if last == nil {
		return
	}

	fmt.Printf("%d->", last.Val)
	current := last.Next
	for {
		if current == last {
			break
		}
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func addToEmpty(last *LinkNode, val int) *LinkNode {
	if last != nil {
		return last
	}
	last = NewLinkNode(val)
	last.Next = last
	return last
}

func addBegin(last *LinkNode, val int) *LinkNode {

	if last == nil {
		return addToEmpty(last, val)
	}

	temp := NewLinkNode(val)
	temp.Next = last.Next
	last.Next = temp
	return temp
}

func addEnd(last *LinkNode, v int) {
	if last == nil {
		return
	}
	lastSec, current := last, last.Next
	for {
		if current == last {
			break
		}
		lastSec = current
		current = current.Next
	}
	temp := NewLinkNode(v)
	temp.Next = lastSec.Next
	lastSec.Next = temp
}

func countLink(last *LinkNode) int {

	if last == nil {
		return 0
	}
	count := 1
	p := last.Next
	for {
		if p == last {
			return count
		}
		count++
		p = p.Next
	}
}

func main() {
	var emptyLink *LinkNode
	last := addToEmpty(emptyLink, 3)
	//fmt.Println(last)
	traverse(last)

	l2 := addBegin(last, 5)
	//fmt.Println(l2)
	traverse(l2)
	addEnd(l2, 1)
	traverse(l2)
	fmt.Println("环长度:", countLink(l2))
}
