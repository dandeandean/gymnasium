package main

import "fmt"

func kClosest(points [][]int, k int) [][]int {
	h := MinHeap()
	for _, point := range points {
		c := make(Coord, 2)
		c[0], c[1] = point[0], point[1]
		h.Push(c)
	}
	out := make([][]int, k)
	for i := range out {
		out[i] = h.Pop()
	}
	return out
}

func main() {
	a := []int{1, 3}
	b := []int{-2, 2}
	in := [][]int{a, b}
	fmt.Println(in)
	out := kClosest(in, 1)
	fmt.Println(out)
}
