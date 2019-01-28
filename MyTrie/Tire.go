package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	Key    string
	nexts  map[string]*Trie //子元素指针s
	size   int              //子元素数目
	isWord bool             //是否为完整单词
}

func NewTrie(k string) *Trie {
	tw := new(Trie)
	tw.Key = k
	tw.nexts = make(map[string]*Trie)
	tw.size = 0
	return tw
}

func Insert(t *Trie, word string) {
	wsArr := strings.Split(word, "")
	if len(wsArr) == 0 {
		return
	}
	cT := t
	for _, s := range wsArr {
		c, ok := cT.nexts[s]
		if !ok {
			c = NewTrie(s)
			cT.nexts[s] = c
			cT.size++
		}
		cT = c
	}
	cT.isWord = true
}

func Search(t *Trie, word string) []string {
	wsArr := strings.Split(word, "")
	res := []string{}
	if len(wsArr) == 0 {
		return res
	}
	cT := t
	for _, s := range wsArr {
		c, ok := cT.nexts[s]
		if !ok {
			return res
		}
		cT = c
	}
	if cT.isWord {
		res = append(res, word)
		return res
	}
	//seaerch  next word
	//fmt.Println("find next word")
	return recurrence_Search(cT, word, res)
}

func recurrence_Search(t *Trie, prex string, res []string) []string {

	if t == nil {
		return nil
	}
	for _, s := range t.nexts {
		ts := fmt.Sprintf("%s%s", prex, s.Key)
		//r1 := recurrence_Search(s, ts, res)
		if s.isWord {
			res = append(res, ts)
			res = append(res, recurrence_Search(s, ts, nil)...)
		} else {
			res = append(res, recurrence_Search(s, ts, res)...)
		}
	}
	return res
}

func Del(t *Trie, word string) {
	wsArr := strings.Split(word, "")
	if len(wsArr) == 0 {
		return
	}
	list := make([]*Trie, len(wsArr))
	cT := t
	for i, s := range wsArr {
		c, ok := cT.nexts[s]
		if !ok {
			return
		}
		cT = c
		list[i] = cT
	}

	if cT.isWord != true {
		return
	}
	for i := len(wsArr) - 2; i > 0; i-- {
		if list[i].isWord {
			//fmt.Printf("del %v\n", list[i])
			delete(list[i].nexts, list[i+1].Key)
		}
	}

}