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

func Test_CombinationSum2(t *testing.T) {
	data := []struct {
		arr    []int
		target int
		res    [][]int
	}{
		{arr: []int{10, 1, 2, 7, 6, 1, 5}, target: 8, res: [][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}}},
		{arr: []int{2, 5, 2, 1, 2}, target: 5, res: [][]int{{1, 2, 2}, {5}}},
	}
	for _, d := range data {
		testres := combinationSum2(d.arr, d.target)
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

func Test_DominantIndex(t *testing.T) {
	data := []struct {
		nums   []int
		target int
	}{
		{nums: []int{3, 6, 1, 0}, target: 1},
		{nums: []int{1, 2, 3, 4}, target: -1},
		{nums: []int{1, 0}, target: 0},
		{nums: []int{0, 0, 3, 2}, target: -1},
	}
	for _, d := range data {
		ret := dominantIndex(d.nums)
		if ret != d.target {
			t.Errorf("test fail , %v should be %d , but return %d\n", d.nums, d.target, ret)
		}
	}
}

func Test_MaxSubArray(t *testing.T) {

	data := []struct {
		nums   []int
		target int
	}{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, target: 6},
		{nums: []int{2, -3}, target: 2},
	}
	for _, d := range data {
		ret := maxSubArray(d.nums)
		if ret != d.target {
			t.Errorf("test fail , %v should be %d , but return %d\n", d.nums, d.target, ret)
		}
	}
}

func Test_ThreeSum(t *testing.T) {
	data := []struct {
		nums   []int
		target int
		res    [][]int
	}{
		{nums: []int{-1, 0, 1, 2, -1, -4}, target: 6, res: [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 00, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, target: 6, res: [][]int{{0, 0, 0}}},
		{nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, target: 6, res: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
		{nums: []int{0, -4, -1, -4, -2, -3, 2}, target: 8, res: [][]int{{-2, 0, 2}}},
	}
	for _, d := range data {
		testres := threeSum(d.nums)
		if len(testres) != len(d.res) {
			t.Errorf("test fail ,return %v\n", testres)
		}
		for _, ires := range testres {

			if len(testres) != len(d.res) {

				t.Errorf("test fail @ %v, target :%d , return :%v\n", d.nums, d.target, ires)
			}
			k := 0
			for _, v := range ires {
				k = k + v
			}
			if k != 0 {
				t.Errorf("test fail @ %v, target :%d , return :%v\n", d.nums, d.target, ires)
			}
		}
	}

}

func Test_LongestCommonPrefix(t *testing.T) {

	data := []struct {
		strs   []string
		target string
	}{
		{strs: []string{"flower", "flow", "flight"}, target: "fl"},
		{strs: []string{"dog", "racecar", "car"}, target: ""},
		{strs: []string{}, target: ""},
	}
	for _, d := range data {
		ret := longestCommonPrefix(d.strs)
		if ret != d.target {
			t.Errorf("test fail , %v should be %s , but return %s\n", d.strs, d.target, ret)
		}
	}
}

func Test_IntegerBreak(t *testing.T) {

	data := []struct {
		input  int
		output int
	}{
		{input: 2, output: 1},
		{input: 10, output: 36},
		{input: 5, output: 6},
		{input: 6, output: 9},
	}
	for _, d := range data {
		ret := integerBreak(d.input)
		if ret != d.output {
			t.Errorf("test fail , %v should be %d , but return %d\n", d, d.output, ret)
		}
	}
}

func Test_ClimbStairs(t *testing.T) {

	data := []struct {
		input  int
		output int
	}{
		{input: 2, output: 2},
		{input: 3, output: 3},
	}
	for _, d := range data {
		ret := climbStairs(d.input)
		if ret != d.output {
			t.Errorf("test fail , %v should be %d , but return %d\n", d, d.output, ret)
		}
	}
}

func Test_Rob(t *testing.T) {
	data := []struct {
		arr    []int //对于-1的值，为生二叉树为空节点
		result int
	}{
		{arr: []int{3, 2, 3, -1, 3, -1, 1}, result: 7},
		{arr: []int{3, 4, 5, 1, 3, -1, 1}, result: 9},
		//	{arr: []int{2, -1, 3, 4, -1, 1}, result: 6},
	}

	for _, d := range data {
		//make tree
		tn := buildTreeByArr(d.arr, 0)
		//	inOrderTree(tn)
		res := rob(tn)
		if res == d.result {

		} else {
			t.Errorf("test fail at %v , result:%d , but return %d\n", d.arr, d.result, res)
		}
	}
}

func Test_VideoStitching(t *testing.T) {
	data := []struct {
		clips [][]int
		T     int
		res   int
	}{
		{clips: [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, T: 10, res: 3},
		{clips: [][]int{{0, 1}, {1, 2}}, T: 5, res: -1},
		{clips: [][]int{{0, 4}, {2, 8}}, T: 5, res: 2},
	}
	for _, d := range data {
		rs := videoStitching(d.clips, d.T)
		if rs != d.res {
			t.Errorf("test fail at %v  , return %d\n", d, rs)
		}
	}
}
