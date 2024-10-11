package main

import (
	"container/heap"
	"fmt"
)

type Point struct {
	X      int
	Y      int
	DistSq int
}

func makePoint(x, y int) Point {
	return Point{
		X:      x,
		Y:      y,
		DistSq: x*x + y*y,
	}
}

func (p Point) Unpack() []int {
	return []int{p.X, p.Y}
}

type PQ []Point

func (h PQ) Len() int            { return len(h) }
func (h PQ) Less(i, j int) bool  { return h[i].DistSq < h[j].DistSq }
func (h PQ) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PQ) Push(x interface{}) { *h = append(*h, x.(Point)) }
func (h *PQ) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func kClosest(points [][]int, k int) [][]int {
	pq := &PQ{}
	for _, coord := range points {
		p := makePoint(coord[0], coord[1])
		heap.Push(pq, p)
	}
	out := make([][]int, k)
	for i := range out {
		out[i] = heap.Pop(pq).(Point).Unpack()
	}
	return out
}

func main() {
	a := []int{1, 3}
	b := []int{-2, 2}
	in := [][]int{a, b}
	fmt.Println(in)
	out := kClosest(in, 1)
	fmt.Println(out)
}
