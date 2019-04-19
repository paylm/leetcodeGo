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

func ArrToLinkList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := NewListNode(arr[0])
	last := head
	for i := 1; i < len(arr); i++ {
		last.Next = NewListNode(arr[i])
		last = last.Next
	}

	return head
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	PreN := head
	c := head.Next
	for {
		if c == nil {
			break
		}
		temp := c
		c = c.Next
		head, PreN = insertSortList(head, temp, PreN)

	}
	return head
}

/**
插入某节点
返回新的head和lastNode
**/
func insertSortList(head *ListNode, insertNode *ListNode, lastNode *ListNode) (*ListNode, *ListNode) {
	fmt.Printf("insertionSortList ===== %v =====\n", insertNode)
	if head == insertNode {
		return head, head
	}

	i := 0
	//defer func() {
	//	fmt.Printf("insertSortList %v i:%d\n", insertNode, i)
	//}()
	c := head
	if head.Val > insertNode.Val {
		lastNode.Next = insertNode.Next
		insertNode.Next = head
		return insertNode, lastNode
	} else {
		for {
			i++
			if c == nil || c == lastNode {
				return head, insertNode
			}

			temp := c //当前节点的上一个
			c = c.Next

			if c != nil && c.Val > insertNode.Val {
				lastNode.Next = insertNode.Next
				temp.Next = insertNode
				insertNode.Next = c
				return head, lastNode
			}

		}
	}
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

	//fmt.Printf("sortList\n")
	//showLinkArr(ls)
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
	i, j, k := start, start, mid
	/**
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("i:%d,j:%d,k:%d\nerr:%v\n", i, j, k, err)
		}
	}()
	**/
	//fmt.Printf("sortMerge ls:%v,start:%d,mid:%d,end:%d\n", ls, start, mid, end)
	//showLinkArr(ls)
	if start >= end {
		return
	}
	for {
		if j >= mid || k > end {
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
	if i >= mid {
		for ; k <= end; k++ {
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

func showList(head *ListNode) {
	c := head
	for {
		if c == nil {
			break
		}
		fmt.Printf("%d->", c.Val)
		c = c.Next
	}
	fmt.Println()
}
