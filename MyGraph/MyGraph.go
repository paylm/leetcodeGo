package main

import (
	"fmt"
)

var visited map[string]bool = make(map[string]bool) //for DFS

/***
Adjacency Matrix impl Graph
**/
type MyGraph struct {
	VRType          int //对于无权图，用 1 或 0 表示是否相邻；对于带权图，直接为权值。
	AdjMatrix       [][]int
	vexnum, edgenum int
	GraphKind       int
}

func NewMyGraph(vex int) *MyGraph {
	g := new(MyGraph)
	g.vexnum = vex
	g.AdjMatrix = make([][]int, vex)
	for i, _ := range g.AdjMatrix {
		g.AdjMatrix[i] = make([]int, vex)
	}
	return g
}

/**
* 添加连接边 , 带权值
**/
func (g *MyGraph) addEdge(x, y, VRType int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("drawEdge:%v", err)
		}
	}()
	if g.AdjMatrix == nil {
		return
	}
	g.AdjMatrix[x][y] = VRType
	g.edgenum++
}

func (g *MyGraph) findFistAdjVex() (x, y int) {
	x, y = 0, 0
	for x = 0; x < len(g.AdjMatrix); x++ {
		for y = 0; y < len(g.AdjMatrix[x]); y++ {
			if g.AdjMatrix[x][y] != 0 {
				return
			}
		}
	}
	return -1, -1
}

func (g *MyGraph) nextAdjVex(x, y int) (int, int) {

	for i := 0; i < len(g.AdjMatrix[y]); i++ {
		if g.AdjMatrix[y][i] != 0 {
			return y, i
		}
	}
	return -1, -1
}

func (g *MyGraph) show() {
	for _, v := range g.AdjMatrix {
		fmt.Println(v)
	}
}

//返回可行路径
func (g *MyGraph) DFS(start, target int) []*Point {
	routes := make([]*Point, 1)
	visDfs := make(map[int]*Point)
	var stack Stack
	stack = NewMyStack()
	p := NewPoint(start)
	visDfs[start] = p
	stack.Push(start)
	for {
		if stack.empty() {
			break
		}
		err, v := stack.Pop()
		if err != nil {
			break
		}
		zR := visDfs[v]
		if v == target {
			//fmt.Println("find a pat")
			routes = append(routes, zR)
		}
		for i := 0; i < len(g.AdjMatrix[v]); i++ {
			if g.AdjMatrix[v][i] != 0 {
				pNext := NewPoint(i)
				pNext.Prev = zR
				visDfs[i] = pNext
				stack.Push(i)
			}
		}

		delete(visDfs, v)
	}
	return routes
}

func (g *MyGraph) DFSearchDep(x, y, l int) {
	for j := y; j < len(g.AdjMatrix); j++ {
		if g.AdjMatrix[x][j] != 0 {
			vet := fmt.Sprintf("%d-%d", x, j)
			if visited[vet] != true {
				fmt.Printf("-> <%s>", vet)
				visited[vet] = true
				g.DFSearchDep(j, 0, l+1)
			}
			//输出路径可视化
			if visited[vet] == true {
				fmt.Println()
				delete(visited, vet) //out stack
				for k := 0; k < l; k++ {
					fmt.Printf("\t")
				}
				//fmt.Printf(" -> ")
			}
		}
	}
}

//找到点节F(5) ,EXIT
func (g *MyGraph) BFS(start, target int) *Point {
	var zR *Point
	visBfs := make(map[int]*Point) //for BFS
	var queue MyQueue
	queue = NewLQueue()
	p := NewPoint(start)
	visBfs[start] = p
	queue.push(start)
	for {
		if queue.empty() {
			break
		}
		err, v := queue.pop()
		if err != nil {
			break
		}
		zR = visBfs[v]
		fmt.Printf("-> %d ", v)
		if v == target {
			fmt.Printf("\n find to target :%d \n", target)
			break
		}
		//time.Sleep(time.Second * 3)
		for i := 0; i < len(g.AdjMatrix[v]); i++ {
			if g.AdjMatrix[v][i] != 0 && visBfs[i] == nil {
				//fmt.Println(i)
				iZr := NewPoint(i)
				iZr.Prev = zR
				queue.push(i)
				visBfs[i] = iZr
			}
		}
	}
	return zR
}

func (g *MyGraph) BFSlevel(start, target int) int {

	vistlevel := make(map[int]int) //for BFS
	var queue MyQueue
	queue = NewLQueue()
	vistlevel[start] = 0
	maxLevel := 0
	queue.push(start)
	for {
		if queue.empty() {
			break
		}
		err, v := queue.pop()
		if err != nil {
			break
		}
		zrlevel := vistlevel[v]
		fmt.Printf("-> %d ", v)
		if zrlevel > maxLevel {
			maxLevel = zrlevel
		}
		if v == target {
			fmt.Printf("\n find to target :%d \n", target)
			//		break
		}
		//time.Sleep(time.Second * 3)
		for i := 0; i < len(g.AdjMatrix[v]); i++ {
			if g.AdjMatrix[v][i] != 0 {
				if _, ok := vistlevel[i]; !ok {
					//fmt.Println(i)
					vistlevel[i] = zrlevel + 1
					queue.push(i)
				}
			}
		}
	}

	for k, v := range vistlevel {
		fmt.Printf("%d => %d\n", k, v)
	}
	return maxLevel
}
