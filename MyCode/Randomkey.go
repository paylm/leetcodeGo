package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func RandomKey() int {
	k := 1
	seed := fmt.Sprintf("%d", &k)
	ks, _ := strconv.ParseInt(seed, 10, 64)
	//fmt.Println(ks)
	rand.Seed(ks)
	v := rand.Intn(100)
	return v
}
