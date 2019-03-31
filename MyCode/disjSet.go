package main

/**https://leetcode-cn.com/problems/number-of-islands/
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1

**/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	//init
	nums := make([]int, len(grid)*len(grid[0])+1)
	for i := 0; i < len(nums); i++ {
		nums[i] = -1
	}
	//fmt.Printf("input grid:%v,nums:%v\n", grid, nums)
	res := 0
	for i := 0; i < len(grid); i++ {
		l := len(grid[i])
		for j := 0; j < l; j++ {
			near := false
			n := i*l + j + 1 //序列号
			//和前面的一样
			if j > 0 && grid[i][j] == '1' && grid[i][j-1] == '1' {
				jn := n - 1
				if FindRoot(nums, n) != FindRoot(nums, jn) {
					UnionSet(nums, jn, n)
				}
				near = true
			}

			//和上面的一样
			if i > 0 && grid[i][j] == '1' && grid[i-1][j] == '1' {
				in := n - l
				if FindRoot(nums, n) != FindRoot(nums, in) {
					//fmt.Printf("i:%d,j:%d,n:%d,root:%d,nums:%v\n", i, j, n, FindRoot(nums, n), nums)
					//合并高度2或以前集合时，去除1计数
					if nums[FindRoot(nums, n)] < -1 {
						res--
						//fmt.Printf("res :%d\n", res)
					}
					UnionSet(nums, in, n)
				}
				near = true
			}
			if near {
				continue
			}
			if grid[i][j] == '1' {
				//nums[n] = n
				//fmt.Printf("i:%d,j:%d,v:%d\n", i, j, nums[n])
				res++
			}
		}
	}
	//fmt.Printf("nums:%v\n", nums)
	return res
}

func UnionSet(r []int, i1 int, i2 int) {
	//fmt.Printf("UnionSet i1:%d,i2:%d\n", i1, i2)
	r1 := FindRoot(r, i1)
	r2 := FindRoot(r, i2)
	unionSet(r, r1, r2)
}

//aussume i1 i2 are root
func unionSet(r []int, i1 int, i2 int) {

	if r[i1] > r[i2] {
		r[i1] = i2
	} else {
		if r[i1] == r[i2] {
			r[i1]--
		}
		r[i2] = i1
	}
}

func FindRoot(r []int, i int) int {
	if r[i] < 0 {
		return i
	} else {
		return FindRoot(r, r[i])
	}
}

/**
https://leetcode-cn.com/problems/surrounded-regions/
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X

**/
func solve(board [][]byte) {
	bvisit := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		bvisit[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1 {
				//first and last line
				if bvisit[i][j] == false {
					search(board, bvisit, i, j)
				}
			}
		}
	}
	// set default value
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if bvisit[i][j] == false {
				board[i][j] = 'X'
			}
		}
	}
}

func search(board [][]byte, bvisit [][]bool, i int, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if bvisit[i][j] {
		return
	}
	if board[i][j] == 'O' {
		bvisit[i][j] = true
		search(board, bvisit, i-1, j)
		search(board, bvisit, i+1, j)
		search(board, bvisit, i, j-1)
		search(board, bvisit, i, j+1)
	}
}
