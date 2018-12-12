package main

import "fmt"

/**
Adjacency List:
An array of lists is used. Size of the array is equal to the number of vertices. Let the array be array[]. An entry array[i] represents the list of vertices adjacent to the ith vertex
**/
type LsGraph struct {
	Vertex  int
	Edge    int
	AdjList []*AdjListNode
}

type AdjListNode struct {
	Val  int
	Next *AdjListNode
}

func NewAdjListNode(v int) *AdjListNode {
	adj := new(AdjListNode)
	adj.Val = v
	return adj
}

func NewLsGraph(s int) *LsGraph {
	lg := new(LsGraph)
	lg.Vertex = s
	lg.AdjList = make([]*AdjListNode, lg.Vertex)
	return lg
}

// Add an edge from src to dest.  A new node is
// added to the adjacency list of src.  The node
// is added at the begining
func (g *LsGraph) addEdge(src, dst int) {
	d := NewAdjListNode(dst)
	d.Next = g.AdjList[src]
	g.AdjList[src] = d
}

func (g *LsGraph) BFS(src, dst int) {
	visBfs := make(map[int]bool)
	queue := NewLQueue()
	queue.push(src)
	visBfs[src] = true
	for {
		if queue.empty() {
			break
		}
		err, v := queue.pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%d->", v)
		if v == dst {
			fmt.Printf("\nfind the dst :%d\n", dst)
			break
		}
		c := g.AdjList[v]
		//fmt.Println(c)
		for {
			if c == nil {
				break
			}
			if visBfs[c.Val] != true {
				queue.push(c.Val)
				visBfs[c.Val] = true
			}
			c = c.Next
		}
	}
}

func (g *LsGraph) show() {
	for i, gls := range g.AdjList {
		c := gls
		fmt.Printf("Adjacency list of vertex %d\n", i)
		for {
			if c == nil {
				fmt.Println()
				break
			}
			fmt.Printf("->%d", c.Val)
			c = c.Next
		}
	}
}
