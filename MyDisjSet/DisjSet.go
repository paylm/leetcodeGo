package MyDisjSet


type DisjSet struct{
	S  []int
}

func NewDisjSet(len int) *DisjSet{
	ds := new(DisjSet)
	ds.S = make([]int,len)
	for i:=0;i<len;i++{
		ds.S[i] = -1
		ds.S[i] = -1
	}
	return ds
}

func UnionSet(dj *DisjSet,root1 int,root2 int){
	if dj.S[root1] > dj.S[root2]{
		dj.S[root1] = root2
	}else{
		if dj.S[root1] == dj.S[root2]{
			dj.S[root1]--
		}
		dj.S[root2] = root1
	}
}

func Find(dj *DisjSet,e int) int{
	if dj.S[e]< 0 {
		return e
	}else{
		return Find(dj,dj.S[e])
	}
}
