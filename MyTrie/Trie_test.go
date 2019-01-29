package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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

func TestTire_Insert(t *testing.T) {
	//test string from txt
	if contents, err := ioutil.ReadFile("word.txt"); err == nil {
		//fmt.Printf("content:%v\n", string(contents))
		text := strings.Replace(string(contents), "\n", "", -1)
		text = strings.Replace(string(contents), ",", " ", -1)
		strs := strings.ToLower(text)
		strArr := strings.Split(strs, " ")
		tw := NewTrie("")
		for _, v := range strArr {
			Insert(tw, v)
		}

		t.Logf("total word :%d\n", countWord(tw))
		for _, tag := range []string{"th", "da", "o"} {
			res := Search(tw, tag)

			if len(res) == 0 {
				t.Errorf("test fail , not found word %v\n", tag)
			} else {
				t.Logf("test pass , foud to words(%s) :%v\n", tag, res)
			}
		}
	} else {
		t.Errorf("load text from word.txt , throw err:%v\n", err)
	}
}

func BenchmarkTrie_Insert(b *testing.B) {
	if contents, err := ioutil.ReadFile("word.txt"); err == nil {
		//fmt.Printf("content:%v\n", string(contents))
		text := strings.Replace(string(contents), "\n", "", -1)
		text = strings.Replace(string(contents), ",", " ", -1)
		strs := strings.ToLower(text)
		strArr := strings.Split(strs, " ")
		tw := NewTrie("")
		for i := 0; i < b.N; i++ {
			for _, v := range strArr {
				Insert(tw, v)
			}
		}
	}
}

func Test_SearchWord(t *testing.T) {
	tw := NewTrie("")
	datas := []string{
		"keys", "hello", "world", "hi", "word", "work", "words",
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

func BenchmarkTrie_Search(b *testing.B) {
	if contents, err := ioutil.ReadFile("word.txt"); err == nil {
		//fmt.Printf("content:%v\n", string(contents))
		b.StopTimer()
		text := strings.Replace(string(contents), "\n", "", -1)
		text = strings.Replace(string(contents), ",", " ", -1)
		strs := strings.ToLower(text)
		strArr := strings.Split(strs, " ")
		tw := NewTrie("")
		for _, v := range strArr {
			Insert(tw, v)
		}
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			Search(tw, "wo")
			Search(tw, "th")
			Search(tw, "a")
		}
	}
}

func Test_DelWord(t *testing.T) {
	t1 := NewTrie("")
	datas := []string{
		"keys", "hello", "world", "hi", "word", "work", "words",
	}
	for _, s := range datas {
		Insert(t1, s)
	}

	fmt.Printf("del word:%v,ret :%v", "words", Del(t1, "words"))

	case1 := Search(t1, "wo")
	if len(case1) == 0 {
		t.Errorf("test fail , not found word %v\n", "wo")
	} else {
		t.Logf("test pass , foud to words :%v\n", case1)
	}
}

func Test_CommonPrefix(t *testing.T) {
	testcase := make(map[string][]string)
	testcase["gee"] = []string{"geeksforgeeks", "geeks", "geek", "geezer"}
	testcase["wo"] = []string{"work", "working", "word", "world", "wood"}

	for k, v := range testcase {
		res := commonPrefix(v)

		if res == k {
			t.Logf("test pass,%v commonPrefix is %v\n", v, k)
		} else {
			t.Errorf("test fail,%v commonPrefix is %v,but return %v\n", v, k, res)
		}
	}
}

func Test_CountWord(t *testing.T) {
	testcase := make(map[int][]string)
	testcase[4] = []string{"geeksforgeeks", "geeks", "geek", "geezer"}
	testcase[5] = []string{"work", "working", "word", "world", "wood"}

	for k, v := range testcase {
		tw := NewTrie("")
		for _, s := range v {
			Insert(tw, s)
		}
		count := countWord(tw)

		if count == k {
			t.Logf("test pass,%v has  %d word \n", v, k)
		} else {
			t.Errorf("test fail,%v has %d word ,but return %v\n", v, k, count)
		}
	}
}
