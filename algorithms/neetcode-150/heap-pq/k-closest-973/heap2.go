package main

type Coord []int

func (c Coord) SquardDist() int {
	X := c[0]
	Y := c[1]
	return X*X + Y*Y
}

func (c Coord) Slice() []int {
	return []int{c[0], c[1]}
}

type Heap struct {
	Vals  []Coord
	IsMax bool
}

func (h Heap) Len() int      { return len(h.Vals) }
func (h Heap) Swap(i, j int) { h.Vals[i], h.Vals[j] = h.Vals[j], h.Vals[i] }

func (h Heap) cmp(a, b int) bool {
	if h.IsMax {
		if h.Vals[a].SquardDist() >= h.Vals[b].SquardDist() {
			return true
		}
		return false
	} else {
		if h.Vals[a].SquardDist() <= h.Vals[b].SquardDist() {
			return true
		}
	}
	return false
}

func (h Heap) Heapify(i int) {
	left, right := i*2+1, i*2+2
	maxima := i
	if left < h.Len() {
		if h.cmp(left, maxima) {
			maxima = left
		}
	}
	if right < h.Len() {
		if h.cmp(right, maxima) {
			maxima = right
		}
	}
	if maxima != i {
		h.Swap(maxima, i)
		h.Heapify(maxima)
	}
}
func (h *Heap) Pop() Coord {
	if h.Len() == 0 {
		panic("WHAT ARE YOU DOING CALLING POP ON AN EMPTY HEAP?")
	}
	old := *h
	x := old.Vals[0]
	if h.Len() <= 1 {
		h.Vals = []Coord{}
	} else {
		h.Vals = old.Vals[1:]
	}
	for i := h.Len() / 2; i >= 0; i-- {
		h.Heapify(i)
	}
	return x
}
func (h *Heap) Push(n Coord) {
	h.Vals = append(h.Vals, n)
	for i := h.Len() / 2; i >= 0; i-- {
		h.Heapify(i)
	}
}

func (h *Heap) From(nums []Coord) {
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
