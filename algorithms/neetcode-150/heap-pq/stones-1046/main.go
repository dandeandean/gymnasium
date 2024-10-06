package main

import "fmt"

func lastStoneWeight(stones []int) int {
	h := HeapFrom(stones)
	fmt.Println(h)
	for h.Len() > 1 {
		a := h.Pop()
		if h.Len() == 0 {
			return a
		}
		b := h.Pop()
		val := a - b
		println(a, b, val)
		if val < 0 {
			val = -val
		}
		if val != 0 {
			h.Push(val)
		}
	}
	if h.Len() > 0 {
		return h.Pop()
	}
	return 0
}

func main() {
	stones := []int{4, 3, 4, 3, 2}
	out := lastStoneWeight(stones)
	fmt.Println(out)
}
