package main

import "testing"

func Test_MajorityElement1(t *testing.T) {
	data := []struct {
		input  []int
		output int
	}{
		{input: []int{3, 2, 3}, output: 3},
		{input: []int{1, 1, 1, 3, 3, 2, 1, 2}, output: 1},
	}
	for _, v := range data {
		ret := majorityElement1(v.input)
		if ret != v.output {
			t.Errorf("test fail at %v ,return %v\n", v, ret)
		}
	}
}
func Test_MajorityElement(t *testing.T) {
	data := []struct {
		input  []int
		output []int
	}{
		{input: []int{3, 2, 3}, output: []int{3}},
		{input: []int{1, 1, 1, 3, 3, 2, 2, 2}, output: []int{1, 2}},
		{input: []int{1, 1}, output: []int{1}},
		{input: []int{0, 0, 0}, output: []int{0}},
	}
	for _, v := range data {
		ret := majorityElement(v.input)
		if len(ret) != len(v.output) {
			t.Errorf("test fail at %v ,return %v\n", v, ret)
		}
	}
}
