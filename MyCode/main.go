package main

import (
	"fmt"
	"strings"
)

//从数组中找到两个数的值和为k
func findNumS(a []int, k int) {
	if len(a) == 0 {
		return
	}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
	i, j := 0, len(a)-1
	for {

		fmt.Printf("i:%d,j:%d\n", i, j)
		if i >= j {
			break
		}
		if a[j] > k {
			j--
			continue
		} else {
			if a[j]+a[i] == k {
				fmt.Printf("%d+%d\n", a[i], a[j])
				return
			} else if a[j]+a[i] > k {
				j--
			} else {
				//if a[j]+a[i] <k {
				i++
			}
		}
	}
}

func quickSort(a []int, start int, end int) {
	if len(a) == 0 || start == end {
		return
	}

	k := a[start]
	i, j := start, end
	for {
		if i == j {
			break
		}
		for {
			if j > i && a[j] > k {
				j--
			} else {
				break
			}
		}

		for {
			if j > i && a[i] < k {
				i++
			} else {
				break
			}
		}
		temp := a[i]
		a[i] = a[j]
		a[j] = temp
	}

	a[i] = k
	//fmt.Println(a)
	quickSort(a, start, i)
	quickSort(a, j+1, end)
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
		fmt.Printf("i:%d,k:%d\n", i, k)
		if k != i {
			swap(&(arr[i]), &(arr[k]))
		}
	}
}

//l 最大边界
func perDown(arr []int, l int) {
	for i := 0; i < l; i++ {
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
		}
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

func main() {
	//fmt.Println("vim-go")
	fmt.Println("quickSort")
	a := []int{10, 5, 4, 1, 16, 8, 2, 13, 6, 3}
	//quickSort(a, 0, len(a)-1)
	//findNumS(a, 9)
	convert("fuck123", 3)
	convert("LEETCODEISHIRING", 3)
	fmt.Println(convert("LEETCODEISHIRING", 4))
	fmt.Println(convert("A", 1))
	fmt.Println(convert("AB", 1))
	fmt.Println(a)
	headSort(a)
	fmt.Println(a)
	k1, k2 := 11, 12
	fmt.Println(k1, k2)
	swap(&k1, &k2)
	fmt.Println(k1, k2)
}
