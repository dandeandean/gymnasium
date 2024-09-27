package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	Letters string
	Vals    []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) dfs(node *TreeNode) {
	if node == nil {
		this.Letters += ",N"
		return
	}
	letter := strconv.Itoa(node.Val)
	if len(this.Letters) == 0 {
		this.Letters += letter
	} else {
		this.Letters += "," + letter
	}
	this.dfs(node.Left)
	this.dfs(node.Right)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.dfs(root)
	// this is a bit dumb
	return this.Letters
}

// Deserializes your encoded data to tree.
func (this *Codec) ddfs() *TreeNode {
	nodeVal, err := strconv.Atoi(this.Vals[0])
	this.Vals = this.Vals[1:]
	if err != nil { // this can't be working bc atoi proll errs on N
		return nil
	}
	node := &TreeNode{Val: nodeVal}
	node.Left = this.ddfs()
	node.Right = this.ddfs()
	return node
}
func (this *Codec) deserialize(data string) *TreeNode {
	this.Vals = strings.Split(data, ",")
	return this.ddfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
func main() {
	s := Constructor()
	d := Constructor()
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	serial := s.serialize(root)
	fmt.Println(serial)
	fmt.Println(strings.Split(serial, ","))
	rootNew := d.deserialize(serial)
	fmt.Println(
		rootNew.Val,
		rootNew.Left,
		root.Right,
	)
}
