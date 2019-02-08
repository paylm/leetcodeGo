package main

import (
	"errors"
	"fmt"
	"sort"
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

//利用前缀树进行同名文件处理(abc,abb,abc,abc) to (abc,bbb,abc1,abc2)
//https://www.geeksforgeeks.org/program-for-assigning-usernames-using-trie/
func suffixWord(strs []string) []string {
	tw := NewTrie("root")
	//insert to trie , if found exist word , rename it to s1,s2,s3 ....
	for _, w := range strs {
		alpw := strings.Split(w, "")
		alTw := tw
		for _, cw := range alpw {
			alItem, ok := alTw.nexts[cw]
			if !ok {
				alItem = NewTrie(cw)
				alTw.size++
				alTw.nexts[cw] = alItem
				alTw = alItem
			} else {
				alTw = alItem
			}
		}
		//fmt.Printf("end word:%s\n", alTw.Key)
		if alTw.isWord {
			temp := NewTrie(fmt.Sprintf("%d", alTw.size))
			temp.isWord = true
			//fmt.Printf("aready exist old value,insert duble item:%v\n", temp)
			alTw.nexts[fmt.Sprintf("%d", alTw.size)] = temp
			alTw.size++
		} else {
			alTw.isWord = true
		}
	}
	return recurrence_Search(tw, "")
}

func display(t *Trie, prefix string) {
	if t == nil {
		return
	}

	for _, s := range t.nexts {
		sprefix := fmt.Sprintf("%s%s", prefix, s.Key)
		if s.isWord {
			fmt.Println(sprefix)
		}
		display(s, sprefix)
	}
}

//output tire with format like a->b->c...z
func sortDisplay(t *Trie, prefix string) {

	if t == nil {
		return
	}
	newMp := []string{}
	for _, w := range t.nexts {
		newMp = append(newMp, w.Key)
	}
	sort.Strings(newMp)

	for _, s := range newMp {
		sprefix := fmt.Sprintf("%s%s", prefix, s)
		st, _ := t.nexts[s]
		if st.isWord {
			fmt.Println(sprefix)
		}
		display(st, sprefix)
	}
}

//对字符串和各字符进行排序输出
func sortApl(word string) []string {

	s := strings.Split(word, "")
	sort.Strings(s)
	return s
}
