package main

type HashMap interface {
	Insert(k, v int)
	Search(k, v int) int
	Del(k int)
}

type HashList struct {
	Cap  int
	Data []*ListNode
}

type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
}

func NewListNode(k, v int) {
	n := new(ListNode)
	n.Key = k
	n.Val = v
	return n
}

func (n *ListNode) Add(k, v int) *ListNode {
	a := NewListNode(k, v)
	if n == nil {
		return a
	}
	n.Next = a
	return n
}

func (n *ListNode) Get(k int) int {
	if n == nil {
		return nil
	}
	current := n
	for {
		if current.Key == k {
			return current.Val
		}
		if current == nil {
			return nil
		}
		current = current.Next
	}
}

func hash(mod, k int) int {
	return abs(mod / k)
}

func NewHashList(cap int) *HashList {
	h := new(HashList)
	h.Cap = cap
	h.Data = [cap]*ListNode{}
	return h
}

func (h *HashList) Insert(k, v int) {
	i := hash(h.Cap, k)
	if h.Data[i] == nil {
		h.Data[i] = NewListNode(k, v)
	} else {
		h.Data[i].Add(k, v)
	}
}

func (h *HashList) Del(k int) {

}

func (h *HashList) Search(k int) int {
	i := hash(h.Cap, k)
	return h.Data[i].Get(i)
}
