package main

import (
	"fmt"
	"errors"
)

type BinomialHeap struct{
	head *BinomialNode
}

type BinomialNode struct {
	Key       int
	degree    int
	Parent    *BinomialNode
	Next      *BinomialNode
	LeftChild *BinomialNode
}

func NewBinomialNode(k int) *BinomialNode {
	bh := new(BinomialNode)
	bh.Key = k
	bh.degree = 0
	return bh
}

func lessBheap(h1, h2 *BinomialNode) bool {
	if h1.Key < h2.Key {
		return true
	}
	return false
}

/**
h1 to be h2 father
*/
func BinLink(h1, h2 *BinomialNode) {

	//h1.Next = h2.Next
	h2.Parent = h1
	h2.Next = h1.LeftChild
	h1.LeftChild = h2
	h1.degree = h2.degree + 1
}

func unitonUintBinHeap(h1, h2 *BinomialNode) *BinomialNode {
	if h1.Key < h2.Key {
		BinLink(h1, h2)
		return h1
	} else {
		BinLink(h2, h1)
		return h2
	}
}

/**
union two BinomialNode to new one
*/
func unionBinomialNode(h1, h2 *BinomialNode) *BinomialNode {
	//fmt.Println("unionBinomialNode => h1:",h1)
	//fmt.Println("unionBinomialNode => h2:",h2)
	var pre, x, x_next *BinomialNode
	head := mergeBinomialNode(h1, h2)
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
func mergeBinomialNode(h1 *BinomialNode, h2 *BinomialNode) *BinomialNode {
	var pos, head *BinomialNode
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

//找到堆中最小的值
func (bh *BinomialNode)FindMinNode() *BinomialNode{
	if bh == nil {
		return nil
	}
	minHeap,c := bh,bh
	for{
		if c == nil{
			break
		}
		if lessBheap(c,minHeap){
			minHeap = c
		}
		c = c.Next
	}
	return minHeap
}

func NewBinomialHeap() *BinomialHeap{
	bh := new(BinomialHeap)
	return bh
}

func (bh *BinomialHeap)Pop() (int,error){
	if bh == nil{
		return -1,errors.New("BinomialNode is empty")
	}
	//find the min Node
	preHeap,minHeap,c := bh.head,bh.head,bh.head
	for{
		if c == nil || c.Next == nil{
			break
		}

		if lessBheap(c.Next,minHeap){
			preHeap = c
			minHeap = c.Next
		}
		c = c.Next
	}

	deletQue := minHeap.LeftChild
	dc := deletQue.Next
	deletQue.Next = nil
	//update parent and reverse ?
	for {
		if dc == nil {
			break
		}
		dc.Parent = nil
		temp := dc
		dc = dc.Next
		temp.Next = deletQue
		deletQue = temp
	}
	//fmt.Println("update parent and reverse")
	//traverseBheap(deletQue,0)

	if preHeap != minHeap{
		preHeap.Next = minHeap.Next //截断删除点
	}else {
		//fmt.Println("bh udpate",bh)
		bh.head = minHeap.Next
	}
	//fmt.Println("bh :",bh)
	if bh == nil{
		bh.head = deletQue
		//fmt.Println("delQueu => bh",bh,deletQue)
	}else{
		//union two new heap
		unionBinomialNode(bh.head,deletQue)
	}

	return minHeap.Key,nil
}

func (bh *BinomialHeap)Push(v int){
	bn := NewBinomialNode(v)
	if bh.head == nil {
		bh.head = bn
		return
	}
	bh.head = unionBinomialNode(bh.head,bn)
}

func traverseBheap(h *BinomialNode, parent int) {
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
