package main

import (
	"testing"
)

func Test_InsertWord(t *testing.T) {
	tw := NewTrie("")
	datas := []string{
		"keys", "hello", "world", "hi", "word",
	}
	for _, s := range datas {
		Insert(tw, s)
	}

	if tw.size == 0 {
		t.Errorf("test fail,after Insert word , trie size len is 0")
	} else {
		t.Log("test pass")
	}
}

func Test_SearchWord(t *testing.T) {
	tw := NewTrie("")
	datas := []string{
		"keys", "hello", "world", "hi", "word",
	}
	for _, s := range datas {
		Insert(tw, s)
	}

	res := Search(tw, "hi")
	if len(res) == 0 {
		t.Errorf("test fail , not found word %v\n", "hi")
	} else {
		t.Logf("test pass , foud to words :%v\n", res)
	}

	case1 := Search(tw, "wo")
	if len(case1) == 0 {
		t.Errorf("test fail , not found word %v\n", "wo")
	} else {
		t.Logf("test pass , foud to words :%v\n", case1)
	}
}
