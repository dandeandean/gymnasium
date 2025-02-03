package main

import "fmt"

func main() {
	test1 := []string{"hrn", "hrf", "er", "enn", "rfnn"}
	test2 := []string{"z", "o"}
	fmt.Println(
		AlienDictinary(test1),
		AlienDictinary(test2),
	)
}

func AlienDictinary(words []string) string {
	// return their lexographical ordering
	// to reprsent an adj list
	comesB4 := make(map[byte]map[byte]bool)
	for _, w := range words {
		for _, r := range w {
			comesB4[byte(r)] = make(map[byte]bool)
		}
	}

	for i := range len(words) - 1 {
		w1, w2 := words[i], words[i+1]
		smallerLen := len(w1)
		if len(w2) < smallerLen {
			smallerLen = len(w2)
		}
		if len(w1) > len(w2) && w1[:smallerLen] == w2[:smallerLen] {
			return ""
		}
		for j := range smallerLen {
			if w1[j] != w2[j] {
				comesB4[w1[j]][w2[j]] = true
				break
			}
		}
	}

	visited := make(map[byte]bool)
	var dfs func(c byte) bool
	dfs = func(c byte) bool {
		if visited[c] {
			return visited[c]
		}
		visited[c] = true
		return true
	}
	dfs(words[0][0])
	return words[0]
}
