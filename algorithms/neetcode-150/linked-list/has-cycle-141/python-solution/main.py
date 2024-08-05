# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head: ListNode | None) -> bool: 
        been_to = set()
        cur = head
        while cur : 
            if cur.next in been_to:
                return True
            been_to.add(cur)
            cur = cur.next
        return False

    def hasCycle0(self, head: ListNode) -> bool: 
        # Slow & Fast pointers
        slow: ListNode | None = head
        fast: ListNode | None = head.next
        while slow and fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow == fast: return True
        return False
