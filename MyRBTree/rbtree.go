package MyRBTree

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
	leaf   bool
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
	if g == nil {
		return
	}
	k1 := n.left
	k1r := k1.right
	n.left = k1r
	if g.left == n {

		g.left = k1
		k1.parent = g
		k1.right = n
		n.parent = k1
		if k1r != nil {
			k1r.parent = n
		}

	} else if g.right == n {
		g.right = k1
		k1.parent = g
		k1.right = n
		n.parent = k1
		if k1r != nil {
			k1r.parent = n
		}
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
	if g == nil {
		return
	}
	k2 := n.right
	k2l := k2.left
	if g.right == n {
		//case 1
		g.right = k2
		k2.parent = g
		k2.left = n
		n.parent = k2
		n.right = k2l
		if k2l != nil {
			k2l.parent = n
		}
	} else {
		//case 2
		g.left = k2
		k2.parent = g
		k2.left = n
		n.parent = k2
		n.right = k2l
		if k2l != nil {
			k2l.parent = n
		}
	}
}

func NewRBNode(v int) *RBNode {
	n := new(RBNode)
	n.val = v
	n.leaf = true
	n.color = red
	return n
}

func insert(root *RBNode, v int) *RBNode {
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
			// n is right of c
			if c.leaf {
				c.right = n
				c.leaf = false
				break
			}
			c = c.right
		} else {
			// n is left of c
			if c.leaf {
				c.left = n
				c.leaf = false
				break
			}
			c = c.left
		}
	}
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
	if p.color == red && u != nil && u.color == black {
		insert_repair_case4(n)
	}
}

func insert_repair_case1(n *RBNode) {
	n.leaf = true
	n.color = black
}

func insert_repair_case2(n *RBNode) {
	return // don't do anything
}

func insert_repair_case3(n *RBNode) {
	p := parent(n)
	u := sibling(p)
	p.color = black
	u.color = black
	grandparent(n).color = red
	insert_repair_tree(grandparent(n))
}

//insert_case4 , repaire color
func insert_repair_case4(n *RBNode) {
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
	g := grandparent(n)
	p := parent(n)

	if n == p.left {
		rotate_right(p)
	} else if n == p.right {
		rotate_left(p)
	}

	p.color = black
	g.color = red

}
