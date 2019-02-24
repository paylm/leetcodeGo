package MyDisjSet

type DisjSet struct {
	S []int
}

func NewDisjSet(len int) *DisjSet {
	ds := new(DisjSet)
	ds.S = make([]int, len)
	for i := 0; i < len; i++ {
		ds.S[i] = -1
	}
	return ds
}

//union by height
func UnionSet(dj *DisjSet, root1 int, root2 int) {
	r1 := Find(dj, root1)
	r2 := Find(dj, root2)
	if r1 != r2 {
		unionSet(dj, r1, r2)
	}
}

func unionSet(dj *DisjSet, root1 int, root2 int) {
	if dj.S[root1] > dj.S[root2] {
		dj.S[root1] = root2
	} else {
		if dj.S[root1] == dj.S[root2] {
			dj.S[root1]--
		}
		dj.S[root2] = root1
	}
}

func UnionSetBySize(dj *DisjSet, root1 int, root2 int) {
	if root1 > root2 {
		dj.S[root1] = root2
	} else {
		dj.S[root2] = root1
	}
}

func Find(dj *DisjSet, e int) int {
	if dj.S[e] < 0 {
		return e
	} else {
		return Find(dj, dj.S[e])
	}
}

/*测试应用
https://leetcode-cn.com/problems/redundant-connection/
输入一个图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N) 的树及一条附加的边构成。附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。

结果图是一个以边组成的二维数组。每一个边的元素是一对[u, v] ，满足 u < v，表示连接顶点u 和v的无向图的边。

返回一条可以删去的边，使得结果图是一个有着N个节点的树。如果有多个答案，则返回二维数组中最后出现的边。答案边 [u, v] 应满足相同的格式 u < v。

示例 1：

输入: [[1,2], [1,3], [2,3]]
输出: [2,3]
解释: 给定的无向图为:
  1
 / \
2 - 3
示例 2：

输入: [[1,2], [2,3], [3,4], [1,4], [1,5]]
输出: [1,4]
解释: 给定的无向图为:
5 - 1 - 2
    |   |
    4 - 3
注意:

输入的二维数组大小在 3 到 1000。
二维数组中的整数在1到N之间，其中N是输入数组的大小。
*/
func findRedundantConnection(edges [][]int) []int {
	ds := NewDisjSet(len(edges) + 1)
	var res []int
	for _, e := range edges {
		if Find(ds, e[0]) != Find(ds, e[1]) {
			UnionSet(ds, e[0], e[1])
		} else {
			res = e
			break
		}
	}
	return res
}
