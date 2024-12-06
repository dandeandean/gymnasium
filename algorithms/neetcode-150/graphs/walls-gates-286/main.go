package main

import (
	"fmt"
)

func main() {
	mapa := [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}

	mapb := [][]int{
		{0, -1},
		{2147483647, 2147483647},
	}
	islandsAndTreasure(mapa)
	islandsAndTreasure(mapb)
	//fmt.Println(mapa, mapb)
}

func islandsAndTreasure(grid [][]int) {
	q := make(qu, 0)
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 0 {
				q.push(coord{r, c})
			}
		}
	}
	fmt.Println(q)

}

var (
	EMPTY = 2147483647
	WALL  = -1
	GOAL  = 0
)

type grid [][]int
type coord [2]int
type qu []coord

func (g grid) isValid(c coord) bool {
	if c[0] >= len(g) || c[0] < 0 {
		return false
	}
	if c[1] >= len(g[0]) || c[1] < 0 {
		return false
	}
	return true
}

func (g grid) canWalk(c coord) bool {
	return g.isValid(c) && g[c[0]][c[1]] == GOAL
}

func (g grid) isTreasure(c coord) bool {
	return g.isValid(c) && g[c[0]][c[1]] == GOAL
}

func (g *grid) set(c coord, n int) bool {
	if !g.canWalk(c) {
		(*g)[c[0]][c[1]] = n
	}
	return (*g)[c[0]][c[1]] == n
}

func (q *qu) push(n coord) *qu {
	*q = append(*q, n)
	return q
}
func (q *qu) pop() (coord, *qu) {
	out := (*q)[0]
	*q = (*q)[1:]
	return out, q
}
