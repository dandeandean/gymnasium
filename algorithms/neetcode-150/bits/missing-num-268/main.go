package main

func missingNumber(nums []int) int {
	/*
		This took me a little while to grok
			consider we weren't missing one, then we could xor
			everything with index i := range [0,n] and we would get zero b/c
			we know everything in nums := range [0, n]. This would result in
			res = 0.

		Remember that 0 ^ N = N.
			Since we are missing one, this res will actually be what's left out
			of the list.
	*/
	res := len(nums)
	for i, n := range nums {
		res ^= (i ^ n)
	}
	return res
}
