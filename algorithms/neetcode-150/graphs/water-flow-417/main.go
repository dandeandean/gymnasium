package main

import "fmt"

func main() {
	mapa := [][]int{
		{4, 2, 7, 3, 4},
		{7, 4, 6, 4, 7},
		{6, 3, 5, 3, 6},
	}
	mapb := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	mapc := [][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	fmt.Println(
		pacificAtlantic(mapa),
		pacificAtlantic(mapb),
		pacificAtlantic(mapc),
	)
}

type coord [2]int

func pacificAtlantic(heights [][]int) [][]int {
	rowLen, colLen := len(heights), len(heights[0])
	bfs := func(q []coord) [][]bool {
		res := make([][]bool, rowLen)
		for i := range res {
			res[i] = make([]bool, colLen)
		}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			fmt.Println(cur)
			res[cur[0]][cur[1]] = true
			for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				cNew := coord{d[0] + cur[0], d[1] + cur[1]}
				valid := cNew[0] >= 0 && cNew[1] >= 0 && cNew[0] < len(heights) && cNew[1] < len(heights[0])
				if valid && heights[cNew[0]][cNew[1]] >= heights[cur[0]][cur[1]] && !res[cNew[0]][cNew[1]] {
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
		startPac = append(startPac, coord{0, c})
		startAtl = append(startAtl, coord{rowLen - 1, c})
	}
	// bfs from right & bottom (Atl)
	fmt.Println(startPac, startAtl)
	pacMap := bfs(startPac)
	atlMap := bfs(startAtl)
	fmt.Println(pacMap, atlMap)
	res := make([][]int, 0)
	for i := range pacMap {
		for j := range pacMap[i] {
			if pacMap[i][j] && atlMap[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
