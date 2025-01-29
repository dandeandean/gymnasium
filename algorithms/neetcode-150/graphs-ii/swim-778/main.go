package main

import (
	"container/heap"
	"fmt"
)

type cell struct {
	x, y, t int
}

type MinHeap []cell

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(cell)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func swimInWater(grid [][]int) int {
	bounds := len(grid)
	h := &MinHeap{cell{x: 0, y: 0, t: grid[0][0]}}
	heap.Init(h)
	beenTo := make(map[[2]int]bool)
	beenTo[[2]int{0, 0}] = true
	for h.Len() > 0 {
		cur := heap.Pop(h).(cell)
		if cur.x == bounds-1 && cur.y == bounds-1 {
			return cur.t
		}
		for _, d := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			newCell := cell{x: cur.x + d[0], y: cur.y + d[1]}
			outOfBounds := newCell.x >= bounds || newCell.y >= bounds || newCell.y < 0 || newCell.x < 0
			if outOfBounds || beenTo[[2]int{newCell.x, newCell.y}] {
				continue
			}
			beenTo[[2]int{newCell.x, newCell.y}] = true
			newCell.t = grid[newCell.x][newCell.y]
			if cur.t > newCell.t {
				newCell.t = cur.t
			}
			heap.Push(h, newCell)
		}
	}
	return -1
}

func main() {
	g := [][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}
	fmt.Println(swimInWater(g))
}
