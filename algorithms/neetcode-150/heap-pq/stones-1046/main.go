package main

import "fmt"

func lastStoneWeight(stones []int) int {
	h := HeapFrom(stones)
	return h.Pop()
}
func main() {
	stones := []int{1, 2, 3, 5}
	h := HeapFrom(stones)
	fmt.Println(h)

}
