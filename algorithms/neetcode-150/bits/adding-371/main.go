package main

import "fmt"

func getSum(a int, b int) int {
	mask := 0xFFFFFFFF
	for b != 0 {
		carry := (a & b) << 1
		a = (a ^ b) & mask
		b = mask & carry
	}
	if a > 2000 || a < -2000 {
		a = ^(a ^ mask)
	}
	return a
}

func main() {
	for i := range 100 {
		for j := range 100 {
			if getSum(i, j) != (i + j) {
				fmt.Println(getSum(i, j), i+j)
			}
		}
	}
}
