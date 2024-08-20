package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func jumpK(start *ListNode, k int) *ListNode {
	cur := start
	for k > 0 && cur != nil {
		if k >= 2 && cur.Next != nil {
			cur = cur.Next
			k--
		}
		cur = cur.Next
		k--
	}
	return cur
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	anchor := &ListNode{Next: head}
	groupPrev := anchor
	for {
		kth := jumpK(groupPrev, k)
		if nil == kth {
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

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("(%d)", cur.Val)
		cur = cur.Next
		if cur != nil {
			fmt.Printf("->")
		} else {
			fmt.Printf("\n")
		}
	}
}
func main() {
	oldList := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printList(oldList)
	newList := reverseKGroup(oldList, 2)
	printList(newList)
}
