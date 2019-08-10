package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	MAX_LEVEL = 64
)

type Data interface {
	compare() bool
}

type ZskiplistNode struct {
	obj      string
	score    int
	backward *ZskiplistNode
	zkLevel  []*ZskiplistNode //
}

type Zskiplist struct {
	head  *ZskiplistNode
	tail  *ZskiplistNode
	level int //最大已用高度
	len   int
}

func NewZskiplistNode(obj string, score int, level int) *ZskiplistNode {
	zk := new(ZskiplistNode)
	zk.obj = obj
	zk.score = score
	zk.zkLevel = make([]*ZskiplistNode, level)
	return zk
}

func NewZskiplist() *Zskiplist {
	zkl := new(Zskiplist)
	zkl.head = NewZskiplistNode("", 0, MAX_LEVEL)
	zkl.tail = NewZskiplistNode("", 0, MAX_LEVEL)
	for i := 0; i < MAX_LEVEL; i++ {
		zkl.head.zkLevel[i] = zkl.tail
	}
	zkl.level = 1
	zkl.tail.backward = zkl.head
	return zkl
}

func randLevel() int {
	l := 1
	rand.Seed(time.Now().UnixNano())
	for {
		if rand.Intn(2) == 0 {
			break
		}
		l++
	}
	if l < MAX_LEVEL {
		return l
	}
	return MAX_LEVEL
}

func (zkl *Zskiplist) zslInsert(obj string, score int) error {
	n := zkl.head
	update := make([]*ZskiplistNode, zkl.level) //侍插入点的前指针
	i := zkl.level - 1
	for {
		if i < 0 {
			break
		}

		if n.score == score && n != zkl.head {
			return errors.New("arealdy exsit,insert fail")
		}

		if n.zkLevel[i] == zkl.tail || n.zkLevel[i].score > score {
			//下沉
			update[i] = n
			i--
		} else {
			n = n.zkLevel[i]
		}

	}
	//fmt.Printf("update:%v\n", update)
	iLevel := randLevel()
	iN := NewZskiplistNode(obj, score, iLevel)
	for j := iLevel - 1; j >= 0; j-- {
		if len(update) > j {
			iN.zkLevel[j] = update[j].zkLevel[j]
			update[j].zkLevel[j] = iN
		} else {
			iN.zkLevel[j] = zkl.head.zkLevel[j]
			zkl.head.zkLevel[j] = iN
		}
	}
	//fix backward
	iN.zkLevel[0].backward = iN
	iN.backward = update[0]

	if iLevel > zkl.level {
		zkl.level = iLevel
	}
	zkl.len++
	return nil
}

func (zkl *Zskiplist) zslDelete(obj string, score int) bool {

	c := zkl.head
	i := zkl.level - 1
	var delNode *ZskiplistNode
	update := make([]*ZskiplistNode, zkl.level)
	for {
		if i < 0 {
			break
		}
		if c.zkLevel[i] == zkl.tail || c.zkLevel[i].score >= score {
			//下降一级查找
			update[i] = c
			if c.zkLevel[i].score == score {
				delNode = c.zkLevel[i]
			}
			i--
		} else {
			c = c.zkLevel[i]
		}
	}

	if delNode == nil || delNode == zkl.head || delNode == zkl.tail {
		return true
	}
	//other level
	for j := len(update) - 1; j > 0; j-- {
		if len(delNode.zkLevel) > j {
			update[j].zkLevel[j] = delNode.zkLevel[j]
		}
	}
	//delte
	delNode.zkLevel[0].backward = delNode.backward
	delNode.backward.zkLevel[0] = delNode.zkLevel[0]

	return true
}

func (zsk *Zskiplist) zslFind(score int) *ZskiplistNode {
	c := zsk.head
	i := zsk.level - 1
	for {
		if i < 0 {
			break
		}
		if c.score == score && c != zsk.head {
			return c
		}
		if c.zkLevel[i] == zsk.tail || c.zkLevel[i].score > score {
			//下降一级查找
			i--
		} else {
			c = c.zkLevel[i]
		}
	}
	return nil
}

func (zsl *Zskiplist) zslUpdate(obj string, score int) bool {
	n := zsl.zslFind(score)
	if n == nil {
		return false
	}
	n.obj = obj
	return true
}

func (zsk *Zskiplist) show() {

	c := zsk.head
	for {
		if c == nil {
			break
		}
		//fmt.Printf("%d => %v\n", &(c), c)
		if c != zsk.head {
			fmt.Printf("%d(%d) \t %s\t", c.score, c.backward.score, c.obj)
		} else {
			fmt.Printf("%d(nil) \t %s\t", c.score, c.obj)
		}
		for i := 0; i < len(c.zkLevel); i++ {
			if c.zkLevel[i] != nil && i <= zsk.level {
				fmt.Printf(" %d\t", c.zkLevel[i].score)
			}
		}
		fmt.Printf("\n")
		c = c.zkLevel[0]
	}
}

func (n *ZskiplistNode) show() {
	fmt.Printf("%d \t %s \t", n.score, n.obj)
	for i := 0; i < len(n.zkLevel); i++ {
		fmt.Printf("%d\t", n.zkLevel[i].score)
	}
	fmt.Println()
}
