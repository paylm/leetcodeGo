package main

import "fmt"

//路径节点
type Node struct {
	Val  string
	Prev *Node
}

func NewNode(v string) *Node {
	n := new(Node)
	n.Val = v
	return n
}

//返回子节点
func addNextNode(n *Node, v string) *Node {
	sNode := NewNode(v)
	if n == nil {
		return sNode
	}
	sNode.Prev = n

	return sNode
}

func traverse(last *Node) {
	if last == nil {
		return
	}
	c := last
	for {
		if c == nil {
			break
		}
		fmt.Printf("%s->", c.Val)
		c = c.Prev
	}
	fmt.Println()
}

/**测试数据
A { B C D }
B { B1 B2 B3 }
C { C1 ,C2, D }
D { B1 , C1 ,D1 }
D1 { Z }
//起点为A，终点为Z

**/
func createRouter() map[string][]string {
	r := make(map[string][]string)
	r["root"] = []string{"A"}
	r["A"] = []string{"B", "C", "D"}
	r["B"] = []string{"B1", "B2", "B3"}
	r["C"] = []string{"C1", "C2", "D"}
	r["D"] = []string{"B1", "C1", "D1"}
	r["D1"] = []string{"Z"}
	return r
}

//通过dfs 算法从route找到终点
func DFS(route map[string][]string, searched map[string]bool, root string, dst string) bool {
	//fmt.Println("stack info :", searched)
	searchQueue := []string{root}
	if len(searchQueue) == 0 {
		return false
	}
	searched[root] = true
	defer func() {
		//	fmt.Println("out stack:", root)
		delete(searched, root)
	}()
	for {
		if len(searchQueue) < 1 {
			break
		}
		path := searchQueue[0]
		searchQueue = searchQueue[1:]
		//fmt.Printf("current search point %s\n", path)
		if path == dst {
			//	fmt.Println("已找到节点:", dst)
			return true
		}

		for _, x := range route[path] {
			//fmt.Println(x, searched[x])
			if _, ok := searched[x]; !ok {
				searchQueue = append(searchQueue, x)
				searched[x] = true
				DFS(route, searched, x, dst)
			} //else {
			//fmt.Printf("%s has checked , skip\n", x)
			//}
		}
	}
	return false
}

//通过bfs 算法从route找到终点
func BFSwithPath(route map[string][]string, root string, dst string) (bool, *Node) {
	fmt.Println("查找节点:", dst)
	searchQueue := []string{root}
	var zR *Node
	if len(searchQueue) == 0 {
		return false, zR
	}
	searched := make(map[string]*Node)
	pNode := NewNode(root)
	searched[root] = pNode
	for {
		if len(searchQueue) < 1 {
			break
		}
		path := searchQueue[0]
		zR = searched[path]
		searchQueue = searchQueue[1:]
		//fmt.Printf("current search point %s\n", path)
		if path == dst {
			fmt.Println("已找到节点:", dst)
			return true, zR
		}

		for _, x := range route[path] {
			//fmt.Println(x, searched[x])
			if _, ok := searched[x]; !ok {
				searchQueue = append(searchQueue, x)
				xZr := addNextNode(zR, x)
				searched[x] = xZr
			} //else {
			//	fmt.Printf("%s has checked , skip\n", x)
			//}
			//searched[x] = true
		}
		//fmt.Println("queue:", searchQueue)
	}
	return false, nil
}

//通过dfs 算法找到所有路径
func DFSwithPath(route map[string][]string, root string, dst string) (bool, []*Node) {
	searchStack := []string{root}
	visited := make(map[string]*Node)
	var p string
	pNode := NewNode(root)
	xrs := make([]*Node, 1)
	visited[root] = pNode
	for {
		if len(searchStack) < 1 {
			break
		}
		p = searchStack[0]
		zR := visited[p]
		//	fmt.Printf("-> %s ", p)
		if p == dst {
			//		fmt.Printf("\n found to dst :%s\n", p)
			//return true, zR
			xrs = append(xrs, zR)
		}

		searchStack = append(searchStack[1:])
		for _, x := range route[p] {
			if _, ok := visited[x]; !ok {
				searchStack = append([]string{x}, searchStack...)
				xZr := addNextNode(zR, x)
				visited[x] = xZr
			}
		}
		delete(visited, p)
	}
	if len(xrs) > 0 {
		return true, xrs
	} else {
		return false, nil
	}
}

func main() {
	fmt.Println("vim-go")
	r := createRouter()
	searched := make(map[string]bool)
	fmt.Println(DFS(r, searched, "A", "Z"))
	//fmt.Println(DFS(r, "A", "C1"))
	ok, zR := BFSwithPath(r, "A", "Z")
	fmt.Println(ok, zR)
	if ok {
		traverse(zR)
	}
	//	a := []string{"abc", "ccc"}
	//	c := append([]string{"bbb"}, a...)
	//	fmt.Println(c)
	//	fmt.Println(append(c[1:]))
	fmt.Println("----find all path pass by dfs  ----")
	ok1, zRs := DFSwithPath(r, "A", "Z")
	if ok1 {
		for _, rs := range zRs {
			traverse(rs)
		}
	}
}
