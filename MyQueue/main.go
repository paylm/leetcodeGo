package main

import (
	"fmt"
)

func test() {
	q1 := NewQueue()
	q1.push(1)
	q1.push(3)
	q1.push(5)
	q1.push(6)
	q1.show()
	fmt.Println("-----取队列测试------")
	fmt.Println(q1.pop())
	fmt.Println(q1.pop())
	q1.show()
	fmt.Println("-----再次插入------")
	q1.push(10)
	q1.push(12)
	q1.push(17)
	q1.show()
	fmt.Println(q1.pop())
	fmt.Println(q1.pop())
	fmt.Println(q1.pop())
	q1.show()
}

func main() {
	fmt.Println("MyQueue")
	fmt.Println("-------队列q2-------")
	q2 := NewQueue()
	q2.push(1)
	q2.push(3)
	q2.push(5)
	q2.push(6)
	q2.show()
	fmt.Println("-----取队列测试------")
	fmt.Println(q2.pop())
	fmt.Println(q2.pop())
	q2.show()
	fmt.Println("-----再次插入------")
	q2.push(10)
	q2.push(12)
	q2.push(17)
	q2.show()
	fmt.Println(q2.pop())
	fmt.Println(q2.pop())
	fmt.Println(q2.pop())
	q2.show()

	pq := NewPriorityQueue(20)
	pq.push(2)
	pq.push(5)
	pq.push(1)
	pq.push(8)
	pq.push(6)
	pq.push(9)
	pq.show()
	for {
		if pq.empty() {
			break
		}
		x, _ := pq.pop()
		fmt.Println("pop val:", x)
		pq.show()
	}
}
