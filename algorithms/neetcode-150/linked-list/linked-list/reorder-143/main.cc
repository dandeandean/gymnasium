// Definition for singly-linked head.
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
    if (!head->next) {
      return;
    }
    ListNode *sprev, *fprev;
    ListNode *slow = head, *fast = head;
    while (fast && fast->next) {
      sprev = slow;
      slow = slow->next;
      fprev = fast->next;
      fast = fast->next->next;
    }
    // separate the two lists
    sprev->next = NULL;
    reverse(slow);
    ListNode *last = fast ? fast : fprev;
    merge(head, last);
  }

private:
  void reverse(ListNode *head) {
    ListNode *prev = NULL;
    ListNode *curr = head;
    ListNode *next = curr->next;
    while (curr) {
      next = curr->next;
      curr->next = prev;
      prev = curr;
      curr = next;
    }
  }
  void merge(ListNode *l1, ListNode *l2) {
    while (l1) {
      ListNode *p1 = l1->next, *p2 = l2->next;
      l1->next = l2;
      if (!p1) {
        break;
      }
      l2->next = p1;
      l1 = p1;
      l2 = p2;
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
  while (cur) {
    std::cout << "(" << cur->val << ")->";
    cur = cur->next;
  }
  std::cout << std::endl;
  return 0;
}
