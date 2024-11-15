package main

import (
	"fmt"
)

type Board []string

func createBoard(n int) *Board {
	res := make(Board, n)
	s := ""
	for i := 0; i < n; i++ {
		s += "."
	}
	for i := 0; i < n; i++ {
		res[i] = s
	}
	return &res
}

func (b Board) rowClear(r, _ int) bool {
	for _, ch := range b[r] {
		if ch != '.' {
			return false
		}
	}
	return true
}

func (b Board) colClear(_, c int) bool {
	for r := range b {
		if b[r][c] != '.' {
			return false
		}
	}
	return true
}

func (b Board) diagLClear(r, c int) bool {
	for r > 0 && c < len(b[0])-1 { // cast up to the right
		r--
		c++
	}
	for r < len(b)-1 && c > 0 {
		if b[r][c] != '.' {
			return false
		}
		r++
		c--
	}
	return true
}

func (b Board) diagRClear(r, c int) bool {
	for r < len(b[0])-1 && c > 0 { // cast up to the right
		r++
		c--
	}
	for r > 0 && c < len(b[0])-1 {
		if b[r][c] != '.' {
			return false
		}
		r--
		c++
	}
	return true
}

func (b Board) canPlaceQ(r, c int) bool {
	return b.rowClear(r, c) && b.colClear(r, c) &&
		b.diagRClear(r, c) && b.diagLClear(r, c)
}

func (b *Board) tryPlaceQ(r, c int) bool {
	if !b.canPlaceQ(r, c) {
		return false
	}
	byts := []byte((*b)[r])
	byts[c] = 'Q'
	(*b)[r] = string(byts)
	return true
}

/*
func diagRClear(r, c int) bool
func canPlaceQ(r, c int) bool
func backtrack(n int)
func solveNQueens(n int) [][]string
*/
func main() {
	b := createBoard(4)
	for i := range 4 {
		for j := range 4 {
			b.tryPlaceQ(i, j)
		}
		fmt.Println(b)
	}
}
