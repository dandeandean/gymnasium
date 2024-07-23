// Definition for singly-linked list.
#include <cstddef>
#include <deque>
#include <iostream>
#include <vector>
struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};
class Solution {
public:
  void reorderList(ListNode *head) {
    ListNode *end, *mid;
    mid = head;
    end = head->next;
    while (end && end->next) {
      mid = end->next;
      end = end->next->next;
    }
    Solution::mid_end(head, mid, end);
    std::cout << "m=" << mid->val << "\ne=" << end->val << std::endl;
  }

private:
  void mergeList(ListNode *head1, ListNode *head2) {}
  void reverseList(ListNode *head) {
    ListNode *cur = head;
    ListNode *prev;
    while (cur) {
      ListNode *old_next = cur->next;
      cur->next = prev;
      prev = cur;
      cur = old_next;
    }
  }
};

int main(void) {
  ListNode *f = new ListNode(6);
  ListNode *e = new ListNode(5, f);
  ListNode *d = new ListNode(4, e);
  ListNode *c = new ListNode(3, d);
  ListNode *b = new ListNode(2, c);
  ListNode *a = new ListNode(1, b);
  Solution *s = new Solution;
  ListNode *cur = a;
  s->reorderList(a);
  cur = a;
  while (cur) {
    std::cout << "(" << cur->val << ")->&" << cur->next << "\n";
    cur = cur->next;
  }
  return 0;
}
