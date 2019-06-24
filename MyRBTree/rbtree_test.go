package main

import "testing"

func Test_Insert(t *testing.T) {
	data := []int{6, 8, 11, 17, 15, 22, 1, 13, 25, 27}
	var root *RBNode
	for _, v := range data {
		root = insert(root, v)
	}
	//print_tree(root)
	for _, d := range data {
		if search(root, d) == nil {
			t.Errorf("test fail , %d can't be found\n", d)
		}
	}
}

func Test_FindMax(t *testing.T) {
	data := []struct {
		arr []int
		max int
	}{
		{arr: []int{1, 7, 9, 3, 1}, max: 9},
		{arr: []int{6}, max: 6},
		{arr: []int{6, 8, 11, 17, 15, 22, 1, 13, 25, 27}, max: 27},
	}
	for _, d := range data {
		var root *RBNode
		for _, v := range d.arr {
			root = insert(root, v)
		}
		res := findMaxNode(root)
		if d.max != res.val {
			t.Errorf("test fail, %d should be max at %v , but return %d\n", d.max, d, res.val)
		}
	}
}

func Test_Delete(t *testing.T) {
	data := []struct {
		arr []int
		del int
	}{
		{arr: []int{1, 7, 9, 3, 1}, del: 9},
		{arr: []int{6}, del: 6},
		{arr: []int{6, 8, 11, 17, 15, 22, 1, 13, 25, 27}, del: 27},
	}
	for _, d := range data {
		var root *RBNode
		for _, v := range d.arr {
			root = insert(root, v)
		}
		root = delete_recurse(root, d.del)
		res := search(root, d.del)
		if res != nil {
			t.Errorf("delete fail , %d should be del at %v, but stil found it\n", d.del, d)
		}
		print_tree(root)
	}
}
