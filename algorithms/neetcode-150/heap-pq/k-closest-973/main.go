package main

import "fmt"

type Coord []int

func (c Coord) SquardDist() int {
	X := c[0]
	Y := c[1]
	return X*X + Y*Y
}
func (c Coord) Slice() []int {
	return []int{c[0], c[1]}
}

func kClosest(points [][]int, k int) [][]int {
	h := MinHeap()
	m := make(map[int]Coord)
	for _, point := range points {
		c := make(Coord, 2)
		c[0], c[1] = point[0], point[1]
		m[c.SquardDist()] = c
		h.Push(c.SquardDist())
	}
	out := make([][]int, k)
	for i := 0; i < k; i++ {
		sq := h.Pop()
		// this may just unpack the slice???
		out = append(out, m[sq].Slice())
	}
	return out
}

func main() {
	fmt.Println("")
}
