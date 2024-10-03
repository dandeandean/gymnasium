package main

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	for _, val := range stones {
		heap.Push(h, val)
	}
	for h.Len() > 1 {
		a := heap.Pop(h)
		b := heap.Pop(h)
		new := a - b
		heap.Push(h, new)
	}

	return 69
}
func main() {
	h := &MaxHeap{}
	heap.Init(h)
	heap.Fix(h, 0)
}
