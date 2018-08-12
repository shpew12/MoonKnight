package main

import (
	"fmt"
	"sort"
)

type KDTree struct {
	points map[int][]float64
	sweeps [][]int
	dim, rows int
}

func (kd *KDTree) Init() {
	kd.dim = len(kd.points[0])
	kd.rows = len(kd.points)
	kd.InitSweeps()
}

func (kd *KDTree) InitSweeps() {
	n := len(kd.points)
	kd.sweeps = make([][]int,kd.dim,kd.dim)
	for i:=0; i<kd.dim;i++ {
		kd.sweeps[i] = make([]int,n,n)
		for j:=0; j<kd.rows; j++ {
			kd.sweeps[i][j] = j
		}
		sb := SweepBy{kd, kd.sweeps[i], i}
		sort.Sort(sb)
	}
}

func (kd *KDTree) Coor(pid, dim int) float64 {
	return kd.points[pid][dim]
}

func (kd *KDTree) Find(val float64, /*lpid, upid,*/ dim int) int {
	sb := SweepBy{kd, kd.sweeps[dim]/*[lpid:upid]*/, dim}
	return sort.Search(sb.Len(), func(i int)bool { return kd.Coor(i, dim) >= val })
}

// Struct and methods for sort Interface
type SweepBy struct {
	kd * KDTree
	s []int
	by int
}

func (sb SweepBy) Len() int {
	return len(sb.s)
}

func (sb SweepBy) Swap(i,j int) {
	sb.s[i],sb.s[j] = sb.s[j], sb.s[i]
}

func (sb SweepBy) Less(i,j int) bool {
	return sb.kd.Coor(i,sb.by) < sb.kd.Coor(j, sb.by)
}

func main(){
	data := map[int][]float64{0:{1.0,2.0}, 1:{2.3,3.2}, 2:{4.3,0.3}, 3:{1.2,2.2}}
	fmt.Println(data)
	kd := KDTree{points: data}
	kd.Init()
	fmt.Println(kd)
	for _, sweep := range kd.sweeps {
		fmt.Println(sweep)
	}
	fmt.Println(kd.Find(1.5,1))
}
