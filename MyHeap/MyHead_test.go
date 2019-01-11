package main

import "testing"

func TestMinHeap_Insert(t *testing.T) {
	h := NewMinHeap(8)

	for _, v := range []int{6, 1, 8, 3, 10, 5, 4, 9} {
		h.Insert(v)
	}

	for _, tv := range []int{1, 3, 4, 5, 6, 8, 9} {
		if h.isEmpty() {
			t.Error("test err, heap is null")
		}

		v, e := h.PopRoot()
		if e != nil {
			t.Errorf("test fail ,pop Min err:%v\n", e)
		}

		if v == tv {
			t.Logf("test pass ,%d == %d\n", v, tv)
		} else {
			t.Errorf("test err , %d != %d\n", v, tv)
		}
	}
}

func TestMinHeap_Insert_1(t *testing.T) {
	h := NewMinHeap(8)

	for _, v := range []int{6, 1, 8, 3, 10, 5, 4, 9} {
		h.Insert(v)
	}

	for i := 0; i < 10; i++ {
		if h.isEmpty() {
			t.Error("test err, heap is null")
		}

		v, e := h.PopRoot()
		if e != nil {
			t.Errorf("test fail ,pop Min err:%v\n", e)
		}

		h.Insert(1)
		if v == 1 {
			t.Logf("test pass ,%d == 1 \n", v)
		} else {
			t.Errorf("test err , %d != 1\n", v)
		}
	}
}
