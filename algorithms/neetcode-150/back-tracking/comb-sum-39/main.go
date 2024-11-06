package main

import "fmt"

type comb struct {
	res    [][]int
	cand   []int
	target int
}

func makeComb(cand []int, target int) comb {
	return comb{
		res:    make([][]int, 0),
		cand:   cand,
		target: target,
	}
}

func (c *comb) dfs(i int, cur []int, tot int) {
	if tot == c.target {
		c.res = append(c.res, cur)
		return
	}
	if i >= len(c.cand) || tot > c.target {
		return
	}
	c.dfs(i, append(cur, c.cand[i]), tot+c.cand[i])
	c.dfs(i+1, cur, tot)
}

func combinationSum(candidates []int, target int) [][]int {
	c := makeComb(candidates, target)
	c.dfs(0, []int{}, 0)
	return c.res
}
func main() {
	fmt.Println(
		combinationSum([]int{2, 3, 6, 7}, 7),
	)
}
