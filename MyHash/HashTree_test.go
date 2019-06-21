package main

import (
	"math/rand"
	"testing"
)

func Test_Put(t *testing.T) {
	n := createHashTreeNode(0, 0, 0)
	for i := 0; i < 2000000; i++ {
		k := rand.Intn(99999)
		n.put(k, i)
	}
}

func Benchmark_Put(b *testing.B) {
	n := createHashTreeNode(0, 0, 0)
	for i := 0; i < b.N; i++ { //use b.N for looping
		k := rand.Intn(99999)
		n.put(k, 8888)
	}
}

func Test_Get(t *testing.T) {
	n := createHashTreeNode(0, 0, 0)
	data := []int{}
	for i := 0; i < 2000000; i++ {
		k := rand.Intn(99999)
		n.put(k, i)
		if i%30 == 0 {
			data = append(data, k)
		}
	}

	//test get
	for _, d := range data {
		if _, err := n.get(d); err != nil {
			t.Errorf("test fail k %d , error :%v\n", d, err)
		}

	}
}

func Benchmark_Get(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	n := createHashTreeNode(0, 0, 0)
	for i := 0; i < 2000000; i++ {
		k := rand.Intn(99999)
		n.put(k, i)
	}
	b.StartTimer()
	//test get
	for i := 0; i < b.N; i++ { //use b.N for looping
		n.get(i)
	}
}

func Test_Remove(t *testing.T) {

	n := createHashTreeNode(0, 0, 0)
	data := []int{}
	for i := 0; i < 2000000; i++ {
		k := rand.Intn(99999)
		n.put(k, i)
		if i%30 == 0 {
			data = append(data, k)
		}
	}

	//test
	for _, d := range data {
		n.remove(d)
		if _, err := n.get(d); err == nil {
			t.Errorf("test fail k %d , error :%v\n", d, err)
		}

	}

}

func Benchmark_Remove(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	n := createHashTreeNode(0, 0, 0)
	for i := 0; i < 2000000; i++ {
		k := rand.Intn(99999)
		n.put(k, i)
	}
	b.StartTimer()
	//test get
	for i := 0; i < b.N; i++ { //use b.N for looping
		n.remove(i)
	}
}
