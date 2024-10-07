package main

import "fmt"

func main() {
	h := HeapFrom([]int{9, 2, 8, 5}, false)
	fmt.Println(h.Vals)
}
