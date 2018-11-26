package main

import "fmt"

var (
	level = 1
)

type Noder interface {
	InsertNode(n *Node)
	Gt(b *Node) bool
	PreOrder()  //前遍历
	InOrder()   //中遍历
	PostOrder() //后遍历
	GetVal() interface{}
	Tree2Search(v int, t int)
	MaxNode() *Node
	MinNode() *Node
	Delete(v int)
}

type Node struct {
	Val   int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	n := new(Node)
	n.Val = val
	return n
}

func (rootNode *Node) GetVal() interface{} {
	if rootNode != nil {
		return rootNode.Val
	}
	return -1
}

func (rootNode *Node) Gt(b *Node) bool {
	if rootNode.Val > b.Val {
		return true
	}
	return false
}

func (rootNode *Node) InsertNode(n *Node) {

	if rootNode == nil {
		return
	}
	//	defer func() {
	//		fmt.Printf("insert %d ok \n", n.Val)
	//	}()
	if rootNode.Gt(n) {
		// root > b
		if rootNode.left == nil {
			rootNode.left = n
			return
		}
		rootNode.left.InsertNode(n)

	} else {
		if rootNode.right == nil {
			rootNode.right = n
			return
		}
		rootNode.right.InsertNode(n)
	}

}

func (rootNode *Node) PreOrder() {
	if rootNode == nil {
		//fmt.Println("is a Empty Node")
		return
	}
	defer func() {
		//出醆
		level--
	}()
	rootNode.left.PreOrder()
	fmt.Printf("level=%d,val=%d \n", level, rootNode.GetVal())
	//入醆
	level++
	rootNode.right.PreOrder()
}

func (rootNode *Node) InOrder() {

	if rootNode == nil {
		//fmt.Println("is a Empty Node")
		return
	}

	fmt.Println(rootNode.GetVal())
	rootNode.left.InOrder()
	rootNode.right.InOrder()
}

func (rootNode *Node) PostOrder() {

	if rootNode == nil {
		//fmt.Println("is a Empty Node")
		return
	}
	rootNode.right.PostOrder()
	fmt.Println(rootNode.GetVal())
	rootNode.left.PostOrder()
}

func (rootNode *Node) MaxNode() *Node {

	var n *Node
	if rootNode != nil {
		n = rootNode
	}

	for {
		if n.right == nil {
			return n
		}
		p := n.right
		n = p
	}
	return n
}

func (rootNode *Node) MinNode() *Node {
	var n *Node
	if rootNode != nil {
		n = rootNode
	}

	for {
		if n.left == nil {
			return n
		}
		p := n.left
		n = p
	}
	return n
}

//生成二叉树
func MakeTree(arr []int) Noder {
	var rootN Noder
	for _, v := range arr {
		n := NewNode(v)
		if rootN == nil {
			rootN = n
		} else {
			rootN.InsertNode(n)
		}
	}
	return rootN
}

func (rootNode *Node) Delete(v int) {
	//册左:从最右取小值补原位
	//删右:
	if rootNode == nil {
		return
	}
	isleft := false
	parent, delNode := rootNode, rootNode

	for {

		if delNode == nil {
			fmt.Printf("找不到删除点:%d\n", v)
			return
		}

		if delNode.Val == v {
			fmt.Printf("foud then del Node :%d\n", v)
			break
		}

		if delNode.Val > v {
			parent = delNode
			delNode = parent.left
			isleft = true
		} else {
			parent = delNode
			delNode = parent.right
			isleft = false
		}
	}

	fmt.Printf("parent:%d,crrent=%d \n", parent.Val, delNode.Val)

	//del action
	if delNode == rootNode {
		fmt.Printf("oldroot:%d,newroot:%d\n", rootNode, n)
		return
	}
	//case 1 : it's last node
	if delNode.right == nil && delNode.left == nil {
		if isleft == true {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return
	}

	//case 2 : left or right is nil
	if delNode.right != nil {
		rmax := delNode.right.MinNode()
		rmax.left = delNode.left
		rmax.right = delNode.right
		if isleft {
			parent.left = rmax
		} else {
			parent.right = rmax
		}
		return
	}

	if delNode.right == nil {
		lmax := delNode.left.MaxNode()
		lmax.left = delNode.left
		lmax.right = delNode.right
		if isleft {
			parent.left = lmax
		} else {
			parent.right = lmax
		}
		return
	}

}

func (rootNode *Node) Tree2Search(v int, t int) {

	if rootNode == nil {
		fmt.Printf("找不到此值:%d,用时:%d \n", v, t)
		return
	}
	if rootNode.Val == v {
		fmt.Printf("已找到此值:%d,用时%d\n", v, t)
		return
	}
	t++
	if rootNode.Val > v {
		rootNode.left.Tree2Search(v, t)
	} else {
		rootNode.right.Tree2Search(v, t)
	}
}

func main() {
	fmt.Println("TwoT-go")
	arr := []int{8, 2, 3, 10, 1, 16, 5, 11, 7, 9}

	ns := MakeTree(arr)
	fmt.Println("------InOrder-----")
	ns.InOrder()
	fmt.Println("------PreOrder-----")
	ns.PreOrder()
	fmt.Println("------PostOrder-----")
	ns.PostOrder()

	fmt.Println("---------二叉树查找-------")
	ns.Tree2Search(7, 1)
	ns.Tree2Search(12, 1)
	ns.Tree2Search(9, 1)
	ns.Tree2Search(6, 1)

	fmt.Println("-------找到最大值-----")
	m := ns.MaxNode()
	fmt.Println(m.GetVal())

	fmt.Println("-------找到最小值-----")
	n2 := ns.MinNode()
	fmt.Println(n2.GetVal())
	fmt.Println("--------del node -------")
	ns.Delete(6)
	ns.Delete(8)
	fmt.Printf("rootaddr:%d", ns)
	ns.InOrder()
}
