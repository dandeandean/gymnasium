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
	var bfs func(r, c int) int
	bfs = func(r, c int) int {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[r]) {
			return -1
		}
		dirs := [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
		count := 0
		q := [][2]int{{r, c}}
		for len(q) > 0 {
			cr, cc := q[0][0], q[0][1]
			q = q[1:]
			for _, v := range dirs {
				rNew, cNew := cr+v[0], cc+v[1]
				isValid := cNew > 0 && cNew < len(grid) &&
					rNew > 0 && rNew < len(grid[cNew])
				if isValid && grid[rNew][cNew] != 0 {
					q = append(q, [2]int{rNew, cNew})
				}
			}
		}
		return count
	}
	res := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				size := bfs(r, c)
				if size > res {
					res = size
				}
			}
		}
	}
	return res
}
