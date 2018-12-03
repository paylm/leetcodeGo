package main

type LRUCache interface{
	put(k string,v string)
	get(k string) string
}

type LRU struct {
	data map[string]inteface{}
	queue MyQueue
	size int
}

func NewLRU(cap int) *LRU{
	lru := new(LRU)
	ru.data = make(map[string]string)
	lru.size = cap
	lru.queue = NewListQueue(size)
	return lru
}

func (lru *LRU)put(k string,val string){

}

func (lru *LRU)get(k string) string{

	return "fuck"
}
