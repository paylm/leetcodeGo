package main

import (
	"fmt"
	"strings"
)

//从数组中找到两个数的值和为k
func findNumS(a []int, k int) {
	if len(a) == 0 {
		return
	}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
	i, j := 0, len(a)-1
	for {

		fmt.Printf("i:%d,j:%d\n", i, j)
		if i >= j {
			break
		}
		if a[j] > k {
			j--
			continue
		} else {
			if a[j]+a[i] == k {
				fmt.Printf("%d+%d\n", a[i], a[j])
				return
			} else if a[j]+a[i] > k {
				j--
			} else {
				//if a[j]+a[i] <k {
				i++
			}
		}
	}
}

/**
回文数
https://leetcode-cn.com/problems/palindrome-number/
**/
func isPalindrome(x int) bool {
	if x > 0 && x < 10 {
		return true
	}
	if x < 0 {
		return false
	}

	arr := []int{}
	t := x
	for t > 0 {
		y := t % 10
		arr = append(arr, y)
		t = t / 10
	}
	fmt.Println(arr)
	i, j := 0, len(arr)-1
	for {
		if i >= j {
			break
		}

		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}

	return true
}

/**
n*n

i,j			j,n-i



n-j,i			n-i,n-j
**/

func rotate(matrix [][]int) {

	n := len(matrix) - 1
	for i := 0; i < n/2+1; i++ {
		for j := i; j < n && j+i < n; j++ {
			k := matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = matrix[j][n-i]
			matrix[j][n-i] = k
		}
	}
}

func showMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}

/**
https://leetcode-cn.com/problems/integer-to-roman/
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

示例 1:

输入: 3
输出: "III"
示例 2:

输入: 4
输出: "IV"

1 < num < 3994

**/
func intToRoman(num int) string {
	k := num

	m := make(map[int][]string)
	m[1] = []string{"I", "V", "X"}
	m[10] = []string{"X", "L", "C"}
	m[100] = []string{"C", "D", "M"}

	xm := []string{"", "M", "MM", "MMM"}

	//fmt.Println(bit)
	s := 1
	sts := []string{}

	for {
		if k <= 0 {
			break
		}
		if s < 1000 {
			sts = append(sts, (intXstr(k%10, (m[s])[0], (m[s])[1], (m[s])[2])))
		} else {
			sts = append(sts, xm[k])
		}
		s = s * 10
		//bit = append(bit, k%10)
		k = k / 10
	}

	for i, j := 0, len(sts)-1; i <= j; {
		temp := sts[i]
		sts[i] = sts[j]
		sts[j] = temp
		i++
		j--
	}

	return strings.Join(sts, "")
}

func intXstr(n int, l string, m string, h string) string {
	s := ""
	if n == 4 {
		s = fmt.Sprintf("%s%s", l, m)
	} else if n == 5 {
		s = fmt.Sprintf("%s", m)
	} else if n == 9 {
		s = fmt.Sprintf("%s%s", l, h)
	}

	k := 0
	si := []string{}
	if n < 4 {
		k = n
	} else if n > 5 && n < 9 {
		k = n - 5
		si = append(si, m)
	}

	for i := 0; i < k; i++ {
		si = append(si, l)
	}
	if k == 0 {
		return s
	} else {
		return strings.Join(si, "")
	}
}

func main() {
	//fmt.Println("vim-go")
	fmt.Println("quickSort")
	a := []int{10, 5, 4, 1, 16, 10, 5, 8, 13, 2, 10, 6, 3}
	//a := []int{10, 10, 10, 10, 10}
	b := []int{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10}
	c := []int{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10}
	d := []int{10, 5, 4, 1, 16, 8, 13, 2, 10, 6, 3, 10}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)

	fmt.Println("三路快排")
	sort3Parttion(b, 0, len(b)-1)
	fmt.Println(b)
	fmt.Println("insert Sort")
	insertSort(c)
	fmt.Println(c)
	fmt.Println("shell sort")
	shellSort(d)
	fmt.Println(d)
	//findNumS(a, 9)
	convert("fuck123", 3)
	//convert("LEETCODEISHIRING", 3)
	//fmt.Println(convert("LEETCODEISHIRING", 4))
	fmt.Println(convert("A", 1))
	fmt.Println(convert("AB", 1))
	//	fmt.Println(a)
	//	headSort(a)
	//	fmt.Println(a)
	//	k1, k2 := 11, 12
	//	fmt.Println(k1, k2)
	//	swap(&k1, &k2)
	//	fmt.Println(k1, k2)
	fmt.Println(isPalindrome(121))

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	showMatrix(matrix)
	matrix4 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	rotate(matrix4)
	showMatrix(matrix4)

	fmt.Println("test for intToRoman")
	intToRoman(3)
	intToRoman(9)
	intToRoman(58)
	fmt.Println(intToRoman(1994))
}
