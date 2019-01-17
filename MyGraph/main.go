package main

import (
	"fmt"
)

type Point struct {
	Val   int
	Space int
	Prev  *Point
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
		fmt.Printf("->(%d) %d ", c.Space, c.Val)
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

/**
判断两个数是是否只有一位不同
eg: 1432   1232 true
    1432   1331 false
	1432   1432 false
**/
func compareSingle(a, b int) bool {
	if a == b {
		return false
	}
	var x, y int
	if a > b {
		x = a
		y = b
	} else {
		x = b
		y = a
	}
	i := 10
	n := 0 //x,y相差位数
	for {
		if x < 1 {
			break
		}
		k1 := x % i
		k2 := y % i
		//	fmt.Println(k1, k2)
		if k1 != k2 {
			n++
		}
		x = x / 10
		y = y / 10
	}
	//fmt.Printf("%d compare %d diff byte  n :%d , i :%d \n", a, b, n, i)
	if n == 1 {
		return true
	} else {
		return false
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

func testdijInput() {
	var size int
	var p1, p2, vd int
	fmt.Println("input graph vet size")
	fmt.Scanln(&size)
	g := NewMyGraph(size)
	for {
		fmt.Scanln(&p1, &p2, &vd)
		fmt.Printf("p1=>%d,p2=>%d,vd=>%d\n", p1, p2, vd)
		if p1 >= size || p2 >= size || p1 == -1 {
			break
		}

		g.addCyEdge(p1, p2, vd)
	}
	g.show()

	fmt.Println("dijkstra SEARCH")
	var target int
	fmt.Println("please input target node")
	for {
		fmt.Scanln(&target)
		if target < len(g.AdjMatrix) {
			fmt.Printf("src:%d,target:%d\n", 0, target)
			break
		}
		fmt.Println("target not in graph , please input again")
	}
	//testStack()
	res := g.dijkstra(0, len(g.AdjMatrix)-1)
	fmt.Println(res)
}

func testListGraph() {

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

//数组生成路径
//https://www.geeksforgeeks.org/minimum-initial-vertices-traverse-whole-matrix-given-conditions/
func testLsGraph(a []int) {
	fmt.Println(a)
	l := len(a)
	g3 := NewLsGraph(l)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			//		fmt.Printf("i:%d,j=%d\n", a[i], a[j])
			if compareSingle(a[i], a[j]) {
				//			fmt.Println("compareSingle true")
				g3.addEdge(a[i], a[j])
				//g3.addEdge(a[j], a[i])
			}
		}
	}
	g3.show()
	r := g3.BFS(a[0], a[l-1])
	if r != nil {
		fmt.Printf("found %d to %d\n", a[0], a[l-1])
		r.traversePath()
	}
	r2 := g3.BFS(a[2], a[l-2])
	if r != nil {
		fmt.Printf("found %d to %d\n", a[2], a[l-2])
		r2.traversePath()
	}

	rs := g3.DFS(a[0], a[l-1])
	fmt.Println(rs)
	for _, v := range rs {
		fmt.Println("DFS PATH")
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
		v.traverseRole(role)
	}
	fmt.Println("level tree")
	fmt.Printf("from %d -> %d , level :%d \n", 0, 5, g.BFSlevel(0, 5))

	//testInput()
	//	fmt.Println(compareSingle(1234, 1244))
	//	fmt.Println(compareSingle(1224, 1334))
	//	fmt.Println(compareSingle(1234, 1234))
	//	fmt.Println(compareSingle(1733, 3733))
	testLsGraph([]int{1033, 1733, 3733, 3739, 3779, 8779, 8179})
	fmt.Println("dijkstra search ")
	//testdijInput()
	g1 := NewMyGraph(6)
	g1.addCyEdge(0, 1, 7)
	g1.addCyEdge(0, 2, 9)
	g1.addCyEdge(0, 5, 14)
	g1.addCyEdge(1, 2, 10)
	g1.addCyEdge(1, 3, 15)
	g1.addCyEdge(2, 3, 11)
	g1.addCyEdge(2, 5, 2)
	g1.addCyEdge(3, 4, 6)
	g1.addCyEdge(4, 5, 9)
	g1.show()
	res := g1.dijkstra(0, 4)
	fmt.Println("dijkstra res : ", res)
	r1 := g1.dijkstraPath(0, 4)
	r1.traversePath()

	r2 := g1.dijkstraPath(0, 5)
	r2.traversePath()

	r3 := g1.dijkstrQue(0, 4)
	fmt.Println("dijkstrQue res :", r3)
}
