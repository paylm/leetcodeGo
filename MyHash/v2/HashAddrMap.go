package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	//Used in hash2 functoin
	PRIME = 7
	//init size
	INISIZE = 100
	//Collicsion X 0.75
	COLLISX = 75
)

type HashAddrMap struct {
	TableSize int
	used      int
	data      []*HashNode
	m         *sync.RWMutex
}

type HashNode struct {
	Key        int
	Val        int
	Collicsion int
}

func NewHashNode(k, v int) *HashNode {
	n := new(HashNode)
	n.Key = k
	n.Val = v
	return n
}

func NewHashAddrMap() *HashAddrMap {
	h := new(HashAddrMap)
	h.TableSize = INISIZE
	h.data = make([]*HashNode, h.TableSize)
	h.used = 0
	h.m = new(sync.RWMutex)
	return h
}

func (h *HashAddrMap) hash1(k int) int {
	return abs(k % h.TableSize)
}

func (h *HashAddrMap) hash2(k int) int {
	return abs(PRIME - (k % PRIME))
}

func (h *HashAddrMap) isFull() bool {

	if (h.used*100)/h.TableSize > COLLISX {
		return true
	}
	return false
}

func (h *HashAddrMap) resize() {
	fmt.Printf("resize map\n")
	//fmt.Println("old data:", h.data)
	h.m.Lock()
	h.TableSize = h.TableSize * 2
	newData := make([]*HashNode, 2*h.TableSize)
	for _, v := range h.data {
		if v == nil {
			continue
		}

		i := h.hash1(v.Key)
		if newData[i] == nil {
			newData[i] = v
			continue
		}
		index2 := h.hash2(v.Key)
		j := 1 //计数器
		for {
			newIndex := (i + index2*j) % h.TableSize
			if newData[newIndex] == nil {
				newData[newIndex] = v
				//fmt.Printf("发现碰撞,原i:%d,新i:%d \n", i, newIndex)
				break
			}
			j++
		}

	}
	h.data = newData
	h.m.Unlock()
}

func (h *HashAddrMap) Insert(k, v int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Insert recover k=%d,hash1:%d,err:%s\n", k, h.hash1(k), err)
		}
	}()
	if h.isFull() {
		h.resize()
	}
	i := h.hash1(k)
	if h.data[i] == nil {
		h.data[i] = NewHashNode(k, v)
		h.used++
		return
	}
	index2 := h.hash2(k)
	j := 1 //计数器
	for {
		newIndex := (i + index2*j) % h.TableSize
		if h.data[newIndex] == nil {
			fmt.Printf("发现碰撞,k=%d,v=%d,原i:%d,新i:%d \n", k, v, i, newIndex)
			h.data[newIndex] = NewHashNode(k, v)
			break
		}
		j++
	}
	h.used++
}

/**
线性查找，
    如果找到的位置key 不等，继续往下查
	如果找到値为 nil , 说明不存在
*/
func (h *HashAddrMap) Search(k int) (error, int) {
	i := h.hash1(k)
	if h.data[i] != nil && h.data[i].Key == k {
		return nil, h.data[i].Val
	}
	index2 := h.hash2(k)
	j := 1
	for {
		newIndex := (i + index2*j) % h.TableSize
		if h.data[newIndex] == nil {
			return errors.New(fmt.Sprintf("%d not found", k)), -1
		}
		if h.data[newIndex].Key == k {
			return nil, h.data[newIndex].Val
		}
		j++
	}
	return nil, -1
}

func (h *HashAddrMap) Del(k int) {

	i := h.hash1(k)
	if h.data[i] != nil && h.data[i].Key == k {
		h.data[i] = nil
		h.used--
		return
	}
	index2 := h.hash2(k)
	j := 1
	for {
		newIndex := (i + index2*j) % h.TableSize
		if h.data[newIndex] == nil {
			break
		}
		if h.data[newIndex].Key == k {
			h.data[newIndex] = nil
			h.used--
			break
		}
		j++
	}
}

func (h *HashAddrMap) HmPrint() {
	for i, v := range h.data {
		if v == nil {
			continue
		}
		fmt.Printf("i:=%d,k=%d,v=%d\n", i, v.Key, v.Val)
	}
}
