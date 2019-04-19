package main

import (
	"fmt"
	"testing"
)

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

func Test_InsertSortList(t *testing.T) {

	tdata1 := NewListNode(4)
	tdata1.Next = NewListNode(2)
	tdata1.Next.Next = NewListNode(1)
	tdata1.Next.Next.Next = NewListNode(3)

	res := insertionSortList(tdata1)
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

	res2 := insertionSortList(tdata2)
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

func Test_InsertSortList2(t *testing.T) {
	data := []struct {
		input []int
	}{
		{input: []int{}},
		{input: []int{1, 2, 3, 4, 5}},
		{input: []int{-1, 5, 3, 4, 0}},
		{input: []int{3, 4, 5, 1, 89, 11, 9, 33, 23, 6, 8, 2}},
	}
	for _, v := range data {
		head := ArrToLinkList(v.input)
		fmt.Printf("<<<<<<======InsertSortList======%v>>>>>>\n", v.input)
		//showList(head)
		res2 := insertionSortList(head)
		resArr := []int{}
		if res2 == nil {
			continue
		}
		v2 := res2.Val
		resArr = append(resArr, v2)
		c2 := res2.Next
		for {
			if c2 == nil {
				break
			}
			if c2.Val < v2 {
				t.Errorf("test fail at %v \n", v)
				break
			}
			resArr = append(resArr, c2.Val)

			c2 = c2.Next
		}
		if len(resArr) != len(v.input) {
			t.Errorf("test fail at %v return :%v\n", v, resArr)
		}
	}
}
