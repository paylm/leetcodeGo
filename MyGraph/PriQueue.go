package main

import (
	"errors"
	"fmt"
)

type PriQueue interface {
	Push(e *Element) error
	Pop() (error, *Element)
	DecreseKey(e *Element)
	empty() bool
}

type Element struct {
	i      int //index
	Weigth int
	Val    int //节点信息
}

type MinHeap struct {
	harr      []*Element // pointer to array of elements in heap
	capacity  int        // maximum possible size of min heap
	heap_size int        // Current number of elements in min heap
}

func NewMinHeap(capacity int) *MinHeap {
	mp := new(MinHeap)
	mp.capacity = capacity
	mp.heap_size = 0
	mp.harr = make([]*Element, mp.capacity)
	return mp
}

//边
type Edge struct {
	Weigth int
	Src    int
	Dst    int
}

type MaxEdegHeap struct {
	eharr     []*Edge
	capacity  int
	heap_size int
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

//todo ... update i
func Swap(a *Element, b *Element) {
	//fmt.Printf("swap %v %v\n", a, b)
	temp := *b
	tmpi := a.i
	*b = *a
	b.i = temp.i
	*a = temp
	a.i = tmpi
}

//上浮某个节点
func (mp *MinHeap) shiftUp(i int) {
	if i == 0 {
		return
	}
	parent := parentId(i)
	if ltElement(mp.harr[i], mp.harr[parent]) {
		Swap(mp.harr[parent], mp.harr[i])
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

	if left < mp.heap_size && ltElement(mp.harr[left], mp.harr[i]) {
		min = left
	}
	if right < mp.heap_size && ltElement(mp.harr[right], mp.harr[min]) {
		min = right
	}
	//fmt.Printf("i:%d,min:%d,left:%d,right:%d\n",i,min,left,right)
	if i != min {
		Swap(mp.harr[i], mp.harr[min])
		mp.shiftDown(min)
	}

}
func (mp *MinHeap) Push(e *Element) error {
	if mp.heap_size >= mp.capacity {
		return errors.New("head is full")
	}
	e.i = mp.heap_size
	mp.harr[mp.heap_size] = e

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

func (mp *MinHeap) Pop() (error, *Element) {
	if mp.empty() {
		return errors.New("heap is empty"), nil
	}
	x := mp.harr[0]
	//fmt.Printf("pop X:%v\n", x)
	//mp.Remove(0)

	if mp.heap_size > 1 {
		mp.harr[0] = mp.harr[mp.heap_size-1]
		mp.harr[mp.heap_size-1] = nil
		mp.heap_size--
		mp.shiftDown(0)
	} else {
		mp.harr[0] = nil
		mp.heap_size--
	}
	return nil, x
}

func (mp *MinHeap) DecreseKey(e *Element) {
	if e == nil {
		return
	}

	i := e.i

	if mp.harr[i] != e {
		return
	}

	mp.shiftDown(i)
}

func (mp *MinHeap) peek() (error, *Element) {
	if mp.heap_size == 0 {
		return errors.New("heap is empty"), nil
	}
	x := mp.harr[0]
	return nil, x
}

func (mp *MinHeap) show() {
	fmt.Println(mp)
}

func NewElement(w int) *Element {
	e := new(Element)
	e.Weigth = w
	return e
}

/**
  return  e1 > e1
**/
func ltElement(e1, e2 *Element) bool {
	if e1.Weigth <= e2.Weigth {
		return true
	}
	return false
}

func NewEdge(weight int, src int, dst int) *Edge {
	e := new(Edge)
	e.Weigth = weight
	e.Src = src
	e.Dst = dst
	return e
}

func NewMinEdgeHeap(size int) *MaxEdegHeap {
	eheap := new(MaxEdegHeap)
	eheap.eharr = make([]*Edge, size)
	eheap.capacity = size
	eheap.heap_size = 0
	return eheap
}

func gtEdge(e1 *Edge, e2 *Edge) bool {
	if e1.Weigth > e2.Weigth {
		return true
	} else {
		return false
	}
}

func swapEdge(e1 *Edge, e2 *Edge) {
	tmp := *e1
	*e1 = *e2
	*e2 = tmp
}

func (mh *MaxEdegHeap) Push(e *Edge) {
	if mh.heap_size >= mh.capacity {
		fmt.Printf("heap is full")
		return
	}
	mh.eharr[mh.heap_size] = e
	mh.shiftUp(mh.heap_size)
	mh.heap_size++
}

func (mh *MaxEdegHeap) Pop() *Edge {
	if mh.heap_size == 0 {
		return nil
	}

	e := mh.eharr[0]
	mh.eharr[0] = mh.eharr[mh.heap_size-1]
	mh.eharr[mh.heap_size-1] = nil
	mh.heap_size--
	mh.shiftDown(0)
	return e
}

func (mh *MaxEdegHeap) shiftUp(i int) {

	if i == 0 {
		return
	}
	parent := parentId(i)
	if gtEdge(mh.eharr[parent], mh.eharr[i]) {
		swapEdge(mh.eharr[i], mh.eharr[parent])
		mh.shiftUp(parent)
	}
}

func (mh *MaxEdegHeap) shiftDown(i int) {

	left := leftId(i)
	right := rightId(i)
	min := i //ex 为待交换点index

	if left < mh.heap_size && gtEdge(mh.eharr[i], mh.eharr[left]) {
		min = left
	}
	if right < mh.heap_size && gtEdge(mh.eharr[min], mh.eharr[right]) {
		min = right
	}
	//fmt.Printf("i:%d,min:%d,left:%d,right:%d\n",i,min,left,right)
	if i != min {
		swapEdge(mh.eharr[i], mh.eharr[min])
		mh.shiftDown(min)
	}
}

func (mh *MaxEdegHeap) empty() bool {
	if mh.heap_size == 0 {
		return true
	}
	return false
}
