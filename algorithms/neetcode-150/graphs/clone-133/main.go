package main

// given
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	var dfs func(*Node) *Node
	alreadyDone := make(map[*Node]*Node)
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		c, ok := alreadyDone[n]
		if ok {
			return c
		}
		strange := &Node{Val: n.Val}
		alreadyDone[n] = strange
		for _, m := range n.Neighbors {
			strange.Neighbors = append(strange.Neighbors, dfs(m))
		}
		return strange
	}
	return dfs(node)
}
