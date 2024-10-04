package main

type MaxHeap []int

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Pop() int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *MaxHeap) Push(n int) {
	*h = append(*h, n)
	for i := len(*h) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
}

func HeapFrom(nums []int) *MaxHeap {
	heap := new(MaxHeap)
	for _, n := range nums {
		heap.Push(n)
	}
	return heap
}

func (h MaxHeap) MaxHeapify(i int) {
	left, right := i*2+1, i*2+2
	ibig := i
	if left < len(h) {
		if h[left] > h[ibig] {
			ibig = left
		}
	}
	if right < len(h) {
		if h[right] > h[ibig] {
			ibig = right
		}
	}
	if ibig != i {
		h.Swap(ibig, i)
		h.MaxHeapify(ibig)
	}
}
