package main

import (
	"testing"
)

func Test_TopKFrequent(t *testing.T) {
	t1 := []int{1, 1, 1, 2, 2, 3}
	t1res := []int{1, 2}
	res := TopKFrequent(t1, 2)

	if len(res) != len(t1res) {
		t.Errorf("test fail , %v not eq  res :%v\n", t1res, res)
		t.Fail()
	}

	for i, v := range t1res {
		if len(res) > i && res[i] == v {
			t.Logf("test pass %d <=> %d \n", v, res[i])
		} else {
			t.Errorf("test fail at %d <=> res:%d", v, res[i])
		}
	}
}
