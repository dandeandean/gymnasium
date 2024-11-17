package main

import (
	"fmt"
)

type Board struct {
	b     []string
	col   map[int]bool
	row   map[int]bool
	diagL map[int]bool
	diagR map[int]bool
}

func createBoard(n int) *Board {
	b := make([]string, n)
	s := ""
	for i := 0; i < n; i++ {
		s += "."
	}
	for i := 0; i < n; i++ {
		b[i] = s
	}
	return &Board{
		b:     b,
		col:   make(map[int]bool),
		row:   make(map[int]bool),
		diagL: make(map[int]bool),
		diagR: make(map[int]bool),
	}
}

func (b Board) rowClear(r, _ int) bool   { return !b.row[r] }
func (b Board) colClear(_, c int) bool   { return !b.col[c] }
func (b Board) diagLClear(r, c int) bool { return !b.diagL[r+c] }
func (b Board) diagRClear(r, c int) bool { return !b.diagR[r-c] }
func (b Board) canPlaceQ(r, c int) bool {
	return b.rowClear(r, c) && b.colClear(r, c) && b.diagRClear(r, c) && b.diagLClear(r, c)
}

func (b *Board) tryPlaceQ(r, c int) *Board {
	if !b.canPlaceQ(r, c) {
		return nil
	}
	bArr := []byte((*b).b[r])
	bArr[c] = 'Q'
	(*b).b[r] = string(bArr)
	(*b).row[r] = true
	(*b).col[c] = true
	(*b).diagL[r+c] = true
	(*b).diagR[r-c] = true
	return b
}

func (b Board) copy() *Board {
	res := createBoard(len(b.b[0]))
	copy(res.b, b.b)
	for k, v := range b.col {
		res.col[k] = v
	}
	for k, v := range b.row {
		res.row[k] = v
	}
	for k, v := range b.diagL {
		res.diagL[k] = v
	}
	for k, v := range b.diagR {
		res.diagR[k] = v
	}
	return res
}

func solveNQueens(n int) [][]string {
	board := *createBoard(n)
	res := make([][]string, 0)
	var backtrack func(int, Board)
	backtrack = func(r int, b Board) {
		if r >= n {
			res = append(res, b.copy().b)
			return
		}
		for c := range b.b[r] {
			cp := b.copy().tryPlaceQ(r, c)
			if cp != nil {
				backtrack(r+1, *cp)
			}
		}
	}
	backtrack(0, board)
	return res
}

func main() {
	fmt.Println(solveNQueens(4))
}
