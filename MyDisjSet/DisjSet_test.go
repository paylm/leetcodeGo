package MyDisjSet

import (
	"testing"
	"fmt"
)

func TestUnionSet(t *testing.T) {
	dj := NewDisjSet(9)
	UnionSet(dj,5,7)
	UnionSet(dj,5,8)
	UnionSet(dj,2,3)

	if dj.S[5] != -2 {
		t.Errorf("test fail , dj.S[%d] should be %d",5,-2)
	}
	if dj.S[2] != -2 {
		t.Errorf("test fail , dj.S[%d] should be %d",2,-2)
	}
	//fmt.Println(dj.S)
	UnionSet(dj,5,2)
	if dj.S[5] != -3 || dj.S[2] != 5 {
		t.Errorf("test fail at  UnionSet(dj,5,2) ")
	}
	fmt.Println(dj.S)
}

func TestFind(t *testing.T) {
	dj := NewDisjSet(12)
	UnionSet(dj,5,7)
	UnionSet(dj,5,8)
	UnionSet(dj,2,3)
	UnionSet(dj,2,4)
	UnionSet(dj,5,2)
	UnionSet(dj,9,10)
	if Find(dj,2) != Find(dj,8){
		t.Errorf("test fail , %d %d belong same set",2,8)
	}

	if Find(dj,1) == Find(dj,8){
		t.Errorf("test fail , %d(%d) %d(%d) is not belong a same set",1,Find(dj,1),8,Find(dj,8))
	}
	fmt.Println(dj.S)
}