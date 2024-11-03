package main

import (
	"fmt"
	//"slices"
)

type Subsets struct {
	Nums []int
	Res  [][]int
}

func (s *Subsets) dfs(i int, subset []int) {
	if i >= len(s.Nums) {
		s.Res = append(s.Res, subset)
		return
	}
	s.dfs(i+1, append(subset, s.Nums[i]))     // branch do append
	s.dfs(i+1, append([]int(nil), subset...)) // branch don't append
}

func subsets(nums []int) [][]int {
	s := Subsets{
		Nums: nums,
		Res:  make([][]int, 0),
	}
	s.dfs(0, make([]int, 0))
	return s.Res
}

func main() {
	// 0 3 5 9 twice
	wrong := subsets([]int{9, 0, 3, 5, 7})
	fmt.Println(len(wrong))
	for _, sl := range wrong {
		fmt.Println(sl)
	}

}
