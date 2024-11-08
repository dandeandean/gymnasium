package main

type state struct {
	res         [][]int
	nums        []int
	usedChoices map[int]bool
}

func (s *state) curse(i int, subset []int) {
	if i > len(subset) {
		s.res = append(s.res, subset)
	}
	sCopy := append([]int{}, subset...)
	s.curse(i+1, sCopy)
	if s.usedChoices[i] == false {
		s.usedChoices[i] = true
		s.curse(i, append(sCopy, s.nums[i]))
	}
}
func subsetsWithDup(nums []int) [][]int {
	s := state{
		res:         make([][]int, 0),
		nums:        nums,
		usedChoices: make(map[int]bool),
	}
	s.curse(0, []int{})
	return s.res
}
