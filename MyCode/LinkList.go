package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	n := new(ListNode)
	n.Val = val
	return n
}

/**
https://leetcode-cn.com/problems/sort-list/
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ls := []*ListNode{}
	cur := head
	for {
		if cur == nil {
			break
		}
		temp := cur
		cur = cur.Next
		temp.Next = nil
		ls = append(ls, temp)
	}
	//fmt.Printf("sortList ls:\n")
	//showLinkArr(ls)
	tls := make([]*ListNode, len(ls))
	mergerSortList(ls, tls, 0, len(ls)-1)

	fmt.Printf("sortList\n")
	showLinkArr(ls)
	for i := 1; i < len(ls); i++ {
		ls[i-1].Next = ls[i]
	}
	//showLinkArr(tls)
	return ls[0]
}

func mergerSortList(ls []*ListNode, tls []*ListNode, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergerSortList(ls, tls, start, mid)
	mergerSortList(ls, tls, mid+1, end)
	sortMerge(ls, tls, start, mid+1, end)
}

func sortMerge(ls []*ListNode, tls []*ListNode, start int, mid int, end int) {
	//fmt.Printf("sortMerge ls:%v,start:%d,mid:%d,end:%d\n", ls, start, mid, end)
	//showLinkArr(ls)
	if start >= end {
		return
	}
	i, j, k := start, start, mid
	for {
		if j > mid || k > end {
			break
		}
		//fmt.Printf("j:%d,ls[j]:%v,k:%d,ls[k]:%v\n", j, ls[j], k, ls[k])
		if ls[j].Val > ls[k].Val {
			tls[i] = ls[k]
			k++
		} else {
			tls[i] = ls[j]
			j++
		}
		i++
	}
	//处理剩下的
	if k > end {
		for ; j < mid; j++ {
			tls[i] = ls[j]
			i++
		}
	}
	if i > mid {
		for ; k < end; k++ {
			tls[i] = ls[k]
			i++
		}
	}
	//fmt.Printf("after sortMerge\n")
	for o := start; o <= end; o++ {
		ls[o] = tls[o]
	}
	//fmt.Println("after sortMerge copy ==>>>")
	//showLinkArr(ls)
}

func showLinkArr(ls []*ListNode) {
	for i := 0; i < len(ls); i++ {
		fmt.Printf("%v\n", ls[i])
	}
}
