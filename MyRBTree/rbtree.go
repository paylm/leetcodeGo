package main

import "fmt"

//https://en.wikipedia.org/wiki/Red%E2%80%93black_tree

type Colortype int

const (
	red Colortype = iota
	black
)

type RBNode struct {
	val    int
	left   *RBNode
	right  *RBNode
	color  Colortype
	parent *RBNode
}

//p node
func parent(n *RBNode) *RBNode {
	if n == nil {
		return nil // NULL for root node
	}
	return n.parent
}

//grandparent node
func grandparent(n *RBNode) *RBNode {
	p := parent(n)
	if p == nil {
		return nil // No parent means no grandparent
	}
	return parent(p) // NULL if parent is root
}

// uncle node
func uncle(n *RBNode) *RBNode {
	g := grandparent(n)
	p := parent(n)
	if g == nil {
		return nil //// No grandparent means no uncle
	}
	return sibling(p)

}

func sibling(n *RBNode) *RBNode {
	p := parent(n)
	if p == nil {
		return nil //// No parent means no sibling
	}

	if n == p.left {
		return p.right
	} else {
		return p.left
	}
}

/*right
      G                G
     / \              / \
	P  U    =>       K1   U
   / \              / \
  K1  K2               P
   \                  / \
   K1R	 			 K1R K2

	 G                 G
    / \               / \
   U   P    =>       U  k1
      / \               / \
     k1  k2                p
	  \                   / \
       K1R               K1R k2

  @param  n : rotate point P
**/
func rotate_right(n *RBNode) {
	g := parent(n)

	k1 := n.left
	k1r := k1.right

	//rotate_right subtree
	k1.right = n
	n.parent = k1
	n.left = k1r
	if g != nil {

		if g.left == n {

			g.left = k1
			k1.parent = g
			if k1r != nil {
				k1r.parent = n
			}

		} else if g.right == n {
			g.right = k1
			k1.parent = g
			if k1r != nil {
				k1r.parent = n
			}
		}
	} else {
		//if G is nil
		k1.parent = nil
	}
}

/*left rotate
    G                 G
   /  \              /   \
  U	   P   =>       U     K2
	 /  \          / \   / \
	k1   k2             P
	    /              / \
	   K2L            K1  K2L


      G              G
     / \            /  \
    P   U          K2   U
   / \       =>   /  \
  K1  K2         P
      / \       / \
     K2L       K1 K2L
  @param  n : rotate point P
**/
func rotate_left(n *RBNode) {
	g := parent(n)
	k2 := n.right
	k2l := k2.left

	//rotate_left subtree
	k2.left = n
	n.parent = k2
	n.right = k2l
	if k2l != nil {
		k2l.parent = n
	}

	//reset the root
	if g != nil {
		if g.right == n {
			//case 1
			g.right = k2
			k2.parent = g
		} else {
			//case 2
			g.left = k2
			k2.parent = g
		}
	} else {
		//if G is nil
		k2.parent = nil
	}
}

func NewRBNode(v int) *RBNode {
	n := new(RBNode)
	n.val = v
	n.color = red
	return n
}

func (n *RBNode) is_leaf() bool {
	if n != nil && (n.left != nil || n.right != nil) {
		return false
	}
	return true
}

func insert(root *RBNode, v int) *RBNode {
	//fmt.Printf("root:%v,v:%d\n", root, v)
	n := NewRBNode(v)
	// insert new node into the current tree
	insert_recurse(root, n)

	// repair the tree in case any of the red-black properties have been violated
	insert_repair_tree(n)

	// find the new root to return
	root = n
	for {
		if root.parent == nil {
			break
		}
		root = root.parent
	}
	//fmt.Printf("insert %d, root:%d\n", v, root.val)
	return root
}

func insert_recurse(root *RBNode, n *RBNode) {
	if root == nil { // if tree is nil
		return
	}
	c := root
	for {
		if c == nil {
			break
		}
		if c.val > n.val {
			// c is right of n
			if c.left == nil {
				c.left = n
				n.parent = c
				break
			}
			c = c.left
		} else {
			// n is left of c
			if c.right == nil {
				c.right = n
				n.parent = c
				break
			}
			c = c.right
		}
	}
	//print_tree(root)
}

