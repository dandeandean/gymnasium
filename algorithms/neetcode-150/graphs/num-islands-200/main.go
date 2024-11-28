package main

import "fmt"

func main() {
	g := [][]byte{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(numIslands(g))
}

type node struct {
	isLand    bool
	neighbors [4]*node
}

func numIslands(grid [][]byte) int {
	return 42
}
