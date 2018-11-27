package main

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func NewLinkNode(v int) *LinkNode {
	n := new(LinkNode)
	n.Val = v
	return n
}

func (n *LinkNode) Add(aNode *LinkNode) {

	//	defer func() {
	//		fmt.Printf("add Node %d done \n", aNode.Val)
	//	}()
	if n == nil {
		n = aNode
		return
	}

	head := n
	for {
		if head.Next == nil {
			head.Next = aNode
			return
		}
		head = head.Next
	}
}

func (n *LinkNode) PrintLinkNode() {
	if n == nil {
		fmt.Println("n is nil")
		return
	}
	p := n
	for {
		fmt.Printf("%d->", p.Val)
		if p.Next == nil {
			break
		}
		p = p.Next
	}

	fmt.Println()
}

//输出中间节点(快慢指针法)
func (n *LinkNode) PrintMidle() {
	fastP, slowP := n, n
	if n == nil {
		return
	}

	for {
		if fastP == nil {
			fmt.Printf("Midle node is :%d\n", slowP.Val)
			return
		}
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}
}

/***
* 引用计数器方式
**/
func (n *LinkNode) PrintMidleV2() {
	step := 0
	mid, p := n, n
	for {
		if p.Next == nil {
			break
		}
		if step%2 == 0 { //每2进1
			mid = mid.Next
		}

		step++
		p = p.Next
	}

	fmt.Printf("MidleV2 Node is :%d\n", mid.Val)
}

func (n *LinkNode) Del(k int) bool {
	fmt.Printf("del node:%d", k)
	if n == nil {
		return false
	}

	if n.Val == k {
		n = nil
		return true
	}

	preN, delN := n, n.Next

	for {

		if delN == nil {
			return false
		}

		if delN.Val == k {
			preN.Next = delN.Next
			delN = nil
			return true
		}
		preN = delN
		delN = delN.Next
	}
	return false
}

func (n *LinkNode) Insert(aNode *LinkNode, v int) {
	if n == nil {
		return
	}
	preN, postN := n, n.Next
	for {
		if postN == nil {
			break
		}
		if postN.Val == v {
			preN.Next = aNode
			aNode.Next = postN
			fmt.Printf("Insert %d ok !!\n", aNode.Val)
			return
		}
		preN = postN
		postN = postN.Next
	}
	fmt.Printf("Insert %d fail , location Node not found !!\n", aNode.Val)
}

/**
 返回是否存在，并返回节点信息
**/
func (n *LinkNode) FindNode(v int) (bool, *LinkNode) {

	np := n
	for {
		if np == nil {
			return false, nil
		}
		if np.Val == v {
			return true, np
		}
		np = np.Next
	}
}

/**
* 检查链表是否有环
 */
func (n *LinkNode) checkLoop() bool {
	fastP, slowP := n, n
	step := 0
	for {
		if slowP == nil || fastP == nil {
			break
		}
		if step%4 == 3 {
			slowP = slowP.Next
		}
		step++
		fastP = fastP.Next
		if slowP == fastP {
			fmt.Printf("此链表有环，环为:%d\n", slowP.Val)
			return true
		}
	}
	fmt.Println("此链表无环")
	return false
}

/**
* 检查链表是否有环
 */
func (n *LinkNode) countLoopLenght() int {

	fastP, slowP := n, n
	step, slwT := 0, 0
	for {
		if slowP == nil || fastP == nil {
			break
		}
		if step%4 == 3 {
			slwT++
			slowP = slowP.Next
		}
		step++
		fastP = fastP.Next
		if slowP == fastP {
			fmt.Printf("此链表有环，环为:%d,环长度为:%d\n", slowP.Val, step-slwT)
			return step - slwT
		}
	}
	fmt.Println("此链表无环")
	return 0
}
func (head *LinkNode) ReverseList() *LinkNode {
	if head == nil {
		return nil
	}
	pre, current := head, head.Next
	pre.Next = nil
	for {
		if current == nil {
			fmt.Println("reverse done")
			break
		}
		fmt.Printf("%d->", current.Val)
		temp := current
		current = current.Next
		temp.Next = pre
		pre = temp
	}
	return pre
}

/***

给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.


思路：找到中结点，反转此部分，再拼接
***/
func reorderList(head *LinkNode) {

	if head == nil {
		return
	}
	pre, current := head, head.Next
	pre.Next = nil
	for {
		if current == nil {
			fmt.Println("reverse done")
			break
		}
		//fmt.Printf("%d->", current.Val)
		temp := current
		current = current.Next
		temp.Next = pre
		pre = temp
	}
	pre.PrintLinkNode()
	// 此时pre 是反向链表
	hn, pn := head, pre
	for {

		if hn == pn {
			hn.Next = nil

			break
		}
		htemp := hn
		ptemp := pn
		hn = hn.Next
		pn = pn.Next

		htemp.Next = ptemp
		ptemp.Next = hn
	}
}

func main() {
	fmt.Println("vim-go")
	arr := [...]int{2, 3, 5, 7, 8, 9, 10}
	var head = NewLinkNode(1)

	for _, v := range arr {
		n := NewLinkNode(v)
		head.Add(n)
	}
	fmt.Println("--------链表如下------")
	head.PrintLinkNode()
	fmt.Println("-----中节点------")
	head.PrintMidle()
	head.PrintMidleV2()

	l2 := NewLinkNode(5)
	l2.PrintMidleV2()

	l2.Add(NewLinkNode(7))
	l2.Add(NewLinkNode(9))
	l2.PrintLinkNode()
	l2.PrintMidleV2()

	fmt.Println("------删除节点-----")
	fmt.Println(head.Del(5))
	head.PrintLinkNode()
	head.Insert(NewLinkNode(99), 9)
	head.PrintLinkNode()
	head.Insert(NewLinkNode(47), 11)

	fmt.Println("----- find one Node ------")
	fmt.Println(head.FindNode(99))
	fmt.Println(head.FindNode(98))
	fmt.Println("------检测链表是否环------")
	head.checkLoop()
	head.PrintLinkNode()
	//加环
	head.Next.Next.Next.Next = head
	head.checkLoop()
	//head.PrintLinkNode()
	head.countLoopLenght()

	l2.Add(NewLinkNode(66))
	l2.countLoopLenght()

	l3 := NewLinkNode(3)
	l3.Add(NewLinkNode(4))
	l3.Add(NewLinkNode(5))
	l3.Add(NewLinkNode(6))
	l3.Add(NewLinkNode(7))
	l3.PrintLinkNode()
	l3r := l3.ReverseList()
	l3r.PrintLinkNode()
	l3.PrintLinkNode()

	fmt.Println("----重排链表------")
	//reorderList(l3)
	//l3.PrintLinkNode()
}
