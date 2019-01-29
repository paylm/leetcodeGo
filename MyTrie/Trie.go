package main

import (
	"errors"
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
	return recurrence_Search(cT, word)
}

func recurrence_Search(t *Trie, prex string) []string {

	if t == nil {
		return nil
	}
	res := []string{}
	for _, s := range t.nexts {
		ts := fmt.Sprintf("%s%s", prex, s.Key)
		//r1 := recurrence_Search(s, ts, res)

		res = append(res, recurrence_Search(s, ts)...)
		if s.isWord {
			res = append(res, ts)
		}
	}
	return res
}

func Del(t *Trie, word string) error {
	wsArr := strings.Split(word, "")
	if len(wsArr) == 0 {
		return nil
	}
	list := make([]*Trie, len(wsArr))
	cT := t
	for i, s := range wsArr {
		c, ok := cT.nexts[s]
		if !ok {
			return nil
		}
		cT = c
		list[i] = cT
	}

	if cT.isWord != true {
		return nil
	}

	if cT.size > 0 {
		return errors.New(fmt.Sprintf("can't not delete %v,becase other world dep it\n", word))
	}

	for i := len(wsArr) - 2; i > 0; i-- {
		if list[i].isWord {
			//fmt.Printf("del %v\n", list[i])
			delete(list[i].nexts, list[i+1].Key)
			list[i].size--
		}
	}
	return nil
}

//公共字符的公共前缀
func commonPrefix(strs []string) string {
	//step 1 . add all string to trie
	t := NewTrie("")
	for _, s := range strs {
		Insert(t, s)
	}
	//step 2 : find the childNode which has more than 1 childen
	ct := t
	reStrArr := []string{}
	for {
		if ct == nil || ct.size > 1 {
			break
		}
		for _, mp := range ct.nexts {
			ct = mp
		}
		reStrArr = append(reStrArr, ct.Key)
	}
	return strings.Join(reStrArr, "")
}

func countWord(t *Trie) int {
	count := 0
	if t.isWord {
		count++
	}
	for _, cT := range t.nexts {
		count = count + countWord(cT)
	}
	return count
}
