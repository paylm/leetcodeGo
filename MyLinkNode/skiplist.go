package main

import (
	"fmt"
	"math/rand"
)

var MAX_LEVEL = 7

type randomKeyer interface {
	randomNum(k int) int //基于最大值k生成随机数
}

type randomKey struct {
}

type Node struct {
	Val       int
	Next      *Node
	levelNode []*Node //存的是每层指向下一指针的头部
	level     int
}

type Skiplist struct {
	level int
	head  *Node
	tail  *Node
	randomKeyer
}

func NewSkiplist(rk randomKeyer) *Skiplist {
	sk := new(Skiplist)
	sk.level = MAX_LEVEL
	sk.randomKeyer = rk
	h := make([]*Node, sk.level)
	for i := 0; i < sk.level; i++ {
		h[i] = nil
	}
	sk.head = NewNode(-100)
	sk.head.levelNode = h

	return sk
}

func NewNode(val int) *Node {
	n := new(Node)
	n.Val = val
	return n
}

func (rdk *randomKey) randomNum(k int) int {
	l := 1
	for {
		if rand.Intn(2) == 0 {
			break
		}
		l++
	}
	if l < k {
		return l
	}
	return k
}

func (sk *Skiplist) Search(k int) *Node {
	if sk.head.Next == nil {
		return nil
	}
	top := len(sk.head.levelNode) - 1

	c := sk.head

	for i := top; i >= 0; i-- {
		if c.levelNode[i] != nil {
			continue
		} else {
			temp := c.levelNode[i]

			for {
				if temp == nil || temp.Val > k {
					break
				}
				if temp.Val == k {
					fmt.Printf("search k found at :%d\n", i)
					return temp
				}
				temp = temp.Next
			}
		}

	}
	return nil
}

func (sk *Skiplist) Insert(Val int) bool {

	top := len(sk.head.levelNode) - 1
	c := sk.head

	for i := top; i >= 0; i-- {
		if c.levelNode[i] == nil {
			continue
		}

		temp := c.levelNode[i]
		if temp.Val == Val {
			fmt.Printf("already exits %d , insert fail\n", Val)
			return false
		}
		for {
			fmt.Printf("i:%d,Val:%d,temp:%v\n", i, Val, temp)
			if temp == nil {
				break
			}
			if temp.Val > Val {
				break
			}

			if temp.Val < Val && temp.Next == nil {
				fmt.Println("到最后结点 ....")
				c = temp
				break
			}
			if temp.Val < Val && temp.levelNode[i] != nil && Val <= temp.levelNode[i].Val {
				fmt.Printf("找到合适插入位置%d\n", c.Val)
				c = temp
				break
			}
			fmt.Printf("%d %v 找不到符合位置,到下一位找\n", temp.Val, temp)
			temp = temp.levelNode[i]
		}
	}

	fmt.Printf("insert point:%d,after:%d\n", Val, c.Val)
	n := NewNode(Val)
	n.level = sk.randomNum(sk.level)
	n.levelNode = make([]*Node, n.level)
	fmt.Printf("add n:%v\n", n)
	current := sk.head
	for i := 0; i < n.level; i++ {
		if i == 0 {
			temp := c.Next
			n.Next = temp
			c.Next = n
		} else {
			//处理levelNode 指向指针
			if current.levelNode[i-1] == nil {
				current.levelNode[i-1] = n
			}
		}
	}

	sk.show()
	return true
}

func (sk *Skiplist) Del(Val int) bool {
	return true
}

func (sk *Skiplist) show() {
	fmt.Println("skiplist :", sk)

	c := sk.head
	for {
		if c == nil {
			break
		}

		fmt.Printf("%d(%d)->", c.Val, c.level)
		c = c.Next
	}
}
