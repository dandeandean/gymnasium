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
	stack := []string{"JFK"}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		res = append(res, cur)
	}
	return res
}
