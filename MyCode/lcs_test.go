package main

import "testing"

func Test_CommonSubLen(t *testing.T) {
	data := []struct {
		w1  string
		w2  string
		ret []string
	}{
		{w1: "fuck", w2: "fabuook", ret: []string{"f", "u", "k"}},
		//	{w1: "aabsdfwerkkop", w2: "abcyp", ret: []string{"a", "b", "p"}},
		{w1: "uploadfile", w2: "addone", ret: []string{"a", "d", "e"}},
	}
	for _, v := range data {

		lcs := commonSubLen(v.w1, v.w2)
		if lcs != len(v.ret) {
			t.Errorf("test fail at %v, ans:%v,but resturn :%v\n", v, v.ret, lcs)
			return
		}
	}
}
func Test_CommonSubArr(t *testing.T) {
	data := []struct {
		w1  string
		w2  string
		ret []string
	}{
		{w1: "fuck", w2: "fabuook", ret: []string{"f", "u", "k"}},
		{w1: "aabsdfwerkkop", w2: "abcyp", ret: []string{"a", "b", "p"}},
		{w1: "uploadfile", w2: "addone", ret: []string{"a", "d", "e"}},
	}
	for _, v := range data {

		lcs := commonSubArr(v.w1, v.w2)
		if len(lcs) != len(v.ret) {
			t.Errorf("test fail at %v, ans:%v,but resturn :%v\n", v, v.ret, lcs)
			return
		}
		for i := 0; i < len(v.ret); i++ {
			if lcs[i] != v.ret[i] {
				t.Errorf("test fail at %v, ans:%v,but resturn :%v\n", v, v.ret, lcs)
			}
		}
	}
}

/**
func Test_CommonSubArr1(t *testing.T) {
	data := []struct {
		w1  string
		w2  string
		ret []string
	}{
		{w1: "fuck", w2: "fabuook", ret: []string{"f", "u", "k"}},
		{w1: "aabsdfwerkkop", w2: "abcyp", ret: []string{"a", "b", "p"}},
		{w1: "uploadfile", w2: "addone", ret: []string{"a", "d", "e"}},
	}
	for _, v := range data {

		lcs := commonSubArr1(v.w1, v.w2)
		if len(lcs) != len(v.ret) {
			t.Errorf("test fail at %v, ans:%v,but resturn :%v\n", v, v.ret, lcs)
			return
		}
		for i := 0; i < len(v.ret); i++ {
			if lcs[i] != v.ret[i] {
				t.Errorf("test fail at %v, ans:%v,but resturn :%v\n", v, v.ret, lcs)
			}
		}
	}
}
**/

func Test_OptArr(t *testing.T) {
	data := []struct {
		arr []int
		opt int
	}{
		{arr: []int{4, 1}, opt: 4},
		{arr: []int{4}, opt: 4},
		{arr: []int{}, opt: 0},
		{arr: []int{4, 1, 1, 9, 2, 3}, opt: 16},
		{arr: []int{4, 1, 10, 9, 2, 3, 7, 1}, opt: 23},
	}
	for _, d := range data {
		nu := opt_arr(d.arr)
		if nu != d.opt {
			t.Errorf("test fail at :%v , answer is :%d , but return :%d\n", d, d.opt, nu)
		}
	}

}

func Test_CombinationSum(t *testing.T) {
	data := []struct {
		arr    []int
		target int
		res    [][]int
	}{
		{arr: []int{2, 3, 6, 7}, target: 7, res: [][]int{{7}, {2, 2, 3}}},
		{arr: []int{2, 3, 5}, target: 8, res: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{arr: []int{7, 8, 3, 4}, target: 11, res: [][]int{{3, 4, 4}, {3, 8}, {4, 7}}},
		{arr: []int{7, 3, 2}, target: 18, res: [][]int{{2, 2, 2, 2, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 3, 3}, {2, 2, 2, 2, 3, 7}, {2, 2, 2, 3, 3, 3, 3}, {2, 2, 7, 7}, {2, 3, 3, 3, 7}, {3, 3, 3, 3, 3, 3}}},
	}
	for _, d := range data {
		testres := combinationSum(d.arr, d.target)
		if len(testres) != len(d.res) {
			t.Errorf("test fail ,return %v\n", testres)
		}
		for _, ires := range testres {

			k := 0
			for _, v := range ires {
				k = k + v
			}
			if k != d.target {
				t.Errorf("test fail @ %v, target :%d , return :%v\n", d.arr, d.target, ires)
			}
		}
	}
}
