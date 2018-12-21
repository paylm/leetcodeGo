package main

import "fmt"

type BinomialHeap struct {
	Key       int
	degree    int
	Parent    *BinomialHeap
	Next      *BinomialHeap
	LeftChild *BinomialHeap
}

func NewBinomialHeap(k int) *BinomialHeap {
	bh := new(BinomialHeap)
	bh.Key = k
	bh.degree = 0
	return bh
}

func lessBheap(h1, h2 *BinomialHeap) bool {
	if h1.Key < h2.Key {
		return true
	}
	return false
}

/**
h1 to be h2 father
*/
func BinLink(h1, h2 *BinomialHeap) {

	//h1.Next = h2.Next
	h2.Parent = h1
	h2.Next = h1.LeftChild
	h1.LeftChild = h2
	h1.degree = h2.degree + 1
}

func unitonUintBinHeap(h1, h2 *BinomialHeap) *BinomialHeap {
	if h1.Key < h2.Key {
		BinLink(h1, h2)
		return h1
	} else {
		BinLink(h2, h1)
		return h2
	}
}

/**
union two BinomialHeap to new one
*/
func unionBinomialHeap(h1, h2 *BinomialHeap) *BinomialHeap {
	var pre, x, x_next *BinomialHeap
	head := mergeBinomialHeap(h1, h2)
	if head == nil {
		return nil
	}
	x = head
	x_next = x.Next
	for {
		if x_next == nil {
			break
		}
		//case 1
		if x.degree != x_next.degree || ((x.degree == x_next.degree) && (x_next.Next != nil && x_next.degree == x_next.Next.degree)) {
			pre = x
			x = x_next
		} else if lessBheap(x, x_next) {
			//x < x_next, d't handle pre potion
			x.Next = x_next.Next
			BinLink(x, x_next)

		} else {
			//x > x_next
			if pre == nil {
				head = x_next
			} else {
				pre.Next = x_next
			}
			BinLink(x_next, x)
			//pre.Next = x_next
			x = x_next
		}

		x_next = x_next.Next
	}
	return head
}

//将h1, h2中的根表合并成一个按度数递增的链表，返回合并后的根节点
func mergeBinomialHeap(h1 *BinomialHeap, h2 *BinomialHeap) *BinomialHeap {
	var pos, head *BinomialHeap
	for {
		if h1 == nil || h2 == nil {
			break
		}

		if h1.degree < h2.degree {
			if pos == nil {
				pos = h1
				head = pos
			} else {
				pos.Next = h1
				pos = pos.Next
			}
			h1 = h1.Next
		} else {
			if pos == nil {
				pos = h2
				head = pos
			} else {
				pos.Next = h2
				pos = pos.Next
			}
			h2 = h2.Next
		}
	}
	if h1 == nil {
		pos.Next = h2
	} else {
		pos.Next = h1
	}
	return head
}

func traverseBheap(h *BinomialHeap, parent int) {
	if h == nil {
		return
	}
	p := h
	for {
		if p == nil {
			break
		}
		fmt.Printf("%v (parent :%d ) -> ", p, parent)
		if p.LeftChild != nil {
			traverseBheap(p.LeftChild, p.Key)
		}
		fmt.Println()
		p = p.Next
	}
}
