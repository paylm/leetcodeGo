package main

import (
	"fmt"
)

type Point struct {
	Val  int
	Prev *Point
}

func NewPoint(v int) *Point {
	p := new(Point)
	p.Val = v
	return p
}

func (p *Point) traversePath() {
	c := p
	for {
		if c == nil {
			break
		}
		fmt.Printf("-> %d ", c.Val)
		c = c.Prev
	}
	fmt.Println()
}

func (p *Point) traverseRole(r []string) {

	c := p
	for {
		if c == nil {
			break
		}
		fmt.Printf("-> %s ", r[c.Val])
		c = c.Prev
	}
	fmt.Println()

}

/***

  Graph :
   A -> B  -> C      //B also has edge to C
   \     \   /
    > D -> E<
       \   /
	    >F

	A B C D E F
  A 0 1 0 1 0 0
  B 0 0 1 0 1 0
  C 0 0 0 1 1 0
  D 0 0 0 0 1 1
  E 0 0 0 0 0 1
**/

var role = []string{"A", "B", "C", "D", "E", "F", "G", "H"}

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
		v.traverseRole(role)
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
		v.traverseRole(role)
	}
	fmt.Println("level tree")
	fmt.Printf("from %d -> %d , level :%d \n", 0, 5, g.BFSlevel(0, 5))

	//testInput()
	g2 := NewLsGraph(5)
	g2.addEdge(0, 1)
	g2.addEdge(0, 2)
	g2.addEdge(1, 2)
	g2.addEdge(1, 3)
	g2.addEdge(2, 3)
	g2.addEdge(3, 4)
	g2.addEdge(4, 1)
	g2.show()
	g2.BFS(0, 4)
}
