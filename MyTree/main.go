package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := []int{0, 1, 2, 3, 4}
	//删除第i个元素
	i := 2
	a = append(a[:i], a[i+1:]...)

	var q Queue
	q = NewMyQueue()
	q.put(1)
	q.put(2)
	q.put(3)
	q.put(7)
	for {
		if q.empty() {
			break
		}
		fmt.Println(q.pop())
	}

	arr := []int{8, 4, 10, 3, 5, 1, 7}
	var root *Node
	for _, v := range arr {
		if root == nil {
			root = NewNode(v)
		} else {
			root.InsertNode(NewNode(v))
		}
	}

	travese(root)
}
