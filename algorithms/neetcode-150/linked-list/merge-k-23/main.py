class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __repr__(self) -> str:
        outstr = f"({self.val}){'->' if self.next else ''}"
        cur = self.next
        while cur:
            outstr += f"({cur.val}){'->' if cur.next else ''}"
            cur = cur.next
        return outstr
    __str__ = __repr__

class Solution:
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


    def mergeKLists(self, lists: list[ListNode]) -> ListNode | None:
        if len(lists) <= 1:
            return lists[0] if len(lists) == 1 else None
        while len(lists) > 1:
            merged_so_far = []
            print("inside",lists)
            for i in range(0,len(lists), 2):
                print("\tmerged", merged_so_far)
                l1 = lists[i]
                l2 = lists[i+1] if (i + 1) < len(lists) else None
                merged_so_far.append(self.mergeList(l1,l2))
            lists = merged_so_far
        return lists[0]
            

if __name__ == "__main__":
    LN = ListNode 
    l1 = LN(1,LN(2,LN(3)))
    l2 = LN(4,LN(5,LN(6)))
    l3 = LN(5,LN(5,LN(6)))
    ls = [l1, l2, l3]
    print(ls)
    s = Solution()
    merged = s.mergeKLists(ls)
    print(merged)
    

