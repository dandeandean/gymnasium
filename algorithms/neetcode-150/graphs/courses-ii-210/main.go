package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	preMap := make(map[int][]int)
	for _, p := range prerequisites {
		preMap[p[0]] = append(preMap[p[0]], p[1])
	}
	beenTo, inCycle, res := make(map[int]bool), make(map[int]bool), make([]int, 0)

	var dfs func(int) bool
	dfs = func(i int) bool {
		if inCycle[i] {
			return false
		}
		if beenTo[i] {
			return true
		}
		inCycle[i] = true
		for _, p := range preMap[i] {
			if !dfs(p) {
				return false
			}
		}
		inCycle[i] = false
		beenTo[i] = true
		res = append(res, i)
		return true
	}

	for c := range numCourses {
		if !dfs(c) {
			return []int{}
		}
	}
	return res
}
