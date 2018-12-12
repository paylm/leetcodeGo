package main

import (
	"errors"
	"fmt"
	"sync"
)

type MyQueue interface {
	pop() (error, int)
	push(v int)
	peek() (error, int)
	empty() bool //队列是否为空
	show()
}

//通过双链表实现Queue
type LQueue struct {
	Head     *DNode
	Last     *DNode
	lenght   int
	listsize int
	lock     *sync.Mutex
}

//双向链表
type DNode struct {
	Val  int
	Next *DNode
	Prev *DNode
}

func NewDNode(v int) *DNode {
	d := new(DNode)
	d.Val = v
	return d
}

func addNextNode(d *DNode, v int) *DNode {
	newNode := NewDNode(v)
	if d == nil {
		return newNode
	}

	d.Next = newNode
	newNode.Prev = d
	return newNode
}

func addPrevNode(d *DNode, v int) *DNode {
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

//DQueue实现
func NewLQueue() *LQueue {
	q := new(LQueue)
	q.lenght = 0
	q.listsize = 0
	q.lock = new(sync.Mutex)
	return q
}

func (lq *LQueue) pop() (error, int) {
	if lq.lenght == 0 {
		return errors.New("Queue is nil"), -1
	}
	var popNode *DNode
	lq.Last, popNode = removeLastNode(lq.Last)
	if popNode != nil {
		lq.lenght--
		lq.listsize--

	}
	return nil, popNode.Val
}

func (lq *LQueue) push(v int) {

	if lq == nil {
		fmt.Println("Queue is nil")
		return
	}
	lq.lock.Lock()
	lq.Head = addPrevNode(lq.Head, v)
	if lq.Last == nil {
		lq.Last = lq.Head
	}

	if lq.lenght == lq.listsize {
		lq.listsize++
	}
	lq.lenght++
	lq.lock.Unlock()
}

func (lq *LQueue) peek() (error, int) {
	peekNode := lq.Last
	if peekNode == nil {
		return errors.New("queue is empty"), -1
	}
	return nil, peekNode.Val
}

func (lq *LQueue) empty() bool {
	if lq.lenght == 0 {
		return true
	}

	return false
}

func (lq *LQueue) show() {
	fmt.Println(lq)
	lq.Head.traverse()
}
