package main

import (
	"fmt"
)

//BS HEAP TEST
func testBSHeap() {
	fmt.Println("test")
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a, b)

	mp := NewMinHeap(10)
	mp.Insert(14)
	mp.Insert(5)
	mp.Insert(2)
	mp.Insert(8)
	mp.Insert(1)
	mp.Insert(9)
	mp.Insert(10)
	mp.Insert(12)
	mp.Insert(66)
	mp.Insert(34)
	mp.GetMin()
	fmt.Println(mp)

	//swap(&(mp.harr[0]),&(mp.harr[1]))
	//fmt.Println(mp)

	mp.Remove(4)
	fmt.Println(mp)
	root, _ := mp.PopRoot()
	fmt.Println("root:", root)
	fmt.Println(mp)
	fmt.Println(mp.PopRoot())
	fmt.Println(mp)
	for {
		if mp.isEmpty() {
			break
		}
		r, _ := mp.PopRoot()
		fmt.Println("root Pop:", r)
		fmt.Println(mp)
	}
	fmt.Println(CheckHeap([]int{1, 2, 13, 4, 5, 6, 7}, 0))
	fmt.Println(CheckHeap([]int{2, 12, 5, 14, 34, 9, 10, 66}, 0))
	fmt.Println("sort heap test")

	arr := []int{5, 8, 3, 1, 11, 9, 16}
	mp1 := NewMinHeap(len(arr))
	mp1.heap_size = len(arr)
	mp1.harr = arr
	fmt.Println(mp1)
	fmt.Println(CheckHeap(mp1.harr, 0))
	mp1.sortHeap()
	fmt.Println(mp1)
	fmt.Println(CheckHeap(mp1.harr, 0))
}

func testBmHeap() {

	bh := NewBinomialHeap()
	bh.Push(1)
	bh.Push(12)
	bh.Push(5)
	bh.Push(8)
	bh.Push(2)

	traverseBheap(bh.head, 0)

	fmt.Println("pop the mix")
	fmt.Println(bh.Pop())

	traverseBheap(bh.head, 0)

	fmt.Println("test for peek")
	fmt.Println(bh.Peek())
	fmt.Println(bh.Peek())
	fmt.Println(bh.Peek())

	for {
		if bh.empty() {
			break
		}
		fmt.Println("pop min => ")
		fmt.Println(bh.Pop())
	}
}

func main() {
	testBmHeap()
}
