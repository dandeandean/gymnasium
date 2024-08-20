package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	anchor := &ListNode{Next: head}
	cur := anchor.Next
	// not O(1) memory
	var m map[int]*ListNode = make(map[int]*ListNode)
	i := 0
	for cur != nil {
		fmt.Println(cur)
		m[i] = cur
		i++
		cur = cur.Next
	}
	groupPrev := anchor
	for {
		kth, ok := m[k]
		if !ok {
			break
		}
		groupNext := kth.Next
		prev, cur := kth.Next, groupPrev.Next
		for cur != groupNext {
			tmp := cur.Next
			cur.Next = prev
			prev = cur
			cur = tmp
		}
		tmp := groupPrev.Next
		groupPrev.Next = prev
		prev = groupPrev
		groupPrev = tmp
	}

	return anchor.Next
}

func main() {
	reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 2)
}
