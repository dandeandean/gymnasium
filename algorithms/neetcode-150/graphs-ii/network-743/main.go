package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

type node []int
type heapPair struct {
	node *node
	dist int
}

func networkDelayTime(times [][]int, n int, k int) int {
	edges := make(map[int][]node, 0)
	for _, t := range times {
		edges[t[0]] = append(edges[t[0]], node{t[1], t[2]})
	}
	q := []int{k}
	beenTo := make([]bool, n+1)
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for _, nei := range edges[cur] {
			if !beenTo[cur] {
				beenTo[cur] = true
				q = append(q, nei[0])
			}
		}

	}
	return 69
}
