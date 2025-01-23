package main

import "container/heap"

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Point struct {
	n int
	x int
	y int
}

func minCostConnectPoints(points [][]int) int {
	frontier := &IntHeap{}
	heap.Init(frontier)
	pRepr := make([]Point, 0)
	for n, p := range points {
		pNew := Point{
			n: n,
			x: p[0],
			y: p[1],
		}
		pRepr = append(pRepr, pNew)
	}
	//beenTo := make(map[int]bool)
	//aList := make(map[Point][]Point)
	//Prims alg
	// loop over the outgoing edges
	return points[0][0]
}
