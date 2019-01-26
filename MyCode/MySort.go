package main

import (
	"fmt"
	"strings"
)

func quickSort(a []int, start int, end int) {

	if len(a) == 0 || start >= end {
		return
	}
	//if len less than 8 , use shell sort
	if end-start < 8 {
		shellSortArea(a, start, end)
		return
	}

	k := a[start]
	i, oi, j, oj := start, start, end, end
	for {
		if i == j {
			break
		}
		for {
			if j > i && a[j] > k {
				j--
			} else if j > i && a[j] == k {
				swap(&(a[j]), &(a[oj]))
				oj--
				j--
			} else {
				break
			}
		}

		for {
			if j > i && a[i] < k {
				i++
			} else if j > i && a[i] == k && i != start {
				swap(&(a[i]), &(a[oi]))
				oi++
				i++
			} else {
				break
			}
		}
		swap(&(a[i]), &(a[j]))
	}

	//a[i] = k
	for o := end; o > oj; o-- {
		swap(&(a[j+1]), &(a[o]))
		j++
	}
	for o := start; o < oi; o++ {
		swap(&(a[i]), &(a[o]))
		i--
	}
	//fmt.Println(a)
	quickSort(a, start, i)
	quickSort(a, j+1, end)
}

func sort3Parttion(a []int, start int, end int) {
	//fmt.Printf("%v,start:%d,end:%d\n", a, start, end)
	if len(a) == 0 || start >= end {
		return
	}

	lt, gt := start, end
	i := lt
	k := a[i]
	for {
		if i > gt {
			break
		}

		if k < a[i] {
			swap(&(a[i]), &(a[gt]))
			gt--
			//i++
		} else if k > a[i] {
			swap(&(a[i]), &(a[lt]))
			lt++

		} else {
			i++
		}
	}
	//fmt.Println(a)
	sort3Parttion(a, start, lt-1)
	sort3Parttion(a, gt+1, end)
}

/***
插入排序
每次找一位,向前移动到合适位置
**/
func insertSort(a []int) {
	if len(a) < 2 {
		return
	}
	i := 1
	for p := 1; p < len(a); p++ {
		temp := a[p]
		for i = p; i > 0 && a[i-1] > temp; i-- {
			a[i] = a[i-1]
		}
		a[i] = temp
	}
}

//Z 字形变换
//https://leetcode-cn.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	sarr := strings.Split(s, "")
	if len(sarr) <= numRows {
		return s
	}
	sgo := make([][]string, numRows)
	k := numRows - 2
	if numRows == 1 {
		k = 0
	}
	for i, v := range sarr {
		a := i % (numRows + k)
		if a < numRows {
			fmt.Println(a, i, v)
			sgo[a] = append(sgo[a], v)
		} else {
			b := a - numRows
			sgo[k-b] = append(sgo[k-b], v)
		}
	}
	//fmt.Println(sgo)
	//return "fuck"
	res := []string{}
	for i := 0; i < numRows; i++ {
		res = append(res, sgo[i]...)
	}
	fmt.Println(res)
	return strings.Join(res, "")
}

func swap(a, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

//把小的数下沉
func perUP(arr []int, l int) {
	for i := l; i >= 0; i-- {
		leftChild := 2*i + 1
		k := i
		if arr[i] < arr[leftChild] {
			k = leftChild
		}

		if len(arr) > leftChild+1 && arr[k] < arr[leftChild+1] {
			k = leftChild + 1
		}
		//fmt.Printf("i:%d,k:%d\n", i, k)
		if k != i {
			swap(&(arr[i]), &(arr[k]))
		}
	}
}

//l 最大边界
func perDown(arr []int, l int) {
	for i := 0; i < l; {
		//fmt.Println("perDown fun() run")
		leftChild := i*2 + 1
		k := i
		if leftChild < l && arr[i] < arr[leftChild] {
			k = leftChild
		}

		if leftChild+1 < l && arr[k] < arr[leftChild+1] {
			k = leftChild + 1
		}

		if k != i {
			swap(&(arr[i]), &(arr[k]))
			i = k
		} else {
			break
		}
	}
}

func shellSortArea(arr []int, start int, end int) {
	if len(arr) < end || start > end {
		return
	}
	gap := end / 2
	for gap > 0 {
		for i := start + gap; i < end; i++ {
			for j := i; j >= gap; j = j - gap {
				if j-gap >= 0 && arr[j] < arr[j-gap] {
					swap(&(arr[j]), &(arr[j-gap]))
				}
			}
		}
		gap = gap / 2
	}
}

func headSort(arr []int) {

	//build head
	perUP(arr, (len(arr)-1)/2)

	for i := len(arr) - 1; i >= 0; i-- {
		//fmt.Println(arr[0])
		swap(&(arr[0]), &(arr[i]))
		perDown(arr, i-1)
	}
}

func shellSort(arr []int) {
	l := len(arr)
	k := 1 //计算时间复杂度
	gap := l / 2
	for gap > 0 {
		for i := gap; i < l; i++ {
			for j := i; j >= gap; j = j - gap {
				k++
				if j-gap >= 0 && arr[j] < arr[j-gap] {
					swap(&(arr[j]), &(arr[j-gap]))
				}
			}
		}
		gap = gap / 2
	}
	fmt.Printf("shell sort :%v,k:%d\n", arr, k)
}
