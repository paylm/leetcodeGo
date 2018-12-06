package main

import (
	"errors"
	"fmt"
)

type HashList struct {
	Cap  int
	Data []*ListNode
}

type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
}

func NewListNode(k, v int) *ListNode {
	n := new(ListNode)
	n.Key = k
	n.Val = v
	return n
}

func (n *ListNode) appendNode(k, v int) *ListNode {
	a := NewListNode(k, v)
	if n == nil {
		return a
	}
	prelast, last := n, n
	for {
		if last == nil {
			break
		}
		prelast = last
		last = last.Next
	}
	prelast.Next = a
	return n
}

func (n *ListNode) Get(k int) int {
	if n == nil {
		return -1
	}
	current := n
	for {
		if current.Key == k {
			return current.Val
		}
		if current == nil {
			return -1
		}
		current = current.Next
	}
}

func hash(mod, k int) int {
	return abs(k % mod)
}

func NewHashList(Cap int) *HashList {
	h := new(HashList)
	h.Cap = Cap
	//h.Data = [Cap]*ListNode{}
	h.Data = make([]*ListNode, Cap)
	return h
}

func (h *HashList) Insert(k, v int) {
	i := hash(h.Cap, k)

	if h.Data[i] == nil {
		h.Data[i] = NewListNode(k, v)
	} else {
		h.Data[i].appendNode(k, v)
	}
}

func (h *HashList) Del(k int) {
	i := hash(h.Cap, k)

	current := h.Data[i]
	if current == nil {
		return
	}

	preDel := current
	for {
		if current.Key == k {
			break
		}
		if current == nil {
			return
		}
		preDel = current
		current = current.Next
	}

	if preDel == current {
		//head node need to del
		h.Data[i] = current.Next
	} else {
		preDel.Next = current.Next
	}
}

func (h *HashList) Search(k int) (error, int) {
	i := hash(h.Cap, k)
	current := h.Data[i]
	if current == nil {
		return errors.New(fmt.Sprintf("%d not found", k)), -1
	}

	for {
		if current == nil {
			break
		}
		if current.Key == k {
			return nil, current.Val
		}
		current = current.Next
	}

	return errors.New(fmt.Sprintf("%d has delete error", k)), -1
}
