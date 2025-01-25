package main

import "container/heap"

// An IntHeap is a min-heap of ints.
type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.([]int)) }
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

func PointsFrom(points [][]int) []Point {
	pRepr := make([]Point, 0)
	for n, p := range points {
		pNew := Point{
			n: n,
			x: p[0],
			y: p[1],
		}
		pRepr = append(pRepr, pNew)
	}
	return pRepr
}

func (p Point) distTo(p2 Point) int {
	a := p.x - p2.x
	if a < 0 {
		a = -a
	}
	b := p.y - p2.y
	if b < 0 {
		b = -b
	}
	return a + b
}

func minCostConnectPoints(points [][]int) int {
	frontier := &IntHeap{}
	heap.Init(frontier)
	pRepr := PointsFrom(points)
	aList := make(map[Point]*IntHeap)
	for a, p1 := range pRepr {
		for b, p2 := range pRepr {
			dist := p1.distTo(p2)
			heap.Push(aList[pRepr[a]], []int{dist, p2.n})
			heap.Push(aList[pRepr[b]], []int{dist, p1.n})
		}
	}
	//Prims alg
	beenTo := make(map[int]bool)
	// loop over the outgoing edges
	cur := &pRepr[0]
	for cur != nil {
		beenTo[cur.n] = true
		outgoingEdge := heap.Pop(aList[*cur]).([]int)
		for beenTo[outgoingEdge[1]] {
			outgoingEdge = heap.Pop(aList[*cur]).([]int)
		}
		cur = &pRepr[outgoingEdge[1]]
	}
	return points[0][0]
}
