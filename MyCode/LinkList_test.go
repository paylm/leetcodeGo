package main

import "testing"

func Test_SortList(t *testing.T) {

	tdata1 := NewListNode(4)
	tdata1.Next = NewListNode(2)
	tdata1.Next.Next = NewListNode(1)
	tdata1.Next.Next.Next = NewListNode(3)

	res := sortList(tdata1)
	v := res.Val
	c := res.Next
	for {
		if c == nil {
			break
		}
		if c.Val < v {
			t.Errorf("test fail,%v < %d \n ", c.Val, v)
			break
		}
		v = c.Val
		c = c.Next
	}

	tdata2 := NewListNode(6)
	tdata2.Next = NewListNode(7)
	tdata2.Next.Next = NewListNode(8)
	tdata2.Next.Next.Next = NewListNode(9)

	res2 := sortList(tdata2)
	v2 := res2.Val
	c2 := res2.Next
	for {
		if c2 == nil {
			break
		}
		if c2.Val < v2 {
			t.Errorf("test fail,%v < %d \n ", c2.Val, v2)
			break
		}
		v2 = c2.Val
		c2 = c2.Next
	}
}
