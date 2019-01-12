package main

import (
	"fmt"
	"math/rand"
)

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

//合并链表,b合并到a
func merge(a, b *LinkNode) {
	ap, bp := a, b
	fmt.Println("---merge---")
	ap.PrintLinkNode()
	bp.PrintLinkNode()
	for {
		if ap == nil {
			break
		}
		if bp == nil {
			break
		}
		atemp := ap
		btemp := bp
		ap = ap.Next
		bp = bp.Next
		atemp.Next = btemp
		if ap != nil {
			btemp.Next = ap
		}
	}
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
	//计算截截一半链表
	slowP, fastP, spliP := head, head, head
	for {
		if fastP == nil {
			break
		}
		if fastP.Next == nil {
			break
		}
		spliP = slowP
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}
	spliP.Next = nil //截断链表
	slowP.PrintLinkNode()
	//此时慢指针刚走到一半
	pre, current := slowP, slowP.Next
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

	merge(head, pre) //合并链表
}

//移除重复元素
// 前提列表已排序
func removeDuplicates(head *LinkNode) {
	if head == nil {
		return
	}
	current, pre := head, head
	for {
		if current == nil {
			break
		}

		if current.Val == pre.Val {
			pre.Next = current.Next
		} else {
			pre = current
		}
		current = current.Next
	}
}

//移除重复元素(递归版)
// 前提列表已排序
func deleteDuplicates(head *LinkNode) {
	if head == nil {
		return
	}
	current := head.Next
	if current != nil && head.Val == current.Val {
		head.Next = current.Next
		//fmt.Println("skip->", current.Val)
		deleteDuplicates(head)
	} else {
		deleteDuplicates(current)
	}
}

//移除重复元素
// 链表未排序(hash 方法 空间复杂度O(n))
func removeDuplicatesUnSort(head *LinkNode) {
	if head == nil {
		return
	}
	lMaps := make(map[int]int)
	current, pre := head, head
	for {
		if current == nil {
			break
		}
		if _, ok := lMaps[current.Val]; ok {
			pre.Next = current.Next
		} else {
			pre = current
			lMaps[current.Val] = 1
		}
		current = current.Next
	}
}

/**
交换元素
1) x and y may or may not be adjacent.
2) Either x or y may be a head node.
3) Either x or y may be last node.
4) x and/or y may not be present in linked list.
**/
func swapNodes(head *LinkNode, x int, y int) {
	if head == nil {
		return
	}
	if x == y {
		return
	}
	current := head
	//found x and y location
	preX, preY := head, head
	var pX *LinkNode
	var pY *LinkNode
	for {
		if current == nil {
			break
		}

		if current.Val == x {
			pX = current
			//break
		}
		if current.Val == y {
			pY = current
			//break
		}

		if pX == nil {
			preX = current
		}
		if pY == nil {
			preY = current
		}
		current = current.Next
	}

	fmt.Printf("preX:%d,preY:%d\n", preX.Val, preY.Val)
	if pX == nil || pY == nil {
		return
	}

	if head == pX { //? 待处理
		temp := pY.Next
		head = pY
		pY.Next = pX.Next
		preY.Next = pX
		pX.Next = temp
		return
	}

	//x或y 都不在首部
	YNext := pY.Next
	preX.Next = pY
	pY.Next = pX.Next
	preY.Next = pX
	pX.Next = YNext

}

/**
* 把链表最后一个节点移至最前
 */
func moveToFront(head *LinkNode) *LinkNode {
	if head == nil {
		return head
	}
	last, secLast := head, head
	for {
		if last.Next == nil {
			break
		}
		secLast = last
		last = last.Next
	}

	secLast.Next = nil
	last.Next = head
	return last
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

	l6 := NewLinkNode(61)
	l6.Add(NewLinkNode(62))
	l6.Add(NewLinkNode(63))
	l6.Add(NewLinkNode(64))
	l6.Add(NewLinkNode(65))
	l7 := NewLinkNode(71)
	l7.Add(NewLinkNode(72))
	l7.Add(NewLinkNode(73))
	l7.Add(NewLinkNode(74))
	l7.Add(NewLinkNode(75))
	merge(l6, l7)
	l6.PrintLinkNode()
	fmt.Println("---------reorderList--------")
	l8 := NewLinkNode(81)
	l8.Add(NewLinkNode(82))
	l8.Add(NewLinkNode(83))
	l8.Add(NewLinkNode(84))
	l8.Add(NewLinkNode(85))
	l8.Add(NewLinkNode(86))
	reorderList(l8)

	test1 := NewLinkNode(100)
	ts := []int{101, 102, 103, 104, 105, 106, 107, 108, 109, 110}
	for _, v := range ts {
		test1.Add(NewLinkNode(v))
	}
	reorderList(test1)
	test1.PrintLinkNode()

	fmt.Println("--------  del Duble val --------")
	l9 := NewLinkNode(90)

	l9.Add(NewLinkNode(91))
	l9.Add(NewLinkNode(91))
	l9.Add(NewLinkNode(90))
	l9.Add(NewLinkNode(90))
	l9.Add(NewLinkNode(92))
	l9.Add(NewLinkNode(93))
	l9.PrintLinkNode()
	//removeDuplicates(l9)
	//deleteDuplicates(l9)
	removeDuplicatesUnSort(l9)
	l9.PrintLinkNode()
	fmt.Println(" ----- swap -----")
	fmt.Println(l9)
	//swapNodes(l9, 90, 92)
	l10 := moveToFront(l9)
	l10.PrintLinkNode()
	//fmt.Println(l9)
	fmt.Println("skip list")
	sk := NewSkiplist(&RandomKey{})
	sk.Search(1)

	fmt.Println(sk)
	for i := 0; i < 50; i++ {
		k := rand.Intn(100)
		fmt.Printf("insert %d ,res :%v\n", k, sk.Insert(k))
	}
	fmt.Println("search for 6:", sk.Search(6))
	fmt.Println("search for 18:", sk.Search(18))
	fmt.Println("search for 25:", sk.Search(25))
	//fmt.Println(sk)
	fmt.Println("col type show skiplist")
	sk.showCol()

	fmt.Println("del element 25")
	sk.Del(25)
	sk.show()
	fmt.Println("del element 27")
	sk.Del(27)
	sk.show()
	sk.showCol()

	sk2 := NewSkiplist(&RandomKey{})
	for _, v := range []int{1, 1, 1, 1, 1, 5, 8} {
		sk2.Insert(v)
	}
	fmt.Println("col type show skiplist")
	sk2.showCol()
}
