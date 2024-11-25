package main

type Node struct {
	c     rune
	isEnd bool
	next  [26]*Node
}

type Trie struct {
	root *Node
}

func newNode(c rune, isEnd bool) *Node {
	return &Node{
		c:     c,
		isEnd: isEnd,
		next:  [26]*Node(make([]*Node, 26)),
	}
}

func Constructor() Trie {
	return Trie{root: newNode('~', false)}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, c := range word {
		i := int(c - rune('a'))
		if cur.next[i] == nil {
			cur.next[i] = newNode(c, false)
		}
		cur = cur.next[i]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, c := range word {
		i := int(c - rune('a'))
		if cur.next[i] == nil {
			return false
		}
		cur = cur.next[i]
	}
	return cur.isEnd
}

func findWords(board [][]byte, words []string) []string {
	// do word search i but with a trie
	trie := Trie{root: newNode('~', false)}
	ROWS, COLS := len(board), len(board[0])
	for _, word := range words {
		trie.Insert(word)
	}
	res := make([]string, 0)
	var dfs func(r int, c int, n *Node, word string) bool
	dfs = func(r int, c int, n *Node, word string) bool {
		return true
	}
	for r := 0; r < ROWS; r++ {
		for c := 0; r < COLS; c++ {
			dfs(r, c, trie.root, "")
		}
	}

}
