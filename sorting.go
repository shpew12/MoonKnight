package main

import (
	"math/rand"
)

// Randomized Partition
func (sb * SortBy) Partition(lbd, ubd int) int {
	sb.Swap(ubd, rand.Intn(ubd-lbd)+lbd)
	i := lbd-1
	for j:=lbd; j<ubd; j++ {
		if ! sb.Less(sb.s[ubd], sb.s[j]) {
			i++
			sb.Swap(i,j)
		}
	}
	sb.Swap(i+1, ubd)
	return i+1
}

// Randomized Select
func (sb * SortBy) Select(lbd, ubd, i int) int {
	if lbd == ubd { return sb.s[lbd] }
	pivot := sb.Partition(lbd, ubd)
	left := pivot - lbd
	if i == left {
		return sb.s[pivot]
	} else if i < left {
		return sb.Select(lbd, pivot-1, i)
	} else {
		return sb.Select(pivot+1, ubd, i-left)
	}
}

// Struct and methods for sort Interface
type SortBy struct {
	kd * KDTree
	s []int
	by int
}

func (sb SortBy) Len() int {
	return len(sb.s)
}

func (sb SortBy) Swap(i,j int) {
	sb.s[i],sb.s[j] = sb.s[j], sb.s[i]
}

func (sb SortBy) Less(i,j int) bool {
	return sb.kd.Coor(i,sb.by) < sb.kd.Coor(j, sb.by)
}
