package main

import (
	"fmt"
	"math/rand"
)

var MAX_LEVEL = 16

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
	h := []*Node{}
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

	//step 1 , make Insert Node
	n := NewNode(Val)
	n.level = sk.randomNum(sk.level)
	n.levelNode = make([]*Node, n.level)

	//step 2 :calc level , resize head level if current len(n) > len(head)
	if sk.level > len(sk.head.levelNode) {
		for i := len(sk.head.levelNode); i < n.level; i++ {
			sk.head.levelNode = append(sk.head.levelNode, nil)
		}
	}

	//step 3 , find every level insert before location
	update := make([]*Node, n.level)
	c := sk.head
	for i := n.level - 1; i >= 0; i-- {

		temp := c.levelNode[i]
		//next node is null or next node > Val
		// set current node to insert location
		if temp == nil || temp.Val > Val {
			update[i] = c
			continue
		}

		for {
			//down
			if temp == nil {
				break
			}
			if temp.Val == Val {
				//fmt.Printf("already exsits %d, insert fail\n", Val)
				return false
			}
			if temp.levelNode[i] != nil && temp.levelNode[i].Val > Val {
				update[i] = temp
				c = temp
				break
			}
			//last element < Val
			if temp.Val < Val && temp.levelNode[i] == nil {
				update[i] = temp
				c = temp
				break
			}

			temp = temp.levelNode[i] //to next
		}
	}

	//fmt.Printf("before insert update=>%v,n=>%v\n", update, n)
	//step 4: insert it
	for i := 0; i < n.level; i++ {
		temp := update[i].levelNode[i]
		update[i].levelNode[i] = n
		n.levelNode[i] = temp
		if i == 0 {
			n.fb = update[i]
		}
	}
	//sk.show()

	return true
}

//del one node from all level
func (sk *Skiplist) Del(k int) {

	level := len(sk.head.levelNode)
	update := make([]*Node, level)
	c := sk.head
	for i := level - 1; i >= 0; i-- {
		temp := c.levelNode[i]
		if temp != nil && temp.Val == k {
			update[i] = c
			continue
		}
		for {
			if temp == nil {
				break
			}

			if temp.levelNode[i] != nil && temp.levelNode[i].Val == k {
				c = temp
				update[i] = temp
				break
			}
			temp = temp.levelNode[i]
		}
	}
	//fmt.Println(update)
	for i := level - 1; i >= 0; i-- {
		ldelNodeNext(update[i], i)
	}
}

func (sk *Skiplist) Free() {
	sk.head = nil
}

//删除某节点的i节点
func ldelNodeNext(n *Node, i int) {
	if n == nil {
		return
	}
	delN := n.levelNode[i]
	fmt.Printf("from %d i:%d del %d\n", n.Val, i, delN.Val)
	if len(delN.levelNode) >= i {
		n.levelNode[i] = delN.levelNode[i]
	} else {
		//基本不存在此情况
		n.levelNode[i] = nil
	}
}
func (sk *Skiplist) show() {
	fmt.Println("skiplist :", sk)

	c := sk.head
	for {
		if c == nil {
			break
		}
		fmt.Printf("%d(%d)->", c.Val, len(c.levelNode))
		c = c.levelNode[0]
	}
	fmt.Println()
}

//显示跳跃表，行转列显示
func (sk *Skiplist) showCol() {
	fmt.Println("---- showCol ---- ")
	c := sk.head
	for {
		if c == nil {
			break
		}
		fmt.Printf("%d", c.Val)
		for i := 0; i < len(c.levelNode)-1; i++ {
			if c.levelNode[i] != nil {
				fmt.Printf("<-%d", c.levelNode[i].Val)
			} else {
				fmt.Printf("<-nil")
			}
		}
		c = c.levelNode[0]
		fmt.Println()
	}
}
