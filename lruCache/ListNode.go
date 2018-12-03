package main

type MyQueue interface {
	push(interface{})
	pop() interface{}
	peek() interface{}
	empty() bool
}

type ListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
	size   int
}

type ListNode struct {
	Val  interface{}
	Next *ListNode
	Prev *ListNode
}

func NewListNode(v interface{}) *ListNode {
	n := new(ListNode)
	n.Val = v
	return n
}

//附加的链表之前
func addHeadNode(n *ListNode, v interface{}) *ListNode {
	nv := NewListNode(v)
	if n == nil {
		return nv
	}
	nv.Next = n
	n.Prev = nv
	return nv
}

func NewListQueue(size int) *ListQueue {
	q := new(ListQueue)
	q.size = size
	q.length = 0
	return q
}

func (q *ListQueue) push(v interface{}) {
	q.length++
	if q.head == nil {
		q.head = addHeadNode(q.head, v)
		q.tail = q.head
		return
	}
	q.head = addHeadNode(q.head, v)
}

func (q *ListQueue) peek() interface{} {
	if q.size == 0 {
		return nil
	}
	return q.tail.Val
}

//推出最后元素
func (q *ListNode) pop() interface{} {
	if q.empty() {
		return nil
	}
	n := q.tail
	if q.head == n {
		q.head = nil
		q.size = 0
		q.tail = nil
	} else {
		q.tail = q.tail.Prev
		q.size = q.size - 1
	}

	return n.Val
}

func (q *ListNode) empty() bool {
	if q.size == 0 {
		return true
	}
	return false
}
