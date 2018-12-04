package main

type Queue interface {
	peek() interface{}
	pop() interface{}
	put(interface{})
	empty() bool
}

//用sline 实现简单Queue
type MyQueue struct {
	data []interface{}
}

func NewMyQueue() *MyQueue {
	q := new(MyQueue)
	return q
}

func (q *MyQueue) empty() bool {
	if len(q.data) == 0 {
		return true
	}
	return false
}

func (q *MyQueue) peek() interface{} {
	if len(q.data) == 0 {
		return nil
	}

	return q.data[0]
}

func (q *MyQueue) pop() interface{} {

	if len(q.data) == 0 {
		return nil
	}
	v := q.data[0]
	q.data = append(q.data[1:])
	return v
}

func (q *MyQueue) put(v interface{}) {
	q.data = append(q.data, v)
}
