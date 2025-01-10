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
	fmt.Println(aList)
	return []string{}
}
