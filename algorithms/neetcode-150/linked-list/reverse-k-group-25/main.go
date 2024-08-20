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
		m[i] = cur
		i++
		cur = cur.Next
	}
	groupPrev := anchor
	i = 0
	for {
		kth, ok := m[k+i]
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
		i += k
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
	oldList := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	printList(oldList)
	newList := reverseKGroup(oldList, 2)
	printList(newList)
}
