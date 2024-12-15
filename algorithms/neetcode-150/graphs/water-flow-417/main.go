package main

import "fmt"

func main() {
	mapa := [][]int{
		{4, 2, 7, 3, 4},
		{7, 4, 6, 4, 7},
		{6, 3, 5, 3, 6},
	}
	fmt.Println(
		pacificAtlantic(mapa),
	)
}

type coord [2]int

func pacificAtlantic(heights [][]int) [][]int {
	rowLen, colLen := len(heights), len(heights[0])
	bfs := func(startingCoords []coord) [][]bool {
		res := make([][]bool, rowLen)
		for i := range res {
			res[i] = make([]bool, colLen)
		}
		fmt.Println(res)
		q := startingCoords
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			res[cur[0]][cur[1]] = true
			for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				cNew := coord{d[0] + cur[0], d[1] + cur[1]}
				valid := cNew[0] >= 0 && cNew[1] >= 0 && cNew[0] < len(heights) && cNew[1] < len(heights[0])
				if valid && heights[cNew[0]][cNew[1]] >= heights[cur[0]][cur[1]] {
					q = append(q, cNew)
				}
			}
		}
		return res
	}
	// bfs from left & top (Pac)
	startPac := make([]coord, 0)
	startAtl := make([]coord, 0)
	for r := range rowLen {
		startPac = append(startPac, coord{r, 0})
		startAtl = append(startAtl, coord{r, colLen - 1})
	}
	for c := range colLen {
		startPac = append(startPac, coord{c, 0})
		startAtl = append(startAtl, coord{rowLen - 1, c})
	}
	// bfs from right & bottom (Atl)
	pacMap := bfs(startPac)
	atlMap := bfs(startAtl)
	res := make([][]int, 0)
	for i := range pacMap {
		for j := range pacMap {
			if pacMap[i][j] && atlMap[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
