package main

import (
	"fmt"
)
func main(){
	fmt.Println("test")
	a := 1
	b := 2
	swap(&a,&b)
	fmt.Println(a,b)

	mp := NewMinHeap(10)
	mp.Insert(14)
	mp.Insert(5)
	mp.Insert(2)
	mp.Insert(8)
	mp.Insert(1)
	mp.Insert(9)
	mp.Insert(10)
	mp.Insert(12)
	mp.Insert(66)
	mp.Insert(34)
	mp.GetMin()
	fmt.Println(mp)

	//swap(&(mp.harr[0]),&(mp.harr[1]))
	//fmt.Println(mp)

	mp.Remove(4)
	fmt.Println(mp)
	root,_ := mp.PopRoot()
	fmt.Println("root:",root)
	fmt.Println(mp)
}