package main

func hammingWeight(n int) int {
	//this is just even or odd foreach bit, then bitshift
	res := 0
	for n > 0 {
		res += n % 2
		n >>= 1
	}
	return res
}
