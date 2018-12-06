package main

import "fmt"

type HashMap interface {
	Insert(k, v int)
	Search(k int) (error, int)
	Del(k int)
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

func testHm() {
	var m HashMap
	m = NewHashList(10)
	m.Insert(4, 4)
	m.Insert(8, 8)
	m.Insert(-8, -8)
	m.Insert(18, 18)
	m.Insert(14, 14)
	m.Insert(0, 0)

	fmt.Println(m.Search(4))
	fmt.Println(m.Search(8))
	fmt.Println(m.Search(-8))
	fmt.Println(m.Search(-18))
	fmt.Println(m.Search(-28))
	fmt.Println(m)
}

func testHm2() {
	var m HashMap
	m = NewHashList(10)
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
}

func main() {
	//test1()
	testHm()
	testHm2()
}
