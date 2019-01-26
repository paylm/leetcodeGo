package main

import (
	"fmt"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	testData := [][]int{
		{5, 1, 9, 6, 6, 11, 3, 30, 6, 0, -1},
		{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10},
	}

	for i := 0; i < len(testData); i++ {
		MergeSort(testData[i])
		fmt.Printf("testData:%v\n", testData[i])
		tmp := testData[i][0]
		for j := 0; j < len(testData[i]); j++ {
			if testData[i][j] < tmp {
				t.Errorf("test fail,%d is less than priv value %d", testData[i][j], tmp)
			}
			tmp = testData[i][j]
		}
	}
}

func Test_QuickSort(t *testing.T) {
	testData := [][]int{
		{5, 1, 9, 6, 6, 11, 3, 30, 6, 0, -1},
		{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10},
	}

	for i := 0; i < len(testData); i++ {
		QuickSort(testData[i], 0, len(testData[i])-1)
		fmt.Printf("testData:%v\n", testData[i])
		tmp := testData[i][0]
		for j := 0; j < len(testData[i]); j++ {
			if testData[i][j] < tmp {
				t.Errorf("test fail,%d is less than priv value %d , at arr :%v\n", testData[i][j], tmp, testData[i])
			}
			tmp = testData[i][j]
		}
	}
}

func Test_HeapSort(t *testing.T) {
	testData := [][]int{
		{5, 1, 9, 6, 6, 11, 3, 30, 6, 0, -1},
		{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10},
	}

	for i := 0; i < len(testData); i++ {
		HeapSort(testData[i])
		fmt.Printf("testData:%v\n", testData[i])
		tmp := testData[i][0]
		for j := 0; j < len(testData[i]); j++ {
			if testData[i][j] < tmp {
				t.Errorf("test fail,%d is less than priv value %d , at arr :%v\n", testData[i][j], tmp, testData[i])
			}
			tmp = testData[i][j]
		}
	}
}
