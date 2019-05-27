package MyRBTree

import "testing"

func Test_Insert(t *testing.T) {
	data := []int{6, 8, 11, 17, 15, 22, 1, 13, 25, 27}
	var root *RBNode
	for _, v := range data {
		root = insert(root, v)
	}
	print_tree(root)
}
