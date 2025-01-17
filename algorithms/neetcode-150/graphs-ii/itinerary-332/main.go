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
		sort.Sort(sort.Reverse(sort.StringSlice(aList[k])))
	}
	res := make([]string, 0)
	var dfs func(string)
	dfs = func(s string) {
		for len(aList[s]) > 0 {
			last := aList[s][len(aList[s])-1]
			aList[s] = aList[s][:len(aList[s])-1]
			dfs(last)
		}
		res = append(res, s)
	}
	dfs("JFK")
	for i, s := range res {
		temp := s
		res[i] = res[len(res)-1-i]
		res[len(res)-1-i] = temp
		if i == (len(res)-1)/2 {
			break
		}
	}
	return res
}