//aussme , new node n must be red node
func insert_repair_tree(n *RBNode) {
	p := parent(n)
	//case 1 , n-> parent is nil
	if p == nil {
		insert_repair_case1(n)
		return
	}
	//case 2, n -> parent is black node , no operation
	if p.color == black {
		insert_repair_case2(n)
		return
	}
	//case 3 , n -> parent is red && n -> uncle is red
	// make parent and uncle to black and grandparent to red
	u := uncle(n)
	if u != nil && u.color == red {
		insert_repair_case3(n)
		return
	}

	//case 4 , The parent P is red but the uncle U is black.
	if p.color == red {
		insert_repair_case4(n)
	}
}

func insert_repair_case1(n *RBNode) {
	//fmt.Printf("insert_repair_case1 \n")
	//n.leaf = false
	n.color = black
}

func insert_repair_case2(n *RBNode) {
	//fmt.Printf("insert_repair_case2 \n")
	return // don't do anything
}

func insert_repair_case3(n *RBNode) {
	//fmt.Printf("insert_repair_case3\n")
	p := parent(n)
	u := sibling(p)
	p.color = black
	u.color = black
	grandparent(n).color = red
	insert_repair_tree(grandparent(n))
}

//insert_case4 , repaire color
func insert_repair_case4(n *RBNode) {
	//fmt.Printf("insert_repair_case4\n")
	g := grandparent(n)
	p := parent(n)
	//N is the left child of the right child of the grandparent or the right child of the left child of the grandparent
	if n == p.right && g.left == p {
		//rotate_right
		rotate_left(p)
		n = n.left
	} else if n == p.left && g.right == p {
		//rotate_left
		rotate_right(p)
		n = n.right
	}

	insert_repair_case4step2(n)
}

func insert_repair_case4step2(n *RBNode) {
	//The current node N is now certain to be on the "outside" of the subtree under G (left of left child or right of right child).
	//fmt.Printf("insert_repair_case4step2\n")
	g := grandparent(n)
	p := parent(n)

	if n == p.left {
		//fmt.Printf("pattern / rotate_right n:%v\n", g)
		rotate_right(g) //todo
	} else if n == p.right {
		//fmt.Printf("pattern \\ rotate_left n:%v\n", g)
		rotate_left(g) //todo
	}

	p.color = black
	g.color = red

}

/**
 find the max node
**/
func findMaxNode(n *RBNode) *RBNode {
	if n == nil {
		return nil
	}
	c := n
	for {
		if c.right == nil {
			break
		}
		c = c.right
	}

	return c
}

/**
标准二叉树节点删除，red-black 比这个复杂
**/

func delete_recurse(n *RBNode, v int) *RBNode {
	if n == nil {
		return nil
	}
	root := n
	if n.val > v {
		n.left = delete_recurse(n.left, v)
	} else if n.val < v {
		n.right = delete_recurse(n.right, v)
	} else {
		//当前节点值 == v
		//case 1 , n is leaf
		if n.left == nil && n.right == nil {
			return nil
		}

		if n.left == nil && n.right != nil {
			return n.right
		}

		if n.right == nil && n.left != nil {
			return n.left
		}

		//左右节占都不为空时
		temp := findMaxNode(n.left)
		n.val = temp.val
		n.left = delete_recurse(n.left, n.val)
		return n
	}

	return root
}

/**
 删除节点根下的某个节点,返回最后的根节点
 todo
 RB - DELETE(T, z)
 if left[z] = nil[T] or right[z] = nil[T]
    then y ← z                                  // 若“z的左孩子” 或 “z的右孩子”为空，则将“z”赋值给 “y”；
    else y ← TREE - SUCCESSOR(z)                  // 否则，将“z的后继节点”赋值给 “y”。
 if left[y] ≠ nil[T]
    then x ← left[y]                            // 若“y的左孩子” 不为空，则将“y的左孩子” 赋值给 “x”；
    else x ← right[y]                           // 否则，“y的右孩子” 赋值给 “x”。
 p[x] ← p[y]                                    // 将“y的父节点” 设置为 “x的父节点”
 if p[y] = nil[T]
    then root[T] ← x                            // 情况1：若“y的父节点” 为空，则设置“x” 为 “根节点”。
    else if y = left[p[y]]
            then left[p[y]] ← x                 // 情况2：若“y是它父节点的左孩子”，则设置“x” 为 “y的父节点的左孩子”
            else right[p[y]] ← x                // 情况3：若“y是它父节点的右孩子”，则设置“x” 为 “y的父节点的右孩子”
 if y ≠ z
    then key[z] ← key[y]                        // 若“y的值” 赋值给 “z”。注意：这里只拷贝z的值给y，而没有拷贝z的颜色！！！
         copy y's satellite data into z
 if color[y] = BLACK
    then RB - DELETE - FIXUP(T, x)                  // 若“y为黑节点”，则调用
 return y
**/
func delete_one(n *RBNode, v int) *RBNode {
	if n == nil {
		return nil
	}
	root := n
	z := search(n, v)
	if z == nil {
		return root
	}
	var y *RBNode // y is leaf or y has one child
	if z.left != nil || z.right != nil {
		y = z
	} else {
		y = findMaxNode(z.left)
	}

	if y == root {
		return nil
	}

	p := y.parent

	if y.left != nil {
		y.val = y.left.val
		y.left = nil
	} else if y.right != nil {
		y.val = y.right.val
		y.right = nil
	} else {
		// y is leaf
		if y.color == black {
			delete_fixup(y) //可能改变
		} else {
			if p.left == y {
				p.left = nil
			} else {
				p.right = nil
			}
		}
	}
	return root
}

