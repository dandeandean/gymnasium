package main

import (
	"fmt"
)

func main() {
	g := [][]byte{
		{'1', '1', '0', '0'},
		{'0', '1', '1', '0'},
		{'0', '0', '0', '0'},
		{'0', '1', '1', '0'},
	}
	fmt.Println(
		numIslands(g),
	)
}

type coord struct {
	row  int
	col  int
	char byte
}
type grid [][]byte

func newGrid(g [][]byte) grid {
	return g
}

func (g grid) coordFrom(r, c int) *coord {
	if r < 0 || c < 0 || r >= len(g) || c >= len(g[0]) {
		return nil
	}
	return &coord{
		row:  r,
		col:  c,
		char: g[r][c],
	}
}

func (g grid) getNeighbors(c coord) []coord {
	dir := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	res := make([]coord, 0)
	for _, v := range dir {
		dx, dy := v[0], v[1]
		n := g.coordFrom(c.row+dx, c.col+dy)
		if n != nil {
			res = append(res, *n)
		}
	}
	return res
}
func (c coord) isLand() bool {
	return c.char == '1'
}
func (g *grid) bfs(r, c int) {
	q := make([]*coord, 0)
	q = append(q, g.coordFrom(r, c))
	fmt.Println(g)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, n := range g.getNeighbors(*cur) {
			if n.isLand() {
				q = append(q, &n)
				(*g)[n.row][n.col] = '0'
			}
		}
	}
	fmt.Println(g)
}

func numIslands(grid [][]byte) int {
	g := newGrid(grid)
	count := 0
	for r := range g {
		for c := range g[r] {
			if g[r][c] == '1' {
				g.bfs(r, c)
				count++
			}
		}
	}
	return count
}
