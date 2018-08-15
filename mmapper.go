package main

import (
	"fmt"
)

type KDTree struct {
	points map[int][]float64
	tree []int
	dim, rows int
}

func (kd *KDTree) Init() *SortBy {
	kd.dim = len(kd.points[0])
	kd.rows = len(kd.points)
	kd.tree = make([]int, len(kd.points), len(kd.points))
	for i:=0; i<len(kd.points); i++ { kd.tree[i] = i }
	sb := SortBy{ kd, kd.tree, 0 }
	sb.MakeTree(0, sb.Len()-1)
	return &sb
}

func (sb * SortBy) MakeTree(lbd, ubd int) {
	if lbd == ubd { return }
	mid := (lbd+ubd+1)/2
	sb.Select(lbd, ubd, mid )
	if sb.by == sb.kd.dim-1 {
		sb.by = 0
	} else { sb.by++ }
	sb.MakeTree(lbd, mid-1)
	sb.MakeTree(mid, ubd)
}

func (kd *KDTree) Coor(pid, dim int) float64 {
	return kd.points[pid][dim]
}

func main(){
	data := map[int][]float64{0:{1.0,2.0}, 1:{2.3,3.2}, 2:{4.3,0.3}, 3:{1.2,2.2}}
	fmt.Println(data)
	kd := KDTree{points: data}
	kd.Init()
	fmt.Println(kd)
}
