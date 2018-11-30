package main

import "fmt"

//双向链表
type DNode struct {
	Val  interface{}
	Next *DNode
	Prev *DNode
}

func NewDNode(v interface{}) *DNode {
	d := new(DNode)
	d.Val = v
	return d
}

func addNextNode(d *DNode, v interface{}) *DNode {
	newNode := NewDNode(v)
	if d == nil {
		return newNode
	}

	d.Next = newNode
	newNode.Prev = d
	return newNode
}

func addPrevNode(d *DNode, v interface{}) *DNode {
	nNode := NewDNode(v)
	if d == nil {
		return nNode
	}

	d.Prev = nNode
	nNode.Next = d
	return nNode
}

func removeLastNode(last *DNode) (l *DNode, del *DNode) {

	if last == nil {
		return nil, nil
	}

	preLast := last.Prev
	if preLast != nil {
		preLast.Next = nil
	}

	return preLast, last
}

func (d *DNode) traverse() {
	if d == nil {
		return
	}
	current := d
	for {
		if current == nil {
			break
		}
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println()

}
