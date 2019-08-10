package Mycode

import "testing"

func Test_RightSideView(t *testing.T) {
	// t.Fatal("not implemented")
	data := []struct {
		arr []int
		ret []int
	}{
		{
			arr: []int{1, 2, 3, -1, 5, -1, 4}, ret: []int{1, 3, 4},
		},
	}
	for _, d := range data {
		root := buildTree(d.arr, 0)
		res := rightSideView(root)
		if len(res) != len(d.ret) {
			t.Errorf("test fail , len:%v not eq len:%v\n", res, d.ret)
			break
		}
		for i := 0; i < len(res); i++ {
			if res[i] != d.ret[i] {
				t.Errorf("test fail %v, %d not eq res:%d\n", d, res[i], d.ret[i])
			}
		}
	}
}
