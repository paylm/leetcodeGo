package main

import (
	"testing"
)

func Test_MergeSort(t *testing.T) {
	testData := [][]int{
		{5, 1, 9, 6, 6, 11, 3, 30, 6, 0, -1},
		{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10},
	}

	for i := 0; i < len(testData); i++ {
		MergeSort(testData[i])
		//fmt.Printf("testData:%v\n", testData[i])
		for j := 0; j < len(testData[i])-1; j++ {
			if testData[i][j+1] < testData[i][j] {
				t.Errorf("test fail at case %d,%d is less than priv value %d", i, testData[i][j+1], testData[i][j])
			}
		}
	}
}

func Benchmark_MergeSort(b *testing.B) {
	arr := []int{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10}
	for i := 0; i < b.N; i++ { //use b.N for looping
		MergeSort(arr)
	}
}

func Test_QuickSort(t *testing.T) {
	testData := [][]int{
		{5, 1, 9, 6, 6, 11, 3, 30, 6, 0, -1},
		{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10},
	}

	for i := 0; i < len(testData); i++ {
		QuickSort(testData[i], 0, len(testData[i])-1)
		//fmt.Printf("testData:%v\n", testData[i])
		for j := 0; j < len(testData[i])-1; j++ {
			if testData[i][j+1] < testData[i][j] {
				t.Errorf("test fail at case %d,%d is less than priv value %d , at arr :%v\n", i, testData[i][j+1], testData[i][j], testData[i])
			}
		}
	}
}

func Benchmark_QuickSort(b *testing.B) {
	arr := []int{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10}
	for i := 0; i < b.N; i++ { //use b.N for looping
		QuickSort(arr, 0, len(arr)-1)
	}
}

func Test_HeapSort(t *testing.T) {
	testData := [][]int{
		{5, 1, 9, 6, 6, 11, 3, 30, 6, 0, -1},
		{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10},
	}

	for i := 0; i < len(testData); i++ {
		HeapSort(testData[i])
		//fmt.Printf("testData:%v\n", testData[i])
		for j := 0; j < len(testData[i])-1; j++ {
			if testData[i][j+1] < testData[i][j] {
				t.Errorf("test fail at case %d,%d is less than priv value %d , at arr :%v\n", i, testData[i][j+1], testData[i][j], testData[i])
			}
		}
	}
}

func Benchmark_HeapSort(b *testing.B) {
	arr := []int{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10}
	for i := 0; i < b.N; i++ { //use b.N for looping
		HeapSort(arr)
	}
}
