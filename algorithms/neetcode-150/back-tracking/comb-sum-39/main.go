package main

import "fmt"

type comb struct {
	i      int   // candidates[i:] that we can use
	cur    []int // cur subset we may add to res
	sum    int   // sum of cur
	res    [][]int
	cand   []int
	target int
}

func makeComb(cand []int, target int) comb {
	return comb{
		i:      0,
		cur:    make([]int, 0),
		sum:    0,
		res:    make([][]int, 0),
		cand:   cand,
		target: target,
	}
}

func (c *comb) append(v int) {
	c.cur = append(c.cur, v)
	c.sum = c.sum + v
}

func (c comb) dump() []int {
	return append([]int(nil), c.cur...)
}

func (c comb) take() comb {
	// copys slices, not cand
	cNew := c
	cNew.res = append([][]int(nil), cNew.res...)
	cNew.cur = append([]int(nil), cNew.cur...)
	if &cNew.res == &c.res || &cNew.cur == &c.cur {
		fmt.Println("something horrible happened")
	}
	return cNew
}

func dfs(c comb) {
	if c.sum == c.target {
		c.res = append(c.res, c.dump())
		return
	}
	if c.i >= len(c.cand) || c.sum > c.target {
		return
	}
	cLeft := c.take()
	cLeft.append(c.cand[c.i])
	dfs(cLeft)

	cRight := c.take()
	cRight.i += 1
	dfs(cRight)
}

func combinationSum(candidates []int, target int) [][]int {
	c := makeComb(candidates, target)
	dfs(c)
	return c.res
}
func main() {
	fmt.Println(
		combinationSum([]int{2, 3, 6, 7}, 7),
	)
}
