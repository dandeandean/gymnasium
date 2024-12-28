package main

import "fmt"

func main() {
	fmt.Println(
		findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}),
		findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}),
	)
}

func findRedundantConnection(edges [][]int) []int {
	return edges[0]
}
