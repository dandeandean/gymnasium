package main

import "fmt"

func main() {
	fmt.Println(
		validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}),
		validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}),
	)
}

func validTree(n int, edges [][]int) bool {
	if n == 0 {
		return true
	}
	adjList := make(map[int][]int)
	for _, e := range edges {
		adjList[e[0]] = append(adjList[e[0]], e[1])
		adjList[e[1]] = append(adjList[e[1]], e[0])
	}
	var dfs func(i int, prev int) bool
	visit := make(map[int]bool)
	dfs = func(i, prev int) bool {
		if visit[i] {
			return false
		}
		visit[i] = true
		for _, nextI := range adjList[i] {
			if nextI != prev && !dfs(nextI, i) {
				return false
			}
		}
		visit[i] = false
		return true
	}
	return dfs(0, -1) && len(visit) == n
}
