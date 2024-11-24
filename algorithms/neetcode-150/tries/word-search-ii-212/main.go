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
	for _, word := range words {
		trie.Insert(word)
	}
	return []string{"foobar"}
}
