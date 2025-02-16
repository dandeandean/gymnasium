package main

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for range 31 {
		res |= (num % 2)
		num >>= 1
		res <<= 1
	}
	return res | (num % 2)
}
