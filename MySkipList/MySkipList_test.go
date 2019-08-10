package main

import (
	"math/rand"
	"testing"
)

func Test_Insert(t *testing.T) {

	zsl := NewZskiplist()

	for i := 1; i < 30; i++ {
		k := rand.Intn(100)
		if err := zsl.zslInsert("sb", k); err != nil {
			t.Logf("insert fail at %d\n", k)
		}
	}
	zsl.show()
	if zsl.len != 11 {
		t.Errorf("test fail , %v\n", zsl)
	} else {
		t.Logf("test pass, zsl %v\n", zsl)
	}
}

func Test_Find(t *testing.T) {

	zsl := NewZskiplist()

	for i := 1; i < 1000; i++ {
		k := rand.Intn(1000)
		if err := zsl.zslInsert("sb", k); err != nil {
			//	t.Logf("insert fail at %d\n", k)
		}
	}
	zsl.zslInsert("test", 865)
	//zsl.show()
	if n := zsl.zslFind(865); n == nil {
		t.Errorf("test fail , %v\n", zsl)
	} else {
		t.Logf("test pass, zsl %v , n =  %v\n", zsl, n)
	}

	if n := zsl.zslFind(1200); n == nil {
		t.Logf("test pass , %v\n", zsl)
	} else {
		t.Errorf("test pass, zsl %v", zsl)
	}
}

func Test_Delete(t *testing.T) {
	data := []struct {
		obj   string
		score int
	}{
		{obj: "test7", score: 7},
		{obj: "test118", score: 118},
		{obj: "test1223", score: 1223},
		{obj: "test0", score: 0},
	}
	zsl := NewZskiplist()

	for i := 1; i < 1000; i++ {
		k := rand.Intn(2000)
		if err := zsl.zslInsert("sb", k); err != nil {
			//	t.Logf("insert fail at %d\n", k)
		}
	}
	for _, d := range data {
		zsl.zslInsert(d.obj, d.score)

		//zsl.show()
		zsl.zslDelete(d.obj, d.score)

		if n := zsl.zslFind(d.score); n != nil {
			t.Errorf("test fail, %v should be delete \n", d)
		}
		//zsl.show()
	}
}
