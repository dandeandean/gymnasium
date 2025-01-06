package main

import (
	"container/heap"
	"fmt"
)

func main() {
	g := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	fmt.Println(networkDelayTime(g, 4, 2))
}

type heapPair struct {
	nodeId int
	dist   int
}
type minHeap []heapPair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(heapPair))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	edges := make(map[int][][]int, 0)
	for _, t := range times {
		edges[t[0]] = append(edges[t[0]], []int{t[1], t[2]})
	}
	mh := &minHeap{heapPair{k, 0}}
	heap.Init(mh)
	beenTo := make([]bool, n+1)
	counter := 0
	for mh.Len() > 0 {
		cur := heap.Pop(mh).(heapPair)
		curNode, curDist := cur.nodeId, cur.dist
		if beenTo[curNode] {
			continue
		}
		if curDist > counter {
			counter = curDist
		}
		beenTo[curNode] = true
		for _, nei := range edges[cur.nodeId] {
			heap.Push(mh, heapPair{nei[0], curDist + nei[1]})
		}
	}
	fmt.Println(beenTo)
	for _, v := range beenTo[1:] {
		if !v { // return -1 if not visited
			return -1
		}
	}
	return counter
}
