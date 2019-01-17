package main

import "testing"

func Test_Swap(t *testing.T) {
	// t.Fatal("not implemented")
	e1 := NewElement(0)
	e1.i = 0
	e2 := NewElement(1)
	e2.i = 1
	Swap(e1, e2)
	if e1.i != e1.Weigth {
		t.Logf("test pass")
	} else {
		t.Errorf("test fail at t1:%v\n", e1)
	}
}

func Test_Swap_1(t *testing.T) {
	// t.Fatal("not implemented")
	e1 := NewElement(0)
	e1.i = 0
	e2 := NewElement(1)
	e2.i = 1
	arr := []*Element{e1, e2}
	t.Logf("test before:%v,e1:%v\n", arr, &e1)
	Swap(arr[0], arr[1])
	if arr[0].i != arr[0].Weigth {
		t.Logf("test pass:%v,e1:%v\n", arr, &e1)
	} else {
		t.Errorf("test fail at t1:%v\n", e1)
	}
}

func TestMinHeap_Push(t *testing.T) {
	mhp := NewMinHeap(10)
	for i := 9; i >= 0; i-- {
		mhp.Push(NewElement(i))
	}

	if mhp.harr[0].Weigth != 0 {
		t.Errorf("test fail: top %v , but corret is :{0,0}\n", mhp.harr[0])
	} else {
		t.Logf("test pass")
	}
}
