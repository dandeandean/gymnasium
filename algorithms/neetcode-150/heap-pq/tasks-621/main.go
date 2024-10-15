package main

import (
	"container/heap"
	"fmt"
)

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
func leastInterval(tasks []byte, n int) int {
	h := &MaxHeap{}
	//q := make([]int, 0)
	count := make(map[byte]int, 0)
	for _, t := range tasks {
		count[t]++
	}
	for _, c := range count {
		heap.Push(h, c)
	}
	fmt.Println(count, n)
	fmt.Println(h, n)
	return 42
}

func main() {
	tasks := []byte{'A', 'B', 'A'}
	leastInterval(tasks, 1)

}
