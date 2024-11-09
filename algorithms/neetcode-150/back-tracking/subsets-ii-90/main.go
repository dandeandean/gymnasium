package main

import (
	"fmt"
	"sort"
)

type state struct {
	res  [][]int
	nums []int
}

func (s *state) curse(i int, subset []int) {
	if i >= len(s.nums) {
		s.res = append(s.res, subset)
		return
	}
	s.curse(i+1, append(subset, s.nums[i]))
	for i+1 < len(s.nums) && s.nums[i] == s.nums[i+1] {
		i++
	}
	s.curse(i+1, append([]int(nil), subset...))
}

func subsetsWithDup(nums []int) [][]int {
	numsSorted := append([]int(nil), nums...)
	sort.Ints(numsSorted)
	s := state{
		res:  make([][]int, 0),
		nums: numsSorted,
	}
	s.curse(0, make([]int, 0))
	return s.res
}

func main() {
	fmt.Println(
		"subsets ii: ",
		subsetsWithDup([]int{2, 1, 2}),
	)
}
