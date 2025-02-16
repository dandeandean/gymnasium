package main

func hammingWeight(n int) int {
	//from 191
	res := 0
	for n > 0 {
		res += n % 2
		n >>= 1
	}
	return res
}
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := range n + 1 {
		res[i] = hammingWeight(i)
	}
	return res
}
