package main

import "fmt"

type Node struct {
	c     *rune
	isEnd bool
	next  [26]*Node
}

func newNode(c *rune) *Node {
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
		root: newNode(nil),
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, c := range word {
		i := int(c - rune('a'))
		if cur.next[i] == nil {
			cur.next[i] = newNode(&c)
		}
		cur = cur.next[i]
	}
	cur.isEnd = true
}

func (this *WordDictionary) nodeSearcher(n *Node, subWord string) bool {
	if subWord == "" {
		return true
	}
	if n == nil {
		return false
	}
	cur := n
	for _, c := range subWord {
		if c == '.' {
			for i := 0; i < 26; i++ {
				if this.nodeSearcher(n.next[i], subWord[1:]) {
					return true
				}
			}
			return false
		} else {
			i := int(c - rune('a'))
			if cur.next[i] == nil {
				return false
			}
			cur = cur.next[i]
		}
	}
	return cur.isEnd
}

func (this *WordDictionary) Search(word string) bool {
	return this.nodeSearcher(this.root, word)
}

func main() {
	wd := Constructor()
	wd.AddWord("a")
	fmt.Println(
		wd.Search("aa"),
		wd.Search(".a"),
		wd.Search("a."),
	)
}
