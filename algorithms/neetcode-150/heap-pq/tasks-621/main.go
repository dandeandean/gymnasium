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

type timeyWimey struct {
	ToGo, CanStartAgain int
}

func leastInterval(tasks []byte, n int) int {
	h := &MaxHeap{}
	count := make(map[byte]int, 0)
	for _, t := range tasks {
		count[t]++
	}
	for _, c := range count {
		heap.Push(h, c)
	}
	t := 0
	q := make([]timeyWimey, 0)
	for h.Len() > 0 || len(q) > 0 {
		if h.Len() > 0 {
			toGo := heap.Pop(h).(int) - 1
			if toGo > 0 {
				canStartAgain := t + n
				q = append(q, timeyWimey{ToGo: toGo, CanStartAgain: canStartAgain})
			}
		}
		if len(q) > 0 && q[0].CanStartAgain <= t {
			heap.Push(h, q[0].ToGo)
			q = q[1:]
		}
		t++
	}
	return t
}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	fmt.Println(
		leastInterval(tasks, 1),
	)

}
