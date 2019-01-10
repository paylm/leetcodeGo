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
	fb        *Node   //上一节点的指针
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

func (sk *Skiplist) Search(Val int) *Node {
	top := len(sk.head.levelNode) - 1
	c := sk.head

	for i := top; i >= 0; i-- {
		if c.levelNode[i] == nil {
			continue
		}

		temp := c.levelNode[i]
		if temp.Val == Val {
			return temp
		}
		for {
			if temp == nil {
				break
			}
			if temp.Val > Val {
				break
			}

			if temp.levelNode[i] != nil && temp.levelNode[i].Val > Val {
				c = temp
				break
			}
			if temp.Val < Val && temp.levelNode[i] == nil {
				c = temp
				break
			}

			temp = temp.levelNode[i]
		}
	}

	if c.Val == Val {
		//fmt.Printf("already exist %d ,insert fail \n", Val)
		return c
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
			//fmt.Printf("already exits %d , insert fail\n", Val)
			return false
		}
		for {
			//fmt.Printf("level:%d,Val:%d,temp:%v\n", i, Val, temp)
			if temp == nil {
				break
			}
			if temp.Val > Val {
				break
			}

			if temp.levelNode[i] != nil && temp.levelNode[i].Val > Val {
				c = temp
				//fmt.Printf("下一节点%v >=%d 下沉处理\n", temp.levelNode[i], Val)
				break
			}
			if temp.Val < Val && temp.levelNode[i] == nil {
				c = temp
				//fmt.Printf("已到最后节点%v >=%d 右移处理 \n", temp.levelNode[i], Val)
				break
			}

			//fmt.Printf("%d %v 找不到符合位置,到下一位找\n", temp.Val, temp)
			temp = temp.levelNode[i]
		}
	}

	if c.Val == Val {
		fmt.Printf("already exist %d ,insert fail \n", Val)
		return false
	}

	//c 为fb 位置
	fmt.Printf("insert point:%d,after:%d\n", Val, c.Val)
	n := NewNode(Val)
	n.level = sk.randomNum(sk.level)
	n.levelNode = make([]*Node, n.level)
	n.fb = c
	fmt.Printf("add n:%v\n", n)
	//current := sk.head
	for i := 0; i < n.level; i++ {
		if i == 0 {
			temp := c.levelNode[i]
			c.levelNode[i] = n
			n.levelNode[i] = temp
		} else {
			//多层处理
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
		c = c.levelNode[0]
	}
	fmt.Println()
}
