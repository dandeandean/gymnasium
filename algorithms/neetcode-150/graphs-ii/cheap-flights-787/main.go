package main

import "fmt"

func main() {

	case1 := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	case2 := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}

	fmt.Println(
		findCheapestPrice(4, case1, 0, 3, 1),
		findCheapestPrice(3, case2, 0, 2, 1),
		findCheapestPrice(3, case2, 0, 2, 0),
	)
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	res := n + src + dst + k + len(flights)
	return res
}
