package main

import "fmt"

func main() {
	data := []string{"the", "joke", "answer", "bbs", "day", "cool", "any", "bye", "content", "their", "according", "there"}
	tw := NewTrie("root")
	for _, ws := range data {
		Insert(tw, ws)
	}
	fmt.Println("---Display---")
	display(tw, "")
	fmt.Println("---sortDisplay---")
	sortDisplay(tw, "")

	testcase := []string{"geek", "geek0", "geek1", "geek", "geek2"}
	res := suffixWord(testcase)
	fmt.Printf("suffixWord :%v\n", res)
}
