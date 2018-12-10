package main

import "fmt"

/***

  Graph :
   A -> B  -> C       //C also has edge to D
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

func main() {
	fmt.Println("vim-go")
	g := NewMyGraph(6)
	//drawImg
	g.drawEdge(0, 1, 1)
	g.drawEdge(0, 3, 1)
	g.drawEdge(1, 2, 1)
	g.drawEdge(1, 4, 1)
	g.drawEdge(2, 3, 1)
	g.drawEdge(2, 4, 1)
	g.drawEdge(3, 4, 1)
	g.drawEdge(3, 5, 1)
	g.drawEdge(4, 5, 1)
	g.show()
	x, y := g.findFistAdjVex()
	fmt.Println(x, y)
	//g.DFS(x, y)
	g.DFSearchDep(x, y, 0)
}
