package main

import (
	"errors"
)

var (
	prime = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
)

type NodeValue interface {
}

type HashTreeNode struct {
	k          int
	v          interface{}
	occupied   bool //用occupied来表示节点是否被占据。如果节点的关键字（key）有效，那么occupied应该设置位true，否则设置为false。
	childNodes []*HashTreeNode
}

func NewHashTreeNode(k int, v interface{}) *HashTreeNode {
	n := new(HashTreeNode)
	n.k = k
	n.v = v
	n.occupied = true
	return n
}

func createHashTreeNode(k int, v interface{}, level int) *HashTreeNode {
	n := new(HashTreeNode)
	n.k = k
	n.v = v
	n.occupied = true
	n.childNodes = make([]*HashTreeNode, prime[level])
	return n
}

func (n *HashTreeNode) put(k int, v interface{}) {
	i := 0
	c := n
	for {
		l := k % prime[i] //位置
		//fmt.Printf("k:%d,v:%v,i:%d,l:%d,c:%v\n", k, v, i, l, c)
		if c.childNodes[l] == nil {
			c.childNodes[l] = createHashTreeNode(k, v, i+1)
			return
		} else if c.childNodes[l].occupied == false {
			c.childNodes[l].k = k
			c.childNodes[l].v = v
			c.childNodes[l].occupied = true
			return
		} else if c.childNodes[l].k == k {
			//已存在
			c.childNodes[l].v = v
			c.childNodes[l].occupied = true
			return
		} else {
			//节点已被占用
			c = c.childNodes[l]
			i++
		}
	}
}

//
func (n *HashTreeNode) get(k int) (interface{}, error) {
	i := 0
	c := n

	for {
		l := k % prime[i] //位置
		if c.childNodes[l] == nil {
			return nil, errors.New("not found")
		} else if c.childNodes[l].k == k && c.childNodes[l].occupied == false {
			return nil, errors.New("not found,key is deleted")
		} else if c.childNodes[l].k == k {
			return c.childNodes[l].v, nil
		} else {
			c = c.childNodes[l]
			i++
		}
	}
}

func (n *HashTreeNode) remove(k int) {
	i := 0
	c := n
	for {
		l := k % prime[i] //位置
		if c.childNodes[l] == nil {
			return
		} else if c.childNodes[l].k == k && c.childNodes[l].occupied == false {
			return
		} else if c.childNodes[l].k == k {
			c.childNodes[l].occupied = false
			return
		} else {
			c = c.childNodes[l]
			i++
		}
	}
}
