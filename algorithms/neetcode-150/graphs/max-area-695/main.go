package main

func maxAreaOfIsland(grid [][]int) int {
	var bfs func(r, c int) int
	bfs = func(r, c int) int {
		dirs := [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[r]) {
			return -1
		}
		q := make([][2]int, 0)
		count := 0
		for len(q) > 0 {
			cr, cc := q[0][0], q[0][1]
			q = q[1:]
			for _, v := range dirs {
				dx, dy := v[0], v[1]
				cNew := cr + dx
				rNew := cc + dy
				if cNew > 0 && cNew < len(grid) && rNew > 0 && rNew < len(grid[cNew]) && grid[cNew][rNew] == '1'{
						grid[cNew][rNew] = '0'
						count += 1
					}
				}
			}
		}
		return count
	}
	res := 0
	for r := range grid {
		for c := range grid[r] {
			size := bfs(r, c)
			if size > res {
				res = size
			}
		}
	}
	return res
}
