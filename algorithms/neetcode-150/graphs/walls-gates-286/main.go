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
	g := newGrid(grid)
	for r := range grid {
		for c := range grid[r] {
			co := coord{r, c}
			if g.isTreasure(co) {
				q.push(co)
			}
		}
	}
	dirs := [4]coord{
		{-1, 0}, {0, -1},
		{0, 1}, {1, 0},
	}
	for len(q) > 0 {
		co := q.pop()
		for _, d := range dirs {
			co2 := coord{
				co[0] + d[0],
				co[1] + d[1],
			}
			fmt.Println(co, co2)
			if g.canWalk(co2) {
				q.push(co2)
				g.set(co2, g[co[0]][co[1]]+1)
			}
		}
	}
	g.show()
}

var (
	EMPTY = 2147483647
	WALL  = -1
	GOAL  = 0
)

type grid [][]int
type coord [2]int
type qu []coord

func newGrid(g [][]int) grid {
	return g
}

func (g grid) show() {
	for r := range g {
		fmt.Println(g[r])
	}
}

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
	return g.isValid(c) && g[c[0]][c[1]] != WALL
}

func (g grid) isTreasure(c coord) bool {
	return g.isValid(c) && g[c[0]][c[1]] == GOAL
}

func (g grid) get(c coord) int {
	if g.isValid(c) {
		fmt.Println("Can't get", c, " in")
		g.show()
		panic("Invalid coord!")
	}
	return g[c[0]][c[1]]
}

func (g *grid) set(c coord, n int) bool {
	if !g.canWalk(c) {
		return false
	}
	(*g)[c[0]][c[1]] = n
	return (*g)[c[0]][c[1]] == n
}

func (q *qu) push(n coord) {
	*q = append(*q, n)
}
func (q *qu) pop() coord {
	out := (*q)[0]
	*q = (*q)[1:]
	return out
}
