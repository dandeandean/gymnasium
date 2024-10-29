package main

type MayHeap struct {
	data  []int
	IsMax bool
}

func (h MayHeap) Len() int            { return len(h.data) }
func (h MayHeap) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *MayHeap) Push(x interface{}) { h.data = append(h.data, x.(int)) }

func (h MayHeap) Peek() interface{} {
	return h.data[h.Len()-1]
}

func (h MayHeap) Less(i, j int) bool {
	if h.IsMax {
		return h.data[i] > h.data[j]
	}
	return h.data[i] < h.data[j]
}

func (h *MayHeap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func MaxHeap() *MayHeap {
	return &MayHeap{
		data:  make([]int, 0),
		IsMax: true,
	}
}

func MinHeap() *MayHeap {
	return &MayHeap{
		data:  make([]int, 0),
		IsMax: false,
	}
}
