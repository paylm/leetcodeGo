package main

import (
	"fmt"
)

/***

  Graph :
   0 -> 1  -> 2      //2 also has edge to 3
   \     \   /
    > 3 -> 4<
       \   /
	    >5

	A B C D E F
  A 0 1 0 1 0 0
  B 0 0 1 0 1 0
  C 0 0 0 1 1 0
  D 0 0 0 0 1 1
  E 0 0 0 0 0 1
**/

var ids byte = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}

func testStack() {
	s := NewMyStack()
	a := []int{1, 3, 5, 7, 9}
	for _, x := range a {
		s.Push(x)
	}

	i := len(a) - 1
	for {
		if i < 0 {
			break
		}
		_, v := s.Pop()
		if v != a[i] {
			fmt.Println("push and pop err")
		}
		i--
		fmt.Printf("pop : %d\n", v)
	}
}

func testInput() {
	var size int
	var p1, p2 int
	fmt.Println("input graph vet size")
	fmt.Scanln(&size)
	g := NewMyGraph(size)
	for {
		fmt.Scanln(&p1, &p2)
		fmt.Printf("p1=>%d,p2=>%d\n", p1, p2)
		if p1 >= size || p2 >= size || p1 == -1 {
			break
		}

		g.addEdge(p1, p2, 1)
	}
	g.show()

	fmt.Println("DFS SEARCH")
	//testStack()
	paths := g.DFS(0, size-1)
	for _, v := range paths {
		v.traversePath()
	}
}

func main() {
	fmt.Println("vim-go")
	g := NewMyGraph(6)
	//drawImg
	g.addEdge(0, 1, 1)
	g.addEdge(0, 3, 1)
	g.addEdge(1, 2, 1)
	g.addEdge(1, 4, 1)
	g.addEdge(2, 3, 1)
	g.addEdge(2, 4, 1)
	g.addEdge(3, 4, 1)
	g.addEdge(3, 5, 1)
	g.addEdge(4, 5, 1)
	g.show()
	x, y := g.findFistAdjVex()
	fmt.Println(x, y)
	//g.DFS(x, y)
	//g.DFSearchDep(x, y, 0)
	//target F(5)
	fmt.Println("BFS SEARCH")
	bZR := g.BFS(0, 5)
	bZR.traversePath()

	fmt.Println("DFS SEARCH")
	//testStack()
	paths := g.DFS(2, 5)
	for _, v := range paths {
		v.traversePath()
	}
	fmt.Println("level tree")
	fmt.Printf("from %d -> %d , level :%d \n", 0, 5, g.BFSlevel(0, 5))

	testInput()
}
