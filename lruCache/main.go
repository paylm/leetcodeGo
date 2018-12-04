package main

import "fmt"

func main() {
	fmt.Println("lruCache test")
	lru := NewLRU(3)
	lru.put("k1", "v111")
	lru.put("k2", "v112")
	lru.put("k3", "v113")
	lru.put("k4", "v114")
	lru.show()
	lru.put("k5", "v115")
	//lru.show()
	lru.get("k3")
	lru.get("k4")
	//lru.show()
	lru.put("k5", "v115")
	lru.show()
	lru.put("k6", "v116")
	lru.show()
	lru.put("k2", "v122")
	//	lru.put("k3", "v133")
	lru.get("k2")
	lru.get("k4")
	//	lru.get("k2")
	//	lru.get("k2")
	lru.show()
}