/*对节点重新着色并旋转，以此来保证删除节点后的树仍然是一颗红黑树
RB - DELETE - FIXUP(T, x)
 while x ≠ root[T] and color[x] = BLACK
     do if x = left[p[x]]
           then w ← right[p[x]]                                             // 若 “x”是“它父节点的左孩子”，则设置 “w”为“x的叔叔”(即x为它父节点的右孩子)
                if color[w] = RED                                           // Case 1: x是“黑+黑”节点，x的兄弟节点是红色。(此时x的父节点和x的兄弟节点的子节点都是黑节点)。
                   then color[w] ← BLACK ? Case 1   //   (01) 将x的兄弟节点设为“黑色”。
                        color[p[x]] ← RED ? Case 1   //   (02) 将x的父节点设为“红色”。
                        LEFT - ROTATE(T, p[x]) ? Case 1   //   (03) 对x的父节点进行左旋。
                        w ← right[p[x]] ? Case 1   //   (04) 左旋后，重新设置x的兄弟节点。
                if color[left[w]] = BLACK and color[right[w]] = BLACK       // Case 2: x是“黑+黑”节点，x的兄弟节点是黑色，x的兄弟节点的两个孩子都是黑色。
                   then color[w] ← RED ? Case 2   //   (01) 将x的兄弟节点设为“红色”。
                        x ←  p[x] ? Case 2   //   (02) 设置“x的父节点”为“新的x节点”。
                   else if color[right[w]] = BLACK                          // Case 3: x是“黑+黑”节点，x的兄弟节点是黑色；x的兄弟节点的左孩子是红色，右孩子是黑色的。
                           then color[left[w]] ← BLACK ? Case 3   //   (01) 将x兄弟节点的左孩子设为“黑色”。
                                color[w] ← RED ? Case 3   //   (02) 将x兄弟节点设为“红色”。
                                RIGHT - ROTATE(T, w) ? Case 3   //   (03) 对x的兄弟节点进行右旋。
                                w ← right[p[x]] ? Case 3   //   (04) 右旋后，重新设置x的兄弟节点。
                         color[w] ← color[p[x]] ? Case 4   // Case 4: x是“黑+黑”节点，x的兄弟节点是黑色；x的兄弟节点的右孩子是红色的。(01) 将x父节点颜色 赋值给 x的兄弟节点。
                         color[p[x]] ← BLACK ? Case 4   //   (02) 将x父节点设为“黑色”。
                         color[right[w]] ← BLACK ? Case 4   //   (03) 将x兄弟节点的右子节设为“黑色”。
                         LEFT - ROTATE(T, p[x]) ? Case 4   //   (04) 对x的父节点进行左旋。
                         x ← root[T] ? Case 4   //   (05) 设置“x”为“根节点”。
        else (same as then clause with "right" and "left" exchanged)        // 若 “x”是“它父节点的右孩子”，将上面的操作中“right”和“left”交换位置，然后依次执行。
 color[x] ← BLACK
**/
func delete_fixup(n *RBNode) *RBNode {
	return nil
}

/**
 中序遍历
**/
func print_tree(root *RBNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d color:(%d) leaf:(%v)\n", root.val, root.color, root.is_leaf())
	print_tree(root.left)
	print_tree(root.right)
}

func search(n *RBNode, v int) *RBNode {
	if n == nil {
		return nil
	}

	if n.val == v {
		return n
	}
	if n.val < v {
		return search(n.right, v)
	} else {
		return search(n.left, v)
	}
}
