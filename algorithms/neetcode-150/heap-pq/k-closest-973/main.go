package main

import "fmt"

func main() {
	h := MaxHeap()
	h.From([]int{9, 2, 8, 5})
	fmt.Println(h.Vals)
	h2 := MinHeap()
	h2.From([]int{9, 2, 8, 5})
	fmt.Println(h2.Vals)
	fmt.Println(
		h.IsMax,
		h2.IsMax,
	)
}
