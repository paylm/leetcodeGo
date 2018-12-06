package main

var (
	//Used in hash2 functoin
	PRIME = 7
)

type HashAddr struct {
	TableSize int
	used      int
	data      []*HashNode
}

type HashNode struct {
	Key        int
	Val        int
	Collicsion int
}

func NewHashNode(k, v int) *HashNode {
	n := new(HashNode)
	n.Key = k
	n.Val = v
	return n
}

func NewHashAddr(tableSize int) *HashAddr {
	h := new(HashAddr)
	h.TableSize = tableSize
	h.data = make([]*HashNode, h.TableSize)
	h.used = 0
	return h
}

func (h *HashAddr) hash1(k int) int {
	return abs(k % h.TableSize)
}

func (h *HashAddr) hash2(k int) int {
	return abs(PRIME - (k % PRIME))
}

func (h *HashAddr) isFull() bool {

	if h.used >= h.TableSize {
		return true
	}

	return false
}

func (h *HashAddr) Insert(k, v int) {
	i := h.hash1(k)
	if h.data[i] == nil {
		h.data[i] = NewHashNode(k, v)
		h.used++
		return
	}

	index2 := h.hash2(k)
	j := 1 //计数器
	for {
		newIndex := (i + index2*j) % h.TableSize
		if h.data[newIndex] == nil {
			h.data[newIndex] = NewHashNode(k, v)
			break
		}
		j++
	}
	h.used++
}

/**
线性查找，
    如果找到的位置key 不等，继续往下查
	如果找到値为 nil , 说明不存在
*/
func (h *HashAddr) Search(k int) (error, int) {
	return nil, -1
}

func (h *HashAddr) Del(k int) {

}
