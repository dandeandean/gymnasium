# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        cur1 = l1
        cur2 = l2
        anchor: ListNode = ListNode()
        cur_out = anchor
        carry:int = 0
        while cur1 or cur2 or carry:
            val1 = 0
            val2 = 0
            if cur1:
                val1 = cur1.val
            if cur2:
                val2 = cur2.val
            p_sum: int = 0
            p_sum = val1 + val2 + carry
            carry = p_sum // 10
            cur_out.next = ListNode(p_sum%10)
            cur_out = cur_out.next
            cur1 = cur1.next if cur1 else None
            cur2 = cur2.next if cur2 else None
        return anchor.next

if __name__ == "__main__":
    l1a = ListNode(0)
    l1b = ListNode(1,l1a)
    l1 = ListNode(3,l1b)

    l2c = ListNode(3)
    l2b = ListNode(1,l2c)
    l2 = ListNode(1,l2b)
    s = Solution()
    s.addTwoNumbers(l1,l2)
