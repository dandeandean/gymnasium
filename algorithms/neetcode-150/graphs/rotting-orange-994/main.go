package main

import "fmt"

func main() {
	g := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	g1 := [][]int{
		{2, 1, 0},
		{0, 1, 1},
		{1, 0, 1},
	}
	fmt.Println(
		orangesRotting(g),
		orangesRotting(g1),
	)
}

type coord [2]int
type graph [][]int

func (g graph) isValid(c coord) bool {
	if c[0] < 0 || c[1] < 0 || c[0] >= len(g) || c[1] >= len(g[0]) {
		return false
	}
	return true
}

func (g graph) isFresh(c coord) bool {
	return g.isValid(c) && g[c[0]][c[1]] == 1
}
func (g *graph) drainLiveForce(c coord) {
	(*g)[c[0]][c[1]] = 2
}

func (g *graph) bfs(r, c int) int {
	q := []coord{{r, c}}
	dirs := [4]coord{
		{-1, 0}, {0, -1},
		{0, 1}, {1, 0},
	}
	res := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, d := range dirs {
			newCoord := coord{cur[0] + d[0], cur[1] + d[1]}
			if g.isFresh(newCoord) {
				g.drainLiveForce(newCoord)
				q = append(q, newCoord)
			}
		}
		res++
	}
	return res
}

func orangesRotting(grid [][]int) int {
	count := -1
	g := graph(grid)
	for r := range g {
		for c := range g[r] {
			if g[r][c] == 2 {
				rotting := g.bfs(r, c)
				if rotting > count {
					count = rotting
				}
			}
		}
	}
	fmt.Println(g)
	return count
}
