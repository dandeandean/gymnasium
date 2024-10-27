package main

import (
	"container/heap"
	"fmt"
)

func main() {
	maxh := MaxHeap()
	minh := MinHeap()
	for i := range 10 {
		heap.Push(maxh, i)
		heap.Push(minh, i)
	}
	fmt.Println(maxh.data, minh.data)
}
