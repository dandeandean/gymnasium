package main

import "fmt"

func main() {
	g := [][]int{
		{2, 1, 1},
		{1, 1, 1},
		{0, 1, 2},
	}
	g1 := [][]int{
		{2, 1, 0},
		{0, 1, 1},
		{0, 0, 1},
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

func (g graph) show() {
	for r := range g {
		fmt.Println(g[r])
	}
}

func orangesRotting(grid [][]int) int {
	g := graph(grid)
	q := make([]coord, 0)
	for r := range g {
		for c := range g[r] {
			if g[r][c] == 2 {
				q = append(q, coord{r, c})
			}
		}
	}
	dirs := [4]coord{
		{-1, 0}, {0, -1},
		{0, 1}, {1, 0},
	}
	res := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		wrote := false
		for _, d := range dirs {
			newCoord := coord{cur[0] + d[0], cur[1] + d[1]}
			if g.isFresh(newCoord) {
				g.drainLiveForce(newCoord)
				q = append(q, newCoord)
				wrote = true
			}
		}
		if !wrote {
			return res
		}
		res++
		fmt.Println(res)
		g.show()
	}
	for r := range g {
		for c := range g[r] {
			if g[r][c] == 1 {
				return -1
			}
		}
	}
	return res
}
