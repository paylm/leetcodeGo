package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	//if (head == nil) {
	//   return head
	//}
	var PreM, PostN *ListNode //替换点前后位
	var LM, LN *ListNode      //返转位起始位
	i := 1
	current, Pre := head.Next, head
	for {
		if current == nil {
			break
		}

		if i == m {
			PreM = Pre
			PreM.Next = nil
			LM = current
		}
		if i == n+1 {
			PostN = current
			LN = Pre
			LN.Next = nil
			break
		}
		i++
		Pre = current
		current = current.Next
	}

	if LM == nil || LN == nil { //位置不存在
		return head
	}

	if m == 1 {
		return head
	}
	//reverse
	prev, pc := LM, LM.Next
	for {
		if pc == LN {
			break
		}
		temp := pc
		pc = pc.Next
		temp.Next = prev
		prev = temp
	}

	PreM.Next = pc
	LM.Next = PostN

	return head
}
