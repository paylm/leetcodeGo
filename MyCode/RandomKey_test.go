package main

import (
	"fmt"
	"testing"
)

func Test_RandomKey(t *testing.T) {
	for i := 0; i < 100; i++ {
		k := RandomKey()
		fmt.Printf("k :%d\n", k)
		if k > 10 {
			t.Errorf("test fail k:%d\n", k)
		} else {
			t.Log("test pass")
		}
	}
}
