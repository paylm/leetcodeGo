package main

import (
	"testing"
	"time"
)

func Test_Acq(t *testing.T) {
	var tbr TkBer
	tbr = NewTokenBucket(10, 50)
	ch := 0
	for i := 0; i < 5; i++ {
		exit := make(chan bool)
		go func() {
			for j := 0; j < 100; j++ {
				if tbr.Acquire() == false {
					t.Logf("%d:get token fail at %d\n", ch, j)
				} else {
					t.Logf("%d:get token success at %d\n", ch, j)
				}
				if j%50 == 0 {
					time.Sleep(1 * time.Second)
				}
			}
			exit <- false
		}()
		ch++
		<-exit
	}
}
