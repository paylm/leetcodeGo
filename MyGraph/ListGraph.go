package main

import "fmt"

/**
Adjacency List:
An array of lists is used. Size of the array is equal to the number of vertices. Let the array be array[]. An entry array[i] represents the list of vertices adjacent to the ith vertex
**/
type LsGraph struct {
	Vertex  int
	Edge    int
	AdjList map[int]*AdjListNode
}

type s struct{}

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
	lg.AdjList = make(map[int]*AdjListNode)
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

func (g *LsGraph) BFS(src, dst int) *Point {
	visBfs := make(map[int]*Point)
	queue := NewLQueue()
	queue.push(src)
	zR := NewPoint(src)
	visBfs[src] = zR
	for {
		if queue.empty() {
			break
		}
		err, v := queue.pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		zR = visBfs[v]
		//fmt.Printf("%d->", v)
		if v == dst {
			fmt.Printf("\nfind the dst :%d\n", dst)
			return zR
			//break
		}
		c := g.AdjList[v]
		//fmt.Println(c)
		for {
			if c == nil {
				break
			}
			if visBfs[c.Val] == nil {
				queue.push(c.Val)
				xZr := NewPoint(c.Val)
				xZr.Prev = zR
				visBfs[c.Val] = xZr
			}
			c = c.Next
		}
	}
	return nil
}

func (g *LsGraph) DFS(src, dst int) []*Point {
	rs := make([]*Point, 1)
	visDfs := make(map[int]*Point)
	stack := NewMyStack()
	stack.Push(src)
	visDfs[src] = NewPoint(src)
	for {
		if stack.empty() {
			break
		}
		err, v := stack.Pop()
		if err != nil {
			break
		}
		zR := visDfs[v]
		if v == dst {
			fmt.Println("router :", zR)
			rs = append(rs, zR)
		}
		c := g.AdjList[v]
		for {
			if c == nil {
				break
			}
			if visDfs[c.Val] == nil {
				xZr := NewPoint(c.Val)
				xZr.Prev = zR
				visDfs[c.Val] = xZr
				stack.Push(c.Val)
			}
			c = c.Next
		}
		//	fmt.Printf("%d remove from visDfs", v)
		delete(visDfs, v)
	}

	return rs
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
