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
  ListNode *removeNthFromEnd(ListNode *head, int n) {
    ListNode *cur = head;
    int count = 0;
    while (cur) {
      count++;
      cur = cur->next;
    }
    std::cout << count << std::endl;
    return head;
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
  ListNode *cur = a;
  s->removeNthFromEnd(a, 3);
  while (cur) {
    std::cout << "->(" << cur->val << ")";
    cur = cur->next;
  }
  std::cout << std::endl;
  return 0;
}
