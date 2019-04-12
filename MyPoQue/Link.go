package main

import (
	"fmt"
	"time"
)

type AE struct {
	Name  string `json:"name"`
	Event string
	Prov  string
}

type QueLink struct {
	que   map[int64][]*AE
	done  chan bool
	chPop chan []*AE
}

func NewQueLink() *QueLink {
	ql := new(QueLink)
	ql.que = make(map[int64][]*AE)
	ql.done = make(chan bool)
	ql.chPop = make(chan []*AE)
	go ql.expire()
	return ql
}

/***
	多少秒后从que 弹出来
**/
func (q *QueLink) Push(d *AE, s int) {

	k := time.Now().UTC().Unix() + int64(s)
	data, ok := q.que[k]
	if !ok {
		data = []*AE{d}
	} else {
		data = append(data, d)
	}
	q.que[k] = data
	fmt.Printf("Push key:%v\n", k)
}

func (q *QueLink) AutoPop(rec chan []*AE) {
	done := false
	for {
		if done {
			break
		}
		select {
		case aes := <-q.chPop:
			rec <- aes
		case <-q.done:
			fmt.Printf("exit que with que is empty")
			done = true
			break
		}
	}
}

/***
自动清除过期的key
**/
func (q *QueLink) expire() {

	for {
		time.Sleep(time.Second)
		k := time.Now().UTC().Unix()
		//fmt.Printf("auto expire key:%v\n", k)
		data, ok := q.que[k]
		if !ok {
			continue
		}
		q.chPop <- data
		fmt.Printf("expire %v:%v\n", k, data)
		delete(q.que, k)
		if len(q.que) == 0 {
			q.done <- true
		}
	}
}

func main() {
	fmt.Println("vim-go")
	q := NewQueLink()
	q.Push(&AE{Name: "666", Event: "lock_6"}, 15)
	q.Push(&AE{Name: "6777", Event: "lock_44"}, 5)
	q.Push(&AE{Name: "88888", Event: "lock_8"}, 12)

	rec := make(chan []*AE)
	go func() {
		q.AutoPop(rec)
	}()
	select {}
}
