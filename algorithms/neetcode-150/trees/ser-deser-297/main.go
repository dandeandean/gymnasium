package main

import (
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
	letters string
	i       int
	vals    []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) dfs(node *TreeNode) {
	if node == nil {
		this.letters += ",N"
		return
	}
	this.letters += "," + string(node.Val)
	this.dfs(node.Left)
	this.dfs(node.Right)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.dfs(root)
	// this is a bit dumb
	return this.letters
}

// Deserializes your encoded data to tree.
func (this *Codec) ddfs() *TreeNode {
	nodeVal, err := strconv.Atoi(this.vals[this.i])
	this.i++
	if err != nil {
		return nil
	}
	node := &TreeNode{Val: nodeVal}
	node.Left = this.ddfs()
	node.Right = this.ddfs()
	return node
}
func (this *Codec) deserialize(data string) *TreeNode {
	this.i = 0
	this.vals = strings.Split(data, ",")
	return this.ddfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
