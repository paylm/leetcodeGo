package main

import "testing"

func TestAvlBSTree_Insert(t *testing.T) {
	n := NewAvlNode(10)
	for _, k := range []int{1, 3, 5, 7, 9, 11, 13, 15, 17} {
		n = Insert(n, k)
	}

	if n.Height > 4 {
		t.Errorf("test fail")
	} else {
		t.Logf("test pass,AvlTree Height :%d\n", n.Height)
	}

}

func Test_AvlFind(t *testing.T) {

	n := NewAvlNode(10)
	for _, k := range []int{1, 3, 5, 7, 9, 11, 13, 15, 17} {
		n = Insert(n, k)
	}

	c3 := AvlFind(n, 3)
	if c3 == nil {
		t.Errorf("test fail , not found %d\n", 3)
	} else {
		t.Logf("test pass ,found %v at key:%d\n", c3, 3)
	}
	c13 := AvlFind(n, 13)
	if c3 == nil {
		t.Errorf("test fail , not found %d\n", 13)
	} else {
		t.Logf("test pass ,found %v at key:%d\n", c13, 13)
	}
}

func Test_AvlDel(t *testing.T) {
	n := NewAvlNode(10)
	for _, k := range []int{1, 3, 5, 7, 9, 11, 13, 15, 17} {
		n = Insert(n, k)
	}

	AvlDel(n, 9)
	c9 := AvlFind(n, 9)
	if c9 != nil {
		t.Errorf("test fail, has del %d,but now found %v\n", 9, c9)
	} else {
		t.Logf("test pass,had del AvlNode %d\n", 9)
	}

	c1 := AvlFind(n, 1)
	if c1 == nil {
		t.Errorf("test fail , not foud %d\n", 1)
	} else {
		t.Logf("test pas , found %v\n", c1)
	}
}
