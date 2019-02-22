package main

import "testing"

func Test_CommonSubLen(t *testing.T) {
	lcs := commonSubLen("fuck", "fabuook")
	if lcs != 3 {
		t.Errorf("test fail, is should be %d , but return %d", 3, lcs)
	}
}
func Test_CommonSubArr(t *testing.T) {
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
