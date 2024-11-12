package main

import "fmt"

func palOrBust(s string, i, j int) string {
	// inclusive
	if i > j {
		return ""
	}
	if j >= len(s) {
		return ""
	}
	if i == j {
		return string(s[i])
	}
	j += 1
	x, y := i, j
	for i < j {
		if s[i] != s[j-1] {
			return ""
		}
		i, j = i+1, j-1
	}
	return s[x:y]
}
func partition(s string) [][]string {
	res := make([][]string, 0)
	part := make([]string, 0)
	var dfs func(int)
	dfs = func(i int) {
		if i >= len(s) {
			res = append(res, append([]string{}, part...))
			return
		}
		for j := i; j <= len(s); j++ {
			p := palOrBust(s, i, j)
			if len(p) > 0 {
				part = append(part, p)
				dfs(j + 1)
				part = part[0 : len(part)-1]
			}
		}
	}
	dfs(0)
	return res
}
func main() {
	fmt.Println(
		partition("aab"),
		partition("a"),
		partition("bb"),
	)
}
