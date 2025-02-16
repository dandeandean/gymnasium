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
	MAX := 9223372036854775807
	prices := make([]int, n)
	for i := range n {
		if i == src {
			continue
		}
		prices[i] = MAX
	}
	for range k + 1 {
		// IDK why we can't make pCopy outside of this loop
		pCopy := make([]int, n)
		copy(pCopy, prices)
		for _, f := range flights {
			src, dst, cost := f[0], f[1], f[2]
			if prices[src] != MAX &&
				prices[src]+cost < pCopy[dst] {
				pCopy[dst] = prices[src] + cost
			}
		}
		prices = pCopy
	}
	if prices[dst] == MAX {
		prices[dst] = -1
	}
	return prices[dst]
}
