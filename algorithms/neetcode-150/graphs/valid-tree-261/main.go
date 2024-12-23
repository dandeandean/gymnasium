package main

import "fmt"

func main() {
	fmt.Println(
		validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}),
		validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}),
	)
	return
}

func validTree(n int, edges [][]int) bool {
	return n > len(edges)
}
