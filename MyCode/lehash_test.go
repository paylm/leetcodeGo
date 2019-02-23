package main

import (
	"testing"
)

func Test_TopKFrequent_1(t *testing.T) {
	testData := []struct {
		arr []int
		k   int
		ans []int
	}{
		{arr: []int{1, 1, 1, 2, 2, 3}, k: 2, ans: []int{1, 2}},
		{arr: []int{3, 0, 0}, k: 1, ans: []int{0}},
		{arr: []int{4, 1, -1, 2, -1, 2, 3}, k: 2, ans: []int{-1, 2}},
		{arr: []int{2, 3, 4, 1, 4, 0, 4, -1, -2, -1}, k: 2, ans: []int{4, -1}},
	}
	for ti, data := range testData {
		res := TopKFrequent(data.arr, data.k)

		if len(res) != len(data.ans) {
			t.Errorf("%d test fail , %v not eq  res :%v\n", ti, data.ans, res)
			t.Fail()
		}

		if compareArr(res, data.ans) == false {
			t.Errorf("test fail at %d, input:%v,output:%v,corret:%v\n", ti, data.arr, res, data.ans)
		}
	}
}
