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
	//	defer func() {
	//		if err := recover(); err != nil {
	//			fmt.Printf("a:%v,start:%d,end:%d,err:%v\n", a, start, end, err)
	//		}
	//	}()

	if len(a) == 0 || start >= end {
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
	fmt.Printf("%v,start:%d,end:%d\n", a, start, end)
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

/**
回文数
https://leetcode-cn.com/problems/palindrome-number/
**/
func isPalindrome(x int) bool {
	if x > 0 && x < 10 {
		return true
	}
	if x < 0 {
		return false
	}

	arr := []int{}
	t := x
	for t > 0 {
		y := t % 10
		arr = append(arr, y)
		t = t / 10
	}
	fmt.Println(arr)
	i, j := 0, len(arr)-1
	for {
		if i >= j {
			break
		}

		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}

	return true
}

/**
n*n

i,j			j,n-i



n-i,j			n-i,n-j
**/

func rotate(matrix [][]int) {

	n := len(matrix) - 1

	for i := 0; i < n; i++ {
		for j := i; j < n && j < (n-2*i)/2+1; j++ {
			k := matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = matrix[j][n-i]
			matrix[j][n-i] = k
			fmt.Printf("after rotate %d:%d\n", i, j)
			showMatrix(matrix)
		}
	}
}

func showMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}

func main() {
	//fmt.Println("vim-go")
	fmt.Println("quickSort")
	a := []int{10, 5, 4, 1, 16, 10, 5, 8, 13, 2, 10, 6, 3}
	//a := []int{10, 10, 10, 10, 10}
	b := []int{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10}
	c := []int{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10}
	d := []int{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)

	fmt.Println("三路快排")
	sort3Parttion(b, 0, len(b)-1)
	fmt.Println(b)
	fmt.Println("insert Sort")
	insertSort(c)
	fmt.Println(c)
	fmt.Println("shell sort")
	shellSort(d)
	fmt.Println(d)
	//findNumS(a, 9)
	convert("fuck123", 3)
	//convert("LEETCODEISHIRING", 3)
	//fmt.Println(convert("LEETCODEISHIRING", 4))
	fmt.Println(convert("A", 1))
	fmt.Println(convert("AB", 1))
	//	fmt.Println(a)
	//	headSort(a)
	//	fmt.Println(a)
	//	k1, k2 := 11, 12
	//	fmt.Println(k1, k2)
	//	swap(&k1, &k2)
	//	fmt.Println(k1, k2)
	fmt.Println(isPalindrome(121))

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	matrix4 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	rotate(matrix4)
}
