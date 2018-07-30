package main

import (
	"fmt"
	"sort"
//	"math"
)

type Node struct {
	dim int
//	lbd, ubd float64
	lpid, hpid int
	parent, left, right *Node
}

func (node *Node) Up(d int) *Node {
	if node.parent == nil { return nil }
	if d == 1 { return node.parent }
	return node.parent.Up(d-1)
}

type KDTree struct {
	root Node
	points map[int][]float64
	sweeps [][]int
	dim, rows int
}

func (kd *KDTree) Init() {
	kd.dim = len(kd.points[0])
	kd.rows = len(kd.points)
	kd.InitSweeps()
	kd.InitTree()
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

/*
func (kd *KDTree) Bounds() ([]float64,[]float64) {
	n := len(kd.points[0])
	mins, maxs := make([]float64, n, n), make([]float64, n, n)
	copy(mins, kd.points[0])
	copy(maxs, kd.points[0])
	for _, pt := range kd.points {
		for j, coor := range pt {
			mins[j] = math.Min(coor, mins[j])
			maxs[j] = math.Max(coor, maxs[j])
		}
	}
	return mins, maxs
}
*/

func (kd *KDTree) InitTree() {
//	mins, maxs:= kd.Bounds()
	kd.root = Node{/*lbd: mins[0], ubd: maxs[0],*/ lpid: kd.sweeps[0][0], hpid: kd.sweeps[0][len(kd.sweeps[0])-1]}
	kd.Split(&kd.root)
}

func (kd *KDTree) SplitTo(node *Node, weight int) {
	//TODO
}

func (kd *KDTree) Split(node *Node) {
	d := node.dim+1
	if d == kd.dim { d = 0 }
	//mid := kd.Find(node.ubd/2, node.lpid, node.hpid, d)
	mid := kd.Find(kd.Coor(node.lpid, d)/2, node.lpid, node.hpid, d)
	node.left = &Node{d, /*node.lbd, node.ubd/2,*/ node.lpid, mid, node, nil, nil}
	node.right = &Node{d,/* node.ubd/2, node.ubd,*/ mid, node.hpid, node, nil, nil}
}

func (kd *KDTree) Coor(pid, dim int) float64 {
	return kd.points[pid][dim]
}

func (kd *KDTree) Find(val float64, lpid, hpid, dim int) int {
	sb := SweepBy{kd, kd.sweeps[dim][lpid:hpid], dim}
	return sort.Search(sb.Len(), func(i int)bool { return kd.Coor(i, dim) >= val })
}

func (kd *KDTree) Weight(node *Node) int {
	//TODO
	return 0
}

// struct and methods for sort Interface
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
	fmt.Println(kd.root.left, kd.root.right)
}
