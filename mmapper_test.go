package main

import (
	"testing"
	"fmt"
)

func TestHello(t *testing.T) {

}

func ExampleHello() {
	fmt.Println("Hello")
	// Output:
	// Hello
}

func TestInit(t *testing.T) {
	data := map[int][]float64{0:{1.0,2.0}, 1:{2.3,3.2}, 2:{4.3,0.3}, 3:{1.2,2.2}}
	kd := KDTree{points: data}
	kd.Init()
}

func BenchmarkHello(b *testing.B) {
	fmt.Println("Hello")
}

func TestPartition(t *testing.T) {

	t.Run("KD with Data 1", func(t *testing.T) {
		data := map[int][]float64{0:{5}, 1:{3}, 2:{1}, 3:{-5}, 4:{0}, 5:{2}, 6:{7}}
		sorted := []int{ 3, 4, 2, 5, 1, 0, 6}
		kd := KDTree{ points: data }
		sb := kd.Init()
		p := sb.Partition(0,sb.Len()-1)
		if kd.Coor(kd.tree[p],0) != kd.Coor(sorted[p], 0) { t.Fail() }
	})

	t.Run("KD with Data 2", func(t *testing.T) {
		data := map[int][]float64{0:{1.0,2.0}, 1:{2.3,3.2}, 2:{4.3,0.3}, 3:{1.2,2.2}}
		sorted := []int{0, 3, 1, 2}
		kd := KDTree{ points: data }
		sb := kd.Init()
		p := sb.Partition(0,sb.Len()-1)
		if kd.Coor(kd.tree[p],0) != kd.Coor(sorted[p], 0) { t.Fail() }
	})

	t.Run("KD Data 2 Sort by 1", func(t * testing.T) {
		data := map[int][]float64{0:{1.0,2.0}, 1:{2.3,3.2}, 2:{4.3,0.3}, 3:{1.2,2.2}}
		sorted := []int{2, 0, 3, 1}
		kd := KDTree{ points: data }
		kd.dim = len(kd.points[0])
		kd.rows = len(kd.points)
		kd.tree = make([]int, len(kd.points), len(kd.points))
		for i:=0; i<len(kd.points); i++ { kd.tree[i] = i }
		sb := SortBy{ &kd, kd.tree, 1 }
		p := sb.Partition(0,sb.Len()-1)
		if kd.Coor(kd.tree[p],1) != kd.Coor(sorted[p], 1) { t.Fail() }
	})

}

func TestSelect(t * testing.T) {

	t.Run("KD with Data 1", func(t *testing.T) {
		data := map[int][]float64{0:{5}, 1:{3}, 2:{1}, 3:{-5}, 4:{0}, 5:{2}, 6:{7}}
		sorted := []int{ 3, 4, 2, 5, 1, 0, 6}
		kd := KDTree{ points: data }
		sb := kd.Init()
		for i:=0; i<sb.Len(); i++ {
			s := sb.Select(0,sb.Len()-1, i)
			if kd.Coor(kd.tree[s],0) != kd.Coor(sorted[s], 0) {
				t.Fail()
				fmt.Println(":(", i)
			}
		}
	})

	t.Run("KD with Data 2", func(t *testing.T) {
		data := map[int][]float64{0:{1.0,2.0}, 1:{2.3,3.2}, 2:{4.3,0.3}, 3:{1.2,2.2}}
		sorted := []int{0, 3, 1, 2}
		kd := KDTree{ points: data }
		sb := kd.Init()
		for i:=0; i<sb.Len(); i++ {
			s := sb.Select(0,sb.Len()-1, i)
			if kd.Coor(kd.tree[s],0) != kd.Coor(sorted[s], 0) { t.Fail() }
		}
	})

	t.Run("KD Data 2 Sort by 1", func(t * testing.T) {
		data := map[int][]float64{0:{1.0,2.0}, 1:{2.3,3.2}, 2:{4.3,0.3}, 3:{1.2,2.2}}
		sorted := []int{2, 0, 3, 1}
		kd := KDTree{ points: data }
		kd.dim = len(kd.points[0])
		kd.rows = len(kd.points)
		kd.tree = make([]int, len(kd.points), len(kd.points))
		for i:=0; i<len(kd.points); i++ { kd.tree[i] = i }
		sb := SortBy{ &kd, kd.tree, 1 }
		for i:=0; i<sb.Len(); i++ {
			s := sb.Select(0,sb.Len()-1, i)
			if kd.Coor(kd.tree[s],1) != kd.Coor(sorted[s], 1) { t.Fail() }
		}
	})

}
