package main

import (
	"fmt"
	"sort"
)

type state struct {
	res        [][]int
	candidates []int
	target     int
}

func (s *state) backtrack(i int, cur []int, tot int) {
	if tot == s.target {
		s.res = append(s.res, cur)
		return
	}
	if i >= len(s.candidates) || tot > s.target {
		return
	}
	s.backtrack(i+1, append(cur, s.candidates[i]), tot+s.candidates[i])
	for i+1 < len(s.candidates) && s.candidates[i] == s.candidates[i+1] {
		i++
	}
	s.backtrack(i+1, append([]int(nil), cur...), tot)
}

func combinationSum2(candidates []int, target int) [][]int {
	candCopy := append([]int{}, candidates...)
	sort.Ints(candCopy)
	fmt.Println(candCopy)
	s := state{
		res:        make([][]int, 0),
		candidates: candCopy,
		target:     target,
	}
	s.backtrack(0, make([]int, 0), 0)
	return s.res
}

func main() {
	fmt.Println(
		combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8),
	)
}
