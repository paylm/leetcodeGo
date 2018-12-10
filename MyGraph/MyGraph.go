package main

import "fmt"

const (
	MAX_VERtEX_NUM int = 20
)

var visited map[string]bool = make(map[string]bool)

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

func (g *MyGraph) drawEdge(x, y, VRType int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("drawEdge:%v", err)
		}
	}()
	if g.AdjMatrix == nil {
		return
	}
	g.AdjMatrix[x][y] = VRType
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

func (g *MyGraph) DFS(x, y int) {
	for j := y; j < len(g.AdjMatrix); j++ {
		if g.AdjMatrix[x][j] != 0 {
			vet := fmt.Sprintf("%d-%d", x, j)
			if visited[vet] != true {
				fmt.Printf("<%s> -> ", vet)
				visited[vet] = true
				g.DFS(j, 0)
			}
			//out stack
			delete(visited, vet)
		}
	}
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
