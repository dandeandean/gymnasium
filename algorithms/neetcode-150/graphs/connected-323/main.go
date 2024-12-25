package main

import "fmt"

func main() {
	fmt.Println(
		countComponents(3, [][]int{{0, 1}, {0, 2}}),
		countComponents(6, [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 5}}),
	)
}

func countComponents(n int, edges [][]int) int {
	adjList := make(map[int][]int)
	beenTo := make(map[int]bool)
	res := 0
	for _, e := range edges {
		from, to := e[0], e[1]
		adjList[from] = append(adjList[from], to)
		adjList[to] = append(adjList[to], from)
	}

	var dfs func(int)
	dfs = func(i int) {
		for _, n := range adjList[i] {
			if !beenTo[n] {
				beenTo[n] = true
				dfs(n)
			}
		}
	}
	for i := range n {
		if !beenTo[i] {
			beenTo[i] = true
			dfs(i)
			res++
		}
	}
	return res
}
