package main

import "fmt"

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int) //m 存在的 element 在 kArr 的索引
	kArr := []*tkF{}
	for _, v := range nums {

		ik, ok := m[v]
		fmt.Printf("ok:%v,v:%d\n", ok, v)
		if !ok {
			tk := NewtkF(v)
			kArr = Insert(kArr, tk)
			m[v] = len(kArr) - 1
		} else {
			fmt.Printf("kArr:%v, v:%d ik:%d\n", kArr, v, ik)
			tk := kArr[ik]
			tk.fq++
			TkfperUp(kArr, ik)
		}

	}

	res := []int{}
	for i := 0; i < k; i++ {
		tk, ok := Pop(kArr)
		if ok {
			res = append(res, tk.key)
		}
	}
	return res
}

type tkF struct {
	key int
	fq  int
}

func NewtkF(k int) *tkF {
	tk := new(tkF)
	tk.key = k
	return tk
}

func compare(t1, t2 *tkF) bool {
	if t1.fq > t2.fq {
		return true
	}
	return false
}

func TkfperDown(arr []*tkF, i int) {

	leftChild := i * 2
	k := i

	if leftChild < len(arr) && compare(arr[leftChild], arr[i]) {
		k = leftChild
	}

	if leftChild+1 < len(arr) && compare(arr[leftChild+1], arr[i]) {
		k = leftChild + 1
	}
	if k != i {
		//swapTkf it
		swapTkf(arr[k], arr[i])
		TkfperDown(arr, k)
	}
}

func TkfperUp(arr []*tkF, i int) {

	parent := (i - 1) / 2
	if parent < 0 {
		return
	}
	if compare(arr[i], arr[parent]) {
		//swapTkf it
		swapTkf(arr[i], arr[parent])
		TkfperUp(arr, parent)
	}
}

func swapTkf(t1, t2 *tkF) {
	temp := t1
	t1 = t2
	t2 = temp
}

func Insert(arr []*tkF, t *tkF) []*tkF {
	arr = append(arr, t)
	return arr
}

func Pop(arr []*tkF) (*tkF, bool) {
	if len(arr) < 0 {
		return nil, false
	}

	res := arr[0]
	last := len(arr)
	swapTkf(arr[0], arr[last-1])
	arr = append(arr[:last-1])
	TkfperDown(arr, 0)

	return res, true
}
