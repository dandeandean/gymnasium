package main

func dfs(i int, nums []int, curRes [][]int) [][]int {
	curNums := make([]int, 0)
	if i >= len(curNums) {
		curRes = append(curRes, curNums)
		return curRes
	}
	dfs(i+1, nums, curRes)
	curNums = append(curNums, nums[i])
	dfs(i+1, nums, curRes)
	return curRes
}
func subsets(nums []int) [][]int {
	res := dfs(0, nums, make([][]int, 0))
	return res
}
