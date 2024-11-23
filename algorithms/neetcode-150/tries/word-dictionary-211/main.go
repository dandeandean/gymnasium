package main

import "fmt"

func main() {
	wd := Constructor()
	wd.AddWord("a")
	fmt.Println(
		wd.Search("a."),
	)
}

type Node struct {
	c     rune
	isEnd bool
	next  [26]*Node
}

func newNode(c rune) *Node {
	return &Node{
		c:     c,
		isEnd: false,
		next:  [26]*Node(make([]*Node, 26)),
	}
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: newNode(rune('~')),
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, c := range word {
		i := int(c - rune('a'))
		if cur.next[i] == nil {
			cur.next[i] = newNode(c)
		}
		cur = cur.next[i]
	}
	cur.isEnd = true
}

func (n Node) getChildren() []*Node {
	out := make([]*Node, 0, 26)
	for i := 0; i < 26; i++ {
		if n.next[i] != nil {
			out = append(out, n.next[i])
		}
	}
	return out
}

func (this *WordDictionary) nodeSearcher(n *Node, subWord string) bool {
	if n == nil {
		return false
	}
	if len(subWord) == 0 {
		if n.isEnd {
			return true
		}
		return false
	}
	i := int(rune(subWord[0]) - rune('a'))
	if i > 25 || i < 0 {
		for _, ch := range n.getChildren() {
			if this.nodeSearcher(ch, subWord[1:]) {
				return true
			}
		}
		return false
	}
	return this.nodeSearcher(n.next[i], subWord[1:])
}

func (this *WordDictionary) Search(word string) bool {
	return this.nodeSearcher(this.root, word)
}

func printNode(n Node) {
	fmt.Print(
		"NODE: (",
		string(n.c),
		"->",
	)
	fmt.Print("[")
	for _, ch := range n.getChildren() {
		fmt.Print("", string(ch.c))
	}
	fmt.Println("])")
}
