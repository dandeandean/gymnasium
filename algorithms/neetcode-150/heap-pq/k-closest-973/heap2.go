package main

type Heap struct {
	Vals  []int
	IsMax bool
}

func (h Heap) Len() int      { return len(h.Vals) }
func (h Heap) Swap(i, j int) { h.Vals[i], h.Vals[j] = h.Vals[j], h.Vals[i] }

func genCmpEq(a, b int, gt bool) bool {
	if (a >= b) && gt {
		return true
	}
	if (a <= b) && !gt {
		return true
	}
	return false
}

func (h Heap) Heapify(i int, isMax bool) {
	left, right := i*2+1, i*2+2
	ibig := i
	if left < h.Len() {
		if genCmpEq(h.Vals[left], h.Vals[ibig], isMax) {
			ibig = left
		}
	}
	if right < h.Len() {
		if genCmpEq(h.Vals[right], h.Vals[ibig], isMax) {
			ibig = right
		}
	}
	if ibig != i {
		h.Swap(ibig, i)
		h.Heapify(ibig, isMax)
	}
}
func (h *Heap) Pop() int {
	if h.Len() == 0 {
		panic("WHAT ARE YOU DOING CALLING POP ON AN EMPTY HEAP?")
	}
	old := *h
	x := old.Vals[0]
	if h.Len() <= 1 {
		//FIXME: used to be a pointer so this may mess up
		h.Vals = []int{}
	} else {
		h.Vals = old.Vals[1:]
	}
	for i := h.Len() / 2; i >= 0; i-- {
		h.Heapify(i, h.IsMax)
	}
	return x
}
func (h *Heap) Push(n int) {
	h.Vals = append(h.Vals, n)
	for i := h.Len() / 2; i >= 0; i-- {
		h.Heapify(i, h.IsMax)
	}
}

func HeapFrom(nums []int, isMax bool) *Heap {
	heap := new(Heap)
	for _, n := range nums {
		heap.Push(n)
	}
	heap.IsMax = isMax
	return heap
}
