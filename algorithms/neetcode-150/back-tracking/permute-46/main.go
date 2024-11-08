package main

import "fmt"

type ans struct {
	res [][]int
}

func popIth(nums []int, i int) []int {
	return append(nums[:i], nums[i+1:]...)
}

func (a *ans) recurse(soFar []int, toGo []int) {
	if len(toGo) == 0 {
		soFarCopy := append([]int(nil), soFar...)
		a.res = append(a.res, soFarCopy)
	}
	for i := 0; i < len(toGo); i++ {
		toGoCopy := append([]int(nil), toGo...)
		soFarCopy := append(soFar, toGoCopy[i])
		toGoCopy = popIth(toGoCopy, i)
		a.recurse(soFarCopy, toGoCopy)
	}
}

func permute(nums []int) [][]int {
	A := ans{make([][]int, 0)}
	A.recurse([]int{}, nums)
	return A.res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{1, 2}))
}
