package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		ladderLength("cat", "sag", []string{"bat", "bag", "sag", "dag", "dot"}),
		ladderLength("cat", "sag", []string{"bat", "bag", "sat", "dag", "dot"}),
	)

}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	nei := make(map[string][]string)
	for _, word := range wordList {
		for i := range word {
			pattern := word[:i] + "*" + word[i+1:]
			nei[pattern] = append(nei[pattern], word)
		}
	}
	beenTo := make(map[string]bool)
	q := []string{beginWord}
	res := 1
	for len(q) != 0 {
		for range q {
			word := q[0]
			q = q[1:]
			if word == endWord {
				return res
			}
			for i := range word {
				pattern := word[:i] + "*" + word[i+1:]
				for _, n := range nei[pattern] {
					if !beenTo[n] {
						beenTo[n] = true
						q = append(q, n)
					}
				}
			}
		}
		res++
	}
	return 0
}
