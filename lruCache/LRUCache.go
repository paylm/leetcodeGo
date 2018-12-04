package main

import "fmt"

type LRUCache interface {
	put(k string, v string)
	get(k string) string
	show()
}

type LRU struct {
	data   map[string]*ListNode
	head   *ListNode
	tail   *ListNode
	length int //当前链表长度
	size   int
}

func NewLRU(cap int) *LRU {
	lru := new(LRU)
	lru.data = make(map[string]*ListNode)
	lru.size = cap
	return lru
}

func (lru *LRU) put(k string, val string) {
	if lru == nil {
		return
	}

	n, ok := lru.data[k]
	fmt.Printf("put %s\n", k)
	if !ok {
		//如果超队列，置空超队的元素
		if lru.size < lru.length {
			//fmt.Println("to many data , lru clean something")
			lru.expire((lru.length - lru.size) + 1)
		}
		n := NewListNode(k, val)
		lru.data[k] = n
		lru.moveToHead(n)
		lru.length++
	} else {
		//fmt.Println("found exist val")
		n.Val = val
		lru.moveToHead(n)
	}
}

func (lru *LRU) get(k string) (bool, string) {

	n, ok := lru.data[k]
	if !ok {
		return false, "k not exist"
	}
	lru.moveToHead(n)
	return true, n.Val.(string)
}

//remove last n node
func (lru *LRU) expire(n int) {
	if n == 0 {
		return
	}
	current := lru.tail
	i := 0
	for {
		if i == n {
			current.Next.Prev = nil
			current.Next = nil
			lru.tail = current
			lru.tail.Next = nil
			//delete(lru.data, current)
			return
		} else {
			delete(lru.data, current.Key)
		}
		fmt.Printf("expire key: %v\n", current)
		current = current.Prev
		lru.length-- //减小长度纪录
		i++
	}
}

//把某个节点移到头部
func (lru *LRU) moveToHead(n *ListNode) {

	PrevN, PostN := n.Prev, n.Next
	//tempHead := lru.head
	defer func() {
		//fmt.Printf("add %s done\n", n.Key)
		if err := recover(); err != nil {
			fmt.Printf("painc error  n:%v,prevN:%v,postN:%v\n", n, PrevN, PostN)
		}
	}()
	if lru.length == 0 || lru.head == nil {
		lru.head = n
		lru.tail = n
		lru.tail.Next = nil
		return
	}

	if n == lru.head {
		//当前已是head
		return
	}

	if PrevN == nil && PostN == nil {
		//此时为新节点
		n.Next = lru.head
		lru.head.Prev = n
		lru.head = n
		return
	}
	//位于第二位
	if PrevN == lru.head {
		lru.head.Next = PostN
		if PostN != nil {
			PostN.Prev = lru.head
		}
		lru.head.Prev = n
		n.Next = lru.head
		lru.head = n
		return
	}

	//重置头部
	n.Next = lru.head
	lru.head.Prev = n
	lru.head = n

	if PostN != nil {
		//n 不为尾时
		PrevN.Next = PostN
		PostN.Prev = PrevN
	} else {
		//重置
		lru.tail = PrevN
		lru.tail.Next = nil
	}
}

func (lru *LRU) pop() *ListNode {
	n := lru.tail
	if n == nil {
		return n
	}
	if lru.head == lru.tail {
		lru.head = nil
		lru.tail = nil

	} else {
		lru.tail = lru.tail.Prev
		lru.tail.Next = nil
	}
	lru.length--
	return n
}

func (lru *LRU) show() {
	fmt.Println("----- lru info -------")
	fmt.Println(lru)
	fmt.Println("head:", lru.head)
	fmt.Println("tail:", lru.tail)
	traverse(lru.head)
}
