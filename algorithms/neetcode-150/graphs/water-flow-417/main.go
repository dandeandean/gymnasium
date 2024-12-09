package main

import "fmt"

type cell struct {
	c   coord
	h   int
	alt bool
	pac bool
}

type grid struct {
	heights  [][]cell
	pacDrain []coord
	atlDrain []coord
}

type coord [2]int

func gridFromHeights(heights [][]int) grid {
	g := grid{}
	for r := range heights {
		row := make([]cell, len(heights[r]))
		for c := range heights[r] {
			row[c] = cell{
				c: coord{r, c},
				h: heights[r][c],
			}
		}
		g.heights = append(g.heights, row)
	}
	return g
}

func (g grid) inBounds(c coord) bool {
	return c[0] > 0 && c[0] < len(g.heights) && c[1] > 0 && c[1] < len(g.heights[0])
}

func (g grid) heightAt(c coord) int {
	if g.inBounds(c) == false {
		return -1
	}
	return g.heights[c[0]][c[1]].h
}

func (g *grid) setAtl(c coord) {
	if g.inBounds(c) {
		(*g).heights[c[0]][c[1]].alt = true
		(*g).atlDrain = append((*g).atlDrain, c)
	}
}

func (g *grid) setPac(c coord) {
	if g.inBounds(c) {
		(*g).heights[c[0]][c[1]].pac = true
		(*g).pacDrain = append((*g).pacDrain, c)
	}
}

func (g grid) get(c coord) *cell {
	if !g.inBounds(c) {
		return nil
	}
	return &g.heights[c[0]][c[1]]
}

func (g *grid) bfs(c coord, isAtl bool) {
	var writeToG func(coord)
	if isAtl {
		writeToG = g.setAtl
	} else {
		writeToG = g.setPac
	}
	q := []coord{c}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, v := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			cNew := coord{v[0] + cur[0], v[1] + cur[1]}
			cellNew := g.get(cNew)
			if cellNew == nil || cellNew.alt || cellNew.pac {
				continue
			}
			if g.heightAt(cNew) >= g.heightAt(c) {
				q = append(q, cNew)
				writeToG(cNew)
			}
		}
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	g := gridFromHeights(heights)
	fmt.Println(g.atlDrain)
	return heights
}
