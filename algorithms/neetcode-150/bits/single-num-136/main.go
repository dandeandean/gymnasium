package main

func singleNumber(nums []int) int {
	//bits XOR with self => 0 ...
	//X xor Y -> (X || Y) && !(X && Y) (according to stack overflow)
	res := 0
	for _, n := range nums {
		res = (res | n) & ^(res & n)
	}
	return res
}
