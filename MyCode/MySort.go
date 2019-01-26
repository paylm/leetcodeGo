package main

import (
	"fmt"
)

func QuickSort(a []int, start int, end int) {

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
	QuickSort(a, start, i)
	QuickSort(a, j+1, end)
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

func swap(a, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

//把大的数上浮
func perUP(arr []int, i int) {
	parent := (i - 1) / 2
	if parent < 0 {
		return
	}

	if len(arr) > i && arr[parent] < arr[i] {
		swap(&(arr[i]), &(arr[parent]))
		perUP(arr, parent)
	}
}

//l 最大边界
func perDown(arr []int, i int, l int) {
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
		perDown(arr, k, l)
	}
}

func shellSortArea(arr []int, start int, end int) {
	if len(arr) < end || start > end {
		return
	}
	gap := end / 2
	for gap > 0 {
		for i := start + gap; i <= end; i++ {
			for j := i; j >= gap; j = j - gap {
				if j-gap >= 0 && arr[j] < arr[j-gap] {
					swap(&(arr[j]), &(arr[j-gap]))
				}
			}
		}
		gap = gap / 2
	}
}

func HeapSort(arr []int) {

	//build head
	for i := len(arr) - 1; i >= len(arr)/2; i-- {
		perUP(arr, i)
	}

	fmt.Printf("after per Up:%v\n", arr)

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("heap max:%v,arr:%v\n", arr[0], arr)
		swap(&(arr[0]), &(arr[i]))
		perDown(arr, 0, i-1)
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

//归并排序
func MergeSort(arr []int) {
	tmpArr := make([]int, len(arr))
	if len(arr) <= 0 {
		return
	}
	mSort(arr, tmpArr, 0, len(arr)-1)
}

func mSort(arr []int, tmpArr []int, left int, right int) {
	if left < right {
		center := (left + right) / 2
		mSort(arr, tmpArr, left, center)
		mSort(arr, tmpArr, center+1, right)
		merge(arr, tmpArr, left, center, right)
	}
}

//merge tmpArr[left to center] && tmpArr[center to right] to arr[left-right]
func merge(arr []int, tmpArr []int, left int, center int, right int) {
	i, k, j := left, left, center+1
	for {
		if i > center || j > right {
			break
		}

		if arr[i] >= arr[j] {
			tmpArr[k] = arr[j]
			j++
		} else if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		}
		k++
	}
	//fmt.Printf("step 1 tmpArr:%v,left:%d,center:%d,right:%d\n", tmpArr, left, center, right)

	if j > right && i <= center {
		for {
			if i > center {
				break
			}
			tmpArr[k] = arr[i]
			k++
			i++
		}
	}
	if i > center && j <= right {
		for {
			if j > right {
				break
			}
			tmpArr[k] = arr[j]
			k++
			j++
		}
	}
	//fmt.Printf("tmpArr:%v,left:%d,center:%d,right:%d\n", tmpArr, left, center, right)
	//copy it back to arr
	for i := left; i <= right; i++ {
		arr[i] = tmpArr[i]
	}
	//fmt.Printf("arr:%v\n", arr)
}
