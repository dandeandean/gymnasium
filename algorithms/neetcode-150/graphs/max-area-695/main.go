package main

import "fmt"

func main() {
	g := [][]int{
		{1, 1, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 1, 1, 0},
	}
	fmt.Println(
		maxAreaOfIsland(g),
	)
}

func maxAreaOfIsland(grid [][]int) int {
	type coord [2]int
	var dfsCount func(r, c int) int
	beenTo := make(map[coord]bool)
	dfsCount = func(r, c int) int {
		if r < 0 || c < 0 ||
			r >= len(grid) || c >= len(grid[r]) ||
			grid[r][c] == 0 || beenTo[coord{r, c}] {
			return 0
		}
		beenTo[coord{r, c}] = true
		return 1 +
			dfsCount(r+1, c) +
			dfsCount(r-1, c) +
			dfsCount(r, c+1) +
			dfsCount(r, c-1)
	}
	res := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				size := dfsCount(r, c)
				if size > res {
					res = size
				}
			}
		}
	}
	return res
}
