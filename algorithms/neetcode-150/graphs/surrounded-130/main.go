package main

func solve(board [][]byte) {
	type coord [2]int
	x, o, beenTo := byte('X'), byte('O'), byte('Y')
	bfs := func(q []coord) {
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, d := range []coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				cNew := coord{cur[0] + d[0], cur[1] + d[1]}
				board[cur[0]][cur[1]] = beenTo
				if cNew[0] >= 0 &&
					cNew[1] >= 0 &&
					cNew[0] < len(board) &&
					cNew[1] < len(board[0]) &&
					board[cNew[0]][cNew[1]] == o {
					q = append(q, cNew)
				}
			}
		}
	}
	starting := make([]coord, 0)
	ROWS, COLS := len(board), len(board[0])
	for r := range ROWS {
		if board[r][0] == o {
			starting = append(starting, coord{r, 0})
		}
		if board[r][COLS-1] == o {
			starting = append(starting, coord{r, COLS - 1})
		}
	}
	for c := range COLS {
		if board[0][c] == o {
			starting = append(starting, coord{0, c})
		}
		if board[ROWS-1][c] == o {
			starting = append(starting, coord{ROWS - 1, c})
		}
	}
	bfs(starting)
	for r := range ROWS {
		for c := range COLS {
			if board[r][c] == o {
				board[r][c] = x
			}
			if board[r][c] == beenTo {
				board[r][c] = o
			}
		}
	}
}
