package main

/**https://leetcode-cn.com/problems/top-k-frequent-elements/
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
说明：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。


未通过 oops!!!

**/
func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int) //m 存在的 element 在 kArr 的索引
	kArr := []*tkF{}
	for _, v := range nums {

		ik, ok := m[v]
		if !ok {
			tk := NewtkF(v)
			kArr = Insert(kArr, tk)
			m[v] = len(kArr) - 1
		} else {
			tk := kArr[ik]
			tk.fq++
			//TkfperUp(kArr, ik)
		}

	}

	BuildHeap(kArr)

	res := []int{}
	for i := 0; i < k; i++ {
		tk, ok, arr := Pop(kArr)
		if ok {
			kArr = arr
			res = append(res, tk)
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

	leftChild := i*2 + 1
	k := i

	if leftChild < len(arr) && compare(arr[leftChild], arr[i]) {
		k = leftChild
	}

	if leftChild+1 < len(arr) && compare(arr[leftChild+1], arr[k]) {
		k = leftChild + 1
	}
	if k != i {
		//swapTkf it
		//fmt.Printf("TkfperDown i:%d arr[i]:%v,arr[child]:%v\n", i, arr[i], arr[k])
		swapTkf(arr[k], arr[i])
		TkfperDown(arr, k)
	}
}

func TkfperUp(arr []*tkF, i int) {

	parent := (i - 1) / 2
	if parent < 0 {
		return
	}
	//fmt.Printf("TkfperUp arr[i]:%v,arr[parent]:%v\n", arr[i], arr[parent])
	if compare(arr[i], arr[parent]) {
		//swapTkf it
		swapTkf(arr[i], arr[parent])
		//fmt.Printf("after TkfperUp arr[i]:%v,arr[parent]:%v\n", arr[i], arr[parent])
		TkfperUp(arr, parent)
	}
}

func swapTkf(t1, t2 *tkF) {
	temp := *t1
	*t1 = *t2
	*t2 = temp
}

func Insert(arr []*tkF, t *tkF) []*tkF {
	arr = append(arr, t)
	return arr
}

func BuildHeap(arr []*tkF) {
	l := len(arr)
	for i := l / 2; i >= 0; i-- {
		TkfperDown(arr, i)
	}
}

func Pop(arr []*tkF) (int, bool, []*tkF) {
	if len(arr) < 0 {
		return 0, false, nil
	}

	res := arr[0].key
	last := len(arr)
	swapTkf(arr[0], arr[last-1])
	arr = append(arr[:last-1])
	TkfperDown(arr, 0)

	return res, true, arr
}

//测试工具方法,单元测试调用
//set 简单比较
func compareArr(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	//用hash 方法减少比较次数
	m := make(map[int]bool)
	for _, v := range arr1 {
		m[v] = true
	}

	for _, v := range arr2 {
		_, ok := m[v]
		if !ok {
			return false
		}
	}
	return true
}
