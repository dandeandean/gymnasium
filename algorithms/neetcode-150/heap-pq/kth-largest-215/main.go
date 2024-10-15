package main

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func findKthLargest(nums []int, k int) int {
	mh := &MaxHeap{}
	for _, n := range nums {
		heap.Push(mh, n)
	}
	for i := 0; i < k-1; i++ {
		heap.Pop(mh)
	}
	out := heap.Pop(mh).(int)
	return out
}
