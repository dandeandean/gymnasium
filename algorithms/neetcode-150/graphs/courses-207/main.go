package main

import "fmt"

func main() {
	fmt.Println(
		canFinish(2, [][]int{{0, 1}}),
		canFinish(2, [][]int{{1, 0}}),
	)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)
	for _, p := range prerequisites {
		preMap[p[0]] = append(preMap[p[0]], p[1])
	}
	beenTo := make(map[int]bool)
	var dfs func(int) bool
	dfs = func(c int) bool {
		ps, _ := preMap[c]
		if beenTo[c] {
			return false
		}
		if len(ps) == 0 {
			return true
		}
		beenTo[c] = true
		for _, pp := range preMap[c] {
			if !dfs(pp) {
				return false
			}
		}
		beenTo[c] = false
		preMap[c] = make([]int, 0)
		return true
	}

	for c := range numCourses {
		if !dfs(c) {
			return false
		}
	}
	return true
}
