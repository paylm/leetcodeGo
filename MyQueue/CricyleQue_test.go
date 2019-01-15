package main

import "testing"

func TestCricyleQue_Push(t *testing.T) {
	cq := NewCricyleQue(10)
	res := false
	for i := 0; i < 11; i++ {
		err := cq.Push(i * 10)
		if err != nil {
			t.Logf("test at push %d, err:%v\n", i*10, err)
			res = true
		}
	}
	if res {
		t.Log("test pass")
	} else {
		t.Error("test fail")
	}
}

func TestCricyleQue_Pop(t *testing.T) {
	cq := NewCricyleQue(50)
	ls := []int{}
	for i := 0; i < 70; i++ {
		cq.Push(i)
		ls = append(ls, i)
	}
	i := 0
	for {
		if cq.empty() {
			break
		}
		v, _ := cq.Pop()
		if ls[i] != v {
			t.Errorf("test fail , %d != Que %d \n", ls[i], v)
		}
		i++
	}
}
