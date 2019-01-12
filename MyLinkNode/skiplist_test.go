package main

import (
	"math/rand"
	"testing"
)

func TestRandomKey_RandomNum(t *testing.T) {
	r := &RandomKey{}
	for i := 0; i < 100; i++ {
		rv := r.RandomNum(10)
		if rv > 10 {
			t.Errorf("%d randRandomNum(%d) is great than 10 , test fail\n", i, rv)
		} else {
			t.Logf("%d => %d test randRandomNum pass\n", i, rv)
		}
	}
}

func TestSkiplist_Insert(t *testing.T) {
	testArr := [...][]int{{1, 14, 6, 97, 9, 8, 14}, {1, 1, 1, 1, 1, 5, 8}, {-100, 2}}
	for _, arr := range testArr {
		sk := NewSkiplist(&RandomKey{})
		size := 0
		for i := 0; i < len(arr); i++ {
			ok := sk.Insert(arr[i])
			if ok {
				size++
			}
		}
		if size > len(arr) {
			t.Errorf("test fail at %v,skiplist szie:%d\n", arr, size)
		} else {
			t.Logf("test pass at %v, skiplist size:%d\n", arr, size)
		}
	}
}

// go test -test.bench=".*"
func BenchmarkRandomKey_RandomNum(b *testing.B) {
	rk := &RandomKey{}
	for i := 0; i < b.N; i++ {
		rk.RandomNum(16)
	}
}

func BenchmarkSkiplist_Insert(b *testing.B) {
	sk := NewSkiplist(&RandomKey{})
	for i := 0; i < b.N; i++ {
		sk.Insert(rand.Intn(1000))
	}
}
func BenchmarkSkiplist_Del(b *testing.B) {
	sk := NewSkiplist(&RandomKey{})
	for i := 0; i < b.N; i++ {
		sk.Del(rand.Intn(1000))
	}
}
