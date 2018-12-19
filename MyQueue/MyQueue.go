package main

import "fmt"

type MyQueue interface {
	pop() interface{}
	push(v interface{})
	empty() bool //队列是否为空
	show()
}

//通过用两个stack实现Queue
type Queue struct {
	s1     Stacker
	s2     Stacker
	lenght int
	size   int
}

//通过双链表实现Queue
type LQueue struct {
	Head     *DNode
	Last     *DNode
	lenght   int
	listsize int
}

func NewQueue() *Queue {
	q := new(Queue)
	q.s1 = NewStack()
	q.s2 = NewStack()
	q.lenght = 0
	q.size = 0
	return q
}

func (q *Queue) pop() interface{} {
	ds1, ds2 := q.s1.getDep(), q.s2.getDep()
	if ds1 == 0 && ds2 == 0 {
		return nil
	}
	if ds2 == 0 && ds1 == 1 {
		return q.s1.pop()
	}

	if ds2 == 0 {
		for {
			if q.s1.getDep() < 1 {
				break
			}
			q.s2.push(q.s1.pop())
		}
	}

	return q.s2.pop()
}

func (q *Queue) empty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func (q *Queue) push(v interface{}) {
	q.s1.push(v)
}

func (q *Queue) show() {
	fmt.Println(q)
	ds1, ds2 := q.s1.getDep(), q.s2.getDep()
	if ds1 == 0 && ds2 == 0 {
		fmt.Println("queue is null")
		return
	}

	if ds2 == 0 {
		q.s1.show()
		return
	}

	q.s2.show()
	q.s1.show()
}

//DQueue实现
func NewLQueue() *LQueue {
	q := new(LQueue)
	q.lenght = 0
	q.listsize = 0
	return q
}

func (lq *LQueue) pop() interface{} {
	if lq.lenght == 0 {
		return nil
	}
	var popNode *DNode
	lq.Last, popNode = removeLastNode(lq.Last)
	if popNode != nil {
		lq.lenght--
		lq.listsize--

	}
	return popNode.Val
}

func (lq *LQueue) push(v interface{}) {

	if lq == nil {
		fmt.Println("Queue is nil")
		return
	}

	lq.Head = addPrevNode(lq.Head, v)
	if lq.Last == nil {
		lq.Last = lq.Head
	}

	if lq.lenght == lq.listsize {
		lq.listsize++
	}
	lq.lenght++
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
