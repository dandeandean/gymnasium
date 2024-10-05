package main

import "fmt"

func lastStoneWeight(stones []int) int {
	h := HeapFrom(stones)
	fmt.Println(h)
	for h.Len() > 1 {
		a := h.Pop()
		if h.Len() <= 1 {
			return a
		}
		b := h.Pop()
		val := a - b
		if val < 0 {
			val = -val
		}
		if val != 0 {
			h.Push(val)
		}
	}
	return h.Pop()
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	out := lastStoneWeight(stones)
	out2 := lastStoneWeight([]int{1})
	fmt.Println(out)
	fmt.Println(out2)
}
