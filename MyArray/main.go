package main

import "fmt"

func QuickSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	k := arr[start]
	i, j := start, end
	for {
		if i == j {
			break
		}
		for {
			//move j,for right to left ,如果找到小于k , 到下一步
			if j > i && arr[j] > k {
				j = j - 1
			} else {
				fmt.Printf("move j:%d v:%d \n", j, arr[j])
				break
			}
		}
		for {
			//移动i , 从左到右找,如果找到大于k ,到下一步
			if j > i && arr[i] < k {
				i = i + 1
			} else {
				fmt.Printf("move i:%d  v:%d \n", i, arr[i])
				break
			}
		}
		if j > i {
			fmt.Printf("change location , i=%d,j=%d \n", i, j)
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	arr[i] = k
	fmt.Println(arr)
	//return
	QuickSort(arr, start, i-1)
	QuickSort(arr, i+1, end)
}

func main() {

	fmt.Println("vim-go")
	arr := []int{4, 3, 8, 1, 6, 5, 2}
	fmt.Println(arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
