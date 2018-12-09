package main

import "fmt"

type HashMap interface {
	Insert(k, v int)
	Search(k int) (error, int)
	Del(k int)
	HmPrint()
}

const MAX int = 200

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func insert(a *[MAX][2]int, k int) {
	if k > 0 {
		a[k][0] = 1
	} else {
		a[abs(k)][1] = 1
	}
}

func search(a *[MAX][2]int, k int) bool {
	if k >= 0 {
		if a[k][0] == 1 {
			return true
		} else {
			return false
		}
	} else {
		ak := abs(k)

		if a[ak][0] == 1 {
			return true
		} else {
			return false
		}
	}

}

func test1() {
	fmt.Println("test hash")
	has := [MAX][2]int{}

	arr := []int{1, 3, 5, 0, -1, -4, -7, -3}
	for _, v := range arr {
		insert(&has, v)
	}
	//fmt.Println(has)
	fmt.Println(3, search(&has, 3))
	fmt.Println(5, search(&has, 5))
	fmt.Println(-3, search(&has, -3))
	fmt.Println(-1, search(&has, -1))
	fmt.Println(7, search(&has, 7))
}

func testHm2() {
	var m HashMap
	m = NewHashLinkMap(10)
	i := -10
	for {
		if i > 20 {
			break
		}
		m.Insert(i, i*10)
		i++
	}

	fmt.Println(m.Search(4))
	fmt.Println(m.Search(-4))
	fmt.Println(m.Search(14))
	fmt.Println(m.Search(8))
	fmt.Println(m.Search(-8))
	fmt.Println(m.Search(-18))
	fmt.Println(m.Search(-28))
	fmt.Println(m)

	m.Del(-4)
	m.Del(-1)

	m.Del(-8)
	m.Del(0)
	fmt.Println(m.Search(-4))
	fmt.Println(m)
	m.HmPrint()
}

func testHm3() {
	var m HashMap
	m = NewHashAddrMap()
	m.Insert(1, 1)
	m.Insert(-1, 1)
	m.Insert(3, 3)
	m.Insert(0, 100)
	m.Insert(11, 11)
	m.Insert(-11, -11)
	m.Insert(5, 5)
	fmt.Println(m)
	m.HmPrint()
	fmt.Println(m.Search(1))
	fmt.Println(m.Search(4))
	fmt.Println(m.Search(11))
	m.Del(11)
	fmt.Println(m.Search(11))
	m.Del(5)
	fmt.Println(m)
}

func testHm4() {
	var m HashMap
	m = NewHashAddrMap()

	fmt.Println(m)
	m.HmPrint()
	for i := -40; i < 80; i++ {
		m.Insert(i, i*100)
	}
	//fmt.Println(m)
	//m.HmPrint()
	fmt.Println(m.Search(-30))
	fmt.Println(m.Search(10))
	fmt.Println(m.Search(-10))
	fmt.Println(m.Search(81))

}

func printVerticalOrder(root *Node, hd int, m HashMap) {

	if root == nil {
		return
	}

	//	err,oodv := m.Search(hd)
	//	if err != nil {
	//		//not this level data
	//	}else{
	//
	//	}
	fmt.Printf("hd=%d,v=%d\n", hd, root.Val)

	printVerticalOrder(root.Left, hd-1, m)
	printVerticalOrder(root.Right, hd+1, m)
}

func printVerticalOrderSys(root *Node, hd int, m map[int][]int) {

	if root == nil {
		return
	}

	oodv := m[hd]
	if oodv != nil {
		//not this level data
		oodv = append(oodv, root.Val)
		m[hd] = oodv
	} else {
		m[hd] = []int{root.Val}
	}
	//fmt.Printf("hd=%d,v=%d\n", hd, root.Val)

	printVerticalOrderSys(root.Left, hd-1, m)
	printVerticalOrderSys(root.Right, hd+1, m)
}

func testPrintvO() {
	var m HashMap
	m = NewHashAddrMap()

	root := NewNode(6)
	root.InsertNode(NewNode(4))
	root.InsertNode(NewNode(8))
	root.Left.InsertNode(NewNode(2))
	root.Right.InsertNode(NewNode(5))
	root.Left.Left.InsertNode(NewNode(1))
	root.Right.InsertNode(NewNode(10))
	PreOrder(root)
	m.HmPrint()
	printVerticalOrder(root, 0, m)

	m2 := make(map[int][]int)
	printVerticalOrderSys(root, 0, m2)
	fmt.Println("use system map")
	for _, v := range m2 {
		fmt.Println(v)
	}
}

/**
*Minimum delete operations to make all elements of array same
Time complexity : O(n)
*O(n)
*/
func minDelete(a []int) int {
	if len(a) < 2 {
		return 1
	}
	m := make(map[int]int)
	maxFeq := 0
	for _, v := range a {
		od, ok := m[v]

		if !ok {
			od = 1
		} else {
			od++
		}
		if maxFeq < od {
			maxFeq = od
		}
		m[v] = od
	}

	fmt.Printf("maxFeq:%d\n", maxFeq)
	return len(a) - maxFeq
}

// Function to find maximum distance between equal elements
func maxDistance(a []int) int {
	if len(a) < 2 {
		return 1
	}
	var m HashMap
	m = NewHashAddrMap()
	MaxDis := 0
	for i, v := range a {
		err, od := m.Search(v)
		//fmt.Println(err, od)
		if err != nil {
			m.Insert(v, i)
		} else {
			if i-od > MaxDis {
				MaxDis = i - od
			}
		}
	}
	fmt.Printf("maxDistance is :%d\n", MaxDis)
	return MaxDis
}

func main() {
	//test1()
	//testHm()
	//testHm2()
	testHm4()
	testPrintvO()
	//fmt.Println((75 * 100) / 100)
	a := []int{5, 0, 0, 9, 5, 1, 5}
	fmt.Println(minDelete(a))
	fmt.Println(maxDistance(a))
	a2 := []int{3, 2, 1, 2, 1, 4, 5, 8, 6, 7, 4, 2}
	fmt.Println(maxDistance(a2))
}
