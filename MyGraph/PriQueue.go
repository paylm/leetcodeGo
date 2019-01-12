package main

import (
	"errors"
	"fmt"
)

type MinHeap struct {
	harr      []int // pointer to array of elements in heap
	capacity  int   // maximum possible size of min heap
	heap_size int   // Current number of elements in min heap
}

func NewMinHeap(capacity int) *MinHeap {
	mp := new(MinHeap)
	mp.capacity = capacity
	mp.heap_size = 0
	mp.harr = make([]int, mp.capacity)
	return mp
}

func parentId(i int) int {
	return (i - 1) / 2
}

func leftId(i int) int {
	return i*2 + 1
}

func rightId(i int) int {
	return i*2 + 2
}

func swap(a *int, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

//上浮某个节点
func (mp *MinHeap) shiftUp(i int) {
	if i == 0 {
		return
	}
	parent := parentId(i)
	if mp.harr[i] < mp.harr[parent] {
		swap(&(mp.harr[parent]), &(mp.harr[i]))
		mp.shiftUp(parent)

	}
}

/**下沉某个节点
 x 与子点节小的比较,下沉到较小的子节点位置
**/
func (mp *MinHeap) shiftDown(i int) {
	left := leftId(i)
	right := rightId(i)
	min := i //ex 为待交换点index

	if left < mp.heap_size && mp.harr[left] < mp.harr[i] {
		min = left
	}
	if right < mp.heap_size && mp.harr[right] < mp.harr[min] {
		min = right
	}
	//fmt.Printf("i:%d,min:%d,left:%d,right:%d\n",i,min,left,right)
	if i != min {
		swap(&(mp.harr[i]), &(mp.harr[min]))
		mp.shiftDown(min)
	}

}
func (mp *MinHeap) Insert(x int) error {
	if mp.heap_size >= mp.capacity {
		return errors.New("head is full")
	}
	mp.harr[mp.heap_size] = x

	//	parent := parentId(mp.heap_size)
	//	//fmt.Printf("p:%d => %d,i:%d = >%d\n",parent,mp.harr[parent],mp.heap_size,mp.harr[mp.heap_size])
	//	if mp.harr[mp.heap_size] < mp.harr[parent] {
	mp.shiftUp(mp.heap_size)
	//	}
	mp.heap_size++
	return nil
}

func (mp *MinHeap) empty() bool {
	if mp.heap_size == 0 {
		return true
	} else {
		return false
	}
}

func (mp *MinHeap) Remove(x int) {
	if x > mp.heap_size {
		return
	}
	if x <= mp.heap_size {
		swap(&(mp.harr[x]), &(mp.harr[mp.heap_size-1]))
		mp.harr[mp.heap_size-1] = 0
		mp.heap_size--
		mp.shiftDown(x)
	}
}

func (mp *MinHeap) pop() (error, int) {
	if mp.heap_size == 0 {
		return errors.New("heap is empty"), -1
	}
	x := mp.harr[0]
	//mp.heap_size--
	mp.Remove(0)
	return nil, x
}

func (mp *MinHeap) peek() (error, int) {
	if mp.heap_size == 0 {
		return errors.New("heap is empty"), -1
	}
	x := mp.harr[0]
	return nil, x
}

func (mp *MinHeap) push(v int) {
	mp.Insert(v)
}

func (mp *MinHeap) show() {
	fmt.Println(mp)
}
