package main

import (
	"testing"
)

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
		t.Logf("test pass,Top")
	}
}

func TestMinHeap_Pop(t *testing.T) {
	mhp := NewMinHeap(10)
	for i := 10; i > 0; i-- {
		mhp.Push(&Element{Weigth: i * 10, Val: i})
	}

	i := -1
	for {
		if mhp.empty() {
			break
		}
		err, e := mhp.Pop()
		if err != nil {
			t.Errorf("test fail , pop out nil ,err:%v\n", err)
		}
		if e.Weigth > i {
			//		t.Logf("test pass , pop min:%v\n", e)
			i = e.Weigth
		} else {
			t.Errorf("test fail , pop min(%v) is great than last,", e)
		}
	}
}

func TestMinHeap_DecreseKey(t *testing.T) {
	mhp := NewMinHeap(10)
	for i := 9; i >= 0; i-- {
		mhp.Push(NewElement(i * 10))
	}

	e := mhp.harr[0]
	e.Val = 10100
	e.Weigth = 300
	mhp.DecreseKey(e)

	err, res := mhp.Pop()
	if err != nil {
		t.Errorf("test fail,pop min with err:%v\n", err)
	}

	if res.Weigth == 10 {
		t.Logf("test pass ,min head %v\n", res)
	} else {
		t.Errorf("test fail, min pop (%v) is not corret\n", res)
	}
}

func Test_SwapEdge(t *testing.T) {
	e1 := NewEdge(6, 4, 2)
	e2 := NewEdge(5, 2, 1)
	swapEdge(e1, e2)
	if e1.Weigth == 6 {
		t.Errorf("test fail e1:%v\n", e1)
	}
	if e2.Weigth == 5 {
		t.Errorf("test fail e2:%v\n", e2)
	}
}

func TestMinEdge_Push(t *testing.T) {
	mhp := NewMinEdgeHeap(6)
	for i := 6; i > 0; i-- {
		mhp.Push(NewEdge(i, 1, i+1))
	}

	k := 0
	for {
		if mhp.empty() {
			break
		}
		e := mhp.Pop()
		//t.Logf("e:%v,%p\n", e, e)
		if e.Weigth >= k {
			k = e.Weigth
		} else {
			t.Errorf("test fail, %v < %d\n", e, k)
		}
	}
}
