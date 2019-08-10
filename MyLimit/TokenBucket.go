package main

import (
	"sync"
	"time"
)

type TkBer interface {
	Acquire() bool
}

type TokenBucket struct {
	r   int64 // 1/r
	cap int64
	bs  int64 //未被消费的token
	mu  *sync.Mutex
}

func NewTokenBucket(r int64, cap int64) *TokenBucket {
	tb := new(TokenBucket)
	tb.r = r
	tb.cap = cap
	tb.mu = new(sync.Mutex)
	go tb.gen()
	return tb
}

//生成TokenBucket
func (bk *TokenBucket) gen() {
	done := false
	for {
		if done {
			return
		}
		bk.bs += bk.r
		if bk.bs >= bk.cap {
			bk.bs = bk.cap
		}
		time.Sleep(1 * time.Second)
	}
}

func (bk *TokenBucket) Acquire() bool {
	bk.mu.Lock()
	defer bk.mu.Unlock()
	if bk.bs > 0 {
		bk.bs--
		return true
	}
	return false
}
