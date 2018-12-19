package main

import (
	"errors"
	"fmt"
)

type PriorityQueue struct {
	data      []int
	heap_size int
	max_size  int
}

func parentId(i int) int {
	return (i - 1) / 2
}

func leftId(i int) int {
	return 2*i + 1
}

func rightId(i int) int {
	return 2*i + 2
}

func NewPriorityQueue(maxsize int) *PriorityQueue {
	pq := new(PriorityQueue)
	pq.max_size = maxsize
	pq.heap_size = 0
	pq.data = make([]int, pq.max_size)
	return pq
}

func swap(a *int, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

func (pq *PriorityQueue) shiftUp(x int) {
	if x == 0 {
		return
	}

	pid := parentId(x)
	if pq.data[pid] < pq.data[x] {
		swap(&(pq.data[pid]), &(pq.data[x]))
		pq.shiftUp(pid)
	}
}

func (pq *PriorityQueue) shiftDown(x int) {
	left := leftId(x)
	right := rightId(x)
	i := x
	if left < pq.heap_size && pq.data[x] < pq.data[left] {
		i = left
	}

	if right < pq.heap_size && pq.data[i] < pq.data[right] {
		i = right
	}

	if i != x {
		swap(&(pq.data[i]), &(pq.data[x]))
		pq.shiftDown(i)
	}
}

func (pq *PriorityQueue) pop() (int, error) {
	if pq.heap_size < 0 {
		return -1, errors.New("queue is empty")
	}
	x := pq.data[0]
	if pq.heap_size > 1 {
		swap(&(pq.data[0]), &(pq.data[pq.heap_size-1]))
		pq.data[pq.heap_size-1] = 0
		pq.heap_size--
		pq.shiftDown(0)
	} else {
		pq.data[pq.heap_size-1] = 0
		pq.heap_size--
	}

	return x, nil
}

func (pq *PriorityQueue) push(v int) error {
	if pq.heap_size >= pq.max_size {
		return errors.New("queue is full")
	}
	pq.data[pq.heap_size] = v
	pq.shiftUp(pq.heap_size)
	pq.heap_size++
	return nil
}

func (pq *PriorityQueue) empty() bool {
	if pq.heap_size == 0 {
		return true
	}
	return false
}

func (pq *PriorityQueue) show() {
	fmt.Println(pq)
}
