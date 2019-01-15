package main

import "errors"

type CricyleQue struct {
	Size int
	font int
	rear int
	data []int
}

func NewCricyleQue(s int) *CricyleQue {
	cq := new(CricyleQue)
	cq.Size = s
	cq.font = 0
	cq.rear = 0
	cq.data = make([]int, s)
	return cq
}

func (cq *CricyleQue) Push(v int) error {
	if (cq.rear+1)%cq.Size == cq.font {
		return errors.New("CircyleQue is full")
	}
	cq.data[cq.rear] = v
	cq.rear = (cq.rear + 1) % cq.Size
	return nil
}

func (cq *CricyleQue) Pop() (int, error) {
	if cq.empty() {
		return 0, errors.New("Queue is empty")
	}
	v := cq.data[cq.font]
	cq.font = (cq.font + 1) % cq.Size
	return v, nil
}

func (cq *CricyleQue) Font() (int, error) {
	if cq.empty() {
		return 0, errors.New("Queue is empty")
	}
	return cq.data[cq.font], nil
}

func (cq *CricyleQue) Rear() (int, error) {
	if cq.empty() {
		return 0, errors.New("Queue is empty")
	}
	return cq.data[cq.rear], nil
}

func (cq *CricyleQue) empty() bool {
	if cq.font == cq.rear {
		return true
	}
	return false
}
