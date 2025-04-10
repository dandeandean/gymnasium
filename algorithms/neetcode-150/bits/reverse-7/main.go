package main

import "fmt"

func reverse(x int) int {
	out := 0
	flipped := 1
	if x < 0 {
		x = -x
		flipped = -1
	}
	MAX := 2147483648
	if flipped == 1 {
		MAX -= 1
	}
	for x > 0 {
		lastDig := x % 10
		out *= 10
		out += lastDig
		if out > MAX {
			return 0
		}
		x = x / 10
	}
	return out * flipped
}

func main() {
	fmt.Println(
		reverse(123),
		reverse(-123),
		reverse(120),
	)
}
