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
		fmt.Println(s.nums, subset)
		s.res = append(s.res, subset)
		return
	}
	sCopy := append([]int(nil), subset...)
	s.curse(i+1, append(sCopy, s.nums[i]))
	for i+1 < len(s.nums) && s.nums[i] == s.nums[i+1] {
		i++
	}
	s.curse(i+1, sCopy)
}
func subsetsWithDup(nums []int) [][]int {
	s := state{
		res:  make([][]int, 0),
		nums: sort.IntSlice(nums),
	}
	s.curse(0, make([]int, 0))
	return s.res
}

func main() {
	fmt.Println(
		"subsets ii: ",
		subsetsWithDup([]int{1, 2, 2}),
	)
}
