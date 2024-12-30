package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}),
		findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}),
	)
}

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	rank := make([]int, len(edges)+1)
	var find func(int) int
	find = func(n int) int {
		if n != parent[n] {
			parent[n] = find(parent[n])
		}
		return parent[n]
	}
	union := func(a, b int) bool {
		pa, pb := find(a), find(b)
		if pa == pb {
			return false
		}

		larg, smal := pa, pb
		if rank[pa] < rank[pb] {
			larg, smal = pb, pa
		}

		parent[smal] = larg
		rank[larg] += rank[smal]
		return true
	}

	for i := range edges {
		parent[i] = i
		rank[i] = 1
	}
	for _, pair := range edges {
		if !union(pair[0], pair[1]) {
			return pair
		}
	}
	return []int{}
}
