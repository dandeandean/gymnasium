#include <iostream>
/* Definition for singly-linked list.*/
struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
  ListNode *removeNthFromEndDumb(ListNode *head, int n) {
    ListNode *cur = head, *prev = NULL;
    int count = 0;
    while (cur) {
      count++;
      cur = cur->next;
    }
    if (count == n) {
      return head->next;
    }
    cur = head;
    while (count > n) {
      count--;
      prev = cur;
      cur = cur->next;
    }
    if (prev) {
      prev->next = cur->next;
    }
    return head;
  }
  ListNode *removeNthFromEnd(ListNode *head, int n) {
    if (!head->next) {
      return NULL;
    }
    ListNode *funny = new ListNode(0, head);
    ListNode *fast = funny, *slow = head;
    for (int i = 0; i <= n; i++) {
      fast = fast->next;
    }
    while (fast) {
      fast = fast->next;
      slow = slow->next;
    }
    if (slow->next) {
      slow->next = slow->next->next;
    }
    return funny->next;
  }
};

int main() {
  ListNode *f = new ListNode(6);
  ListNode *e = new ListNode(5, f);
  ListNode *d = new ListNode(4, e);
  ListNode *c = new ListNode(3, d);
  ListNode *b = new ListNode(2, c);
  ListNode *a = new ListNode(1, b);
  Solution *s = new Solution;
  int n = 3;
  std::cout << "n=" << n << std::endl;
  s->removeNthFromEnd(a, n);
  ListNode *cur = a;
  while (cur) {
    std::cout << "->(" << cur->val << ")";
    cur = cur->next;
  }
  std::cout << std::endl;
  return 0;
}
