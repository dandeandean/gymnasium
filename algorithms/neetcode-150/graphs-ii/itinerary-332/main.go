package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(
		findItinerary(
			[][]string{{"HOU", "JFK"}, {"SEA", "JFK"}, {"JFK", "SEA"}, {"JFK", "HOU"}},
		),
		findItinerary(
			[][]string{{"BUF", "HOU"}, {"HOU", "SEA"}, {"JFK", "BUF"}},
		),
	)
}

func findItinerary(tickets [][]string) []string {
	aList := make(map[string][]string)
	for _, t := range tickets {
		aList[t[0]] = append(aList[t[0]], t[1])
	}
	for k := range aList {
		sort.Strings(aList[k])
	}
	res := []string{"JFK"}
	var dfs func(string) bool
	dfs = func(w string) bool {
		if len(res) == len(tickets)+1 {
			return true
		}
		if _, ok := aList[w]; !ok {
			return false
		}
		temp := make([]string, len(aList))
		copy(temp, aList[w])
		for i, str := range temp {
			aList[w] = append(aList[w][0:i], aList[w][i+1:]...)
			if dfs(str) {
				return true
			}
		}
	}
	dfs("JFK")

	return res
}
