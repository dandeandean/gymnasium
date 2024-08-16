class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def mergeKLists(self, lists: list[ListNode]) -> ListNode | None:
        if len(lists) <= 1:
            return lists[0] if len(lists) == 1 else None

    def mergeList(self, l1:ListNode | None, l2: ListNode | None) -> ListNode | None:
        anchor: ListNode = ListNode()
        cur: ListNode | None = anchor
        while l1 and l2:
            if l1.val < l2.val:
                cur.next = l1
                l1 = l1.next
            else:
                cur.next = l2
                l2 = l2.next
            cur = cur.next
        if l1:
            cur.next = l1
        if l2:
            cur.next = l2
        return anchor.next



