package main

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

func (this *WordDictionary) nodeSearcher(n *Node, subWord string) *Node {
	if n == nil || len(subWord) == 0 {
		return nil
	}
	if n.next[int(subWord[0]-'a')] != nil {
	}
}

func (this *WordDictionary) Search(word string) bool {
}

func main() {
	return
}
