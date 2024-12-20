package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	preMap := make(map[int][]int)
	for _, p := range prerequisites {
		preMap[p[0]] = append(preMap[p[0]], p[1])
	}
	beenTo, cycle := make(map[int]bool), make(map[int]bool)
	var dfs func(int) bool
	dfs = func(i int) bool {
		prereqs := preMap[i]
		if beenTo[i] || cycle[i] {
			return false
		}
		beenTo[i] = true
		for _, p := range prereqs {
			if !dfs(p) {
				return false
			}
		}
		beenTo[i] = false
	}

	res := make([]int, 0)
	return res
}
