package main

import "fmt"

type coord struct {
	x int
	y int
	c byte
}
type Board struct {
	data [][]byte
	rows int
	cols int
	word string
	been map[coord]bool // this map may cause weird behavior
}

func boardFrom(board [][]byte, word string) Board {
	return Board{
		data: board,
		rows: len(board),
		cols: len(board[0]),
		word: word,
		been: make(map[coord]bool),
	}
}
func (b Board) inBounds(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= b.rows || y >= b.cols {
		return false
	}
	return true
}
func (b Board) getCoord(x, y int) coord {
	if !b.inBounds(x, y) {
		panic("called getCoord out of bounds!!")
	}
	c := coord{x: x, y: y, c: b.data[x][y]}
	return c
}
func (b Board) beenTo(x, y int) bool {
	if !b.inBounds(x, y) {
		return false
	}
	c := b.getCoord(x, y)
	return b.been[c]
}
func (b *Board) visit(x, y int) {
	if !b.inBounds(x, y) {
		return
	}
	b.been[b.getCoord(x, y)] = true
}
func (b *Board) unvisit(x, y int) {
	if !b.inBounds(x, y) {
		return
	}
	b.been[b.getCoord(x, y)] = false
}

func (b *Board) recurse(x, y, i int) bool {
	if i == len(b.word) {
		return true
	}
	if !b.inBounds(x, y) || b.beenTo(x, y) {
		return false
	}
	if b.word[i] == b.getCoord(x, y).c {
		return false
	}
	b.visit(x, y)
	i++
	res := b.recurse(x+1, y, i) || b.recurse(x-1, y, i) ||
		b.recurse(x, y+1, i) || b.recurse(x, y-1, i)
	b.unvisit(x, y)
	return res
}

func exist(board [][]byte, word string) bool {
	b := boardFrom(board, word)
	for i := range b.data {
		for j := range b.data[i] {
			if b.recurse(i, j, 0) {
				return true
			}
		}
	}
	return false
}
func main() {
	b := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//fmt.Println(exist(b, "ABCCED"))
	fmt.Println(exist(b, "ABCB"))
}
