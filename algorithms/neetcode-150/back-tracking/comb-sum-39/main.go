package main

import "fmt"

type comb struct {
	res    [][]int
	cand   []int
	target int
}

func (c *comb) dfs(i int, cur []int, tot int) {
	if tot == c.target {
		fmt.Println(tot, c.res, cur)
		c.res = append(c.res, cur)
		fmt.Println(c.res)
		return
	}
	if i >= len(c.cand) || tot > c.target {
		return
	}
	c.dfs(i, append(cur, c.cand[i]), tot+c.cand[i])
	c.dfs(i+1, append([]int(nil), cur...), tot)
}

func combinationSum(candidates []int, target int) [][]int {
	c := comb{
		res:    make([][]int, 0),
		cand:   candidates,
		target: target,
	}
	c.dfs(0, []int{}, 0)
	return c.res
}

func main() {
	fmt.Println(
		//combinationSum([]int{2, 3, 6, 7}, 7),
		combinationSum([]int{2, 3, 5}, 8),
	)
}
