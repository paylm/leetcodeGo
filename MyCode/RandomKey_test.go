package main

import (
	"testing"
)

func Test_RandomKey(t *testing.T) {
	for i := 0; i < 100; i++ {
		k := RandomKey()
		//fmt.Printf("k :%d\n", k)
		if k > 100 {
			t.Errorf("test fail k:%d\n", k)
			break
		} else {
			t.Log("test pass")
		}
	}
}
