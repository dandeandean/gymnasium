package main

import "fmt"

type Heap struct {
	Vals  []int
	IsMax bool
}

func (h Heap) Len() int      { return len(h.Vals) }
func (h Heap) Swap(i, j int) { h.Vals[i], h.Vals[j] = h.Vals[j], h.Vals[i] }

func (h Heap) cmp(a, b int) bool {
	if h.IsMax {
		if a >= b {
			return true
		}
		return false
	} else {
		if a <= b {
			return true
		}
	}
	return false
}

func (h Heap) Heapify(i int) {
	left, right := i*2+1, i*2+2
	maxima := i
	fmt.Println("ismax= ", h.IsMax)
	if left < h.Len() {
		if h.Vals[left] >= h.Vals[maxima] && h.IsMax {
			fmt.Println("switching")
			maxima = left
		}
		if h.Vals[left] <= h.Vals[maxima] && !h.IsMax {
			fmt.Println("switching")
			maxima = left
		}
	}
	if right < h.Len() {
		if h.Vals[right] >= h.Vals[maxima] && h.IsMax {
			fmt.Println("switching")
			maxima = right
		}
		if h.Vals[right] <= h.Vals[maxima] && !h.IsMax {
			fmt.Println("switching")
			maxima = right
		}
	}
	if maxima != i {
		h.Swap(maxima, i)
		h.Heapify(maxima)
	}
}
func (h *Heap) Pop() int {
	if h.Len() == 0 {
		panic("WHAT ARE YOU DOING CALLING POP ON AN EMPTY HEAP?")
	}
	old := *h
	x := old.Vals[0]
	if h.Len() <= 1 {
		// FIXME: used to be a pointer so this may mess up
		h.Vals = []int{}
	} else {
		h.Vals = old.Vals[1:]
	}
	for i := h.Len() / 2; i >= 0; i-- {
		h.Heapify(i)
	}
	return x
}
func (h *Heap) Push(n int) {
	h.Vals = append(h.Vals, n)
	for i := h.Len() / 2; i >= 0; i-- {
		h.Heapify(i)
	}
}

func (h *Heap) From(nums []int) {
	for _, n := range nums {
		h.Push(n)
	}
}

func MaxHeap() *Heap {
	return &Heap{IsMax: true}
}
func MinHeap() *Heap {
	return &Heap{IsMax: false}
}
