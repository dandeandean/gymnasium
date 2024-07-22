// Definition for singly-linked list.
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
    ListNode *cur = head;
    std::vector<ListNode *> stack;
    while (cur) { // First Pass to init stack
      stack.push_back(cur);
      cur = cur->next;
    }
    cur = head;
    while (cur) {
      ListNode *back = stack.back();
      if (!stack.empty()) {
        stack.pop_back();
        // stack.back() = nullptr;
      }
      if (back) {
        back->next = cur->next;
        ListNode *old_next = cur->next;
        cur->next = back;
        cur = old_next->next;
      }
      cur = cur->next;
      if (back == cur) {
        std::cout << "(" << cur->val << ")" << std::endl;
        // cur->next = nullptr;
        return;
      }
    }
  }
};
int main(void) {
  ListNode *d = new ListNode(4);
  ListNode *c = new ListNode(3, d);
  ListNode *b = new ListNode(2, c);
  ListNode *a = new ListNode(1, b);
  Solution *s = new Solution;
  ListNode *cur = a;
  while (cur) {
    std::cout << "(" << cur->val << ")->";
    cur = cur->next;
  }
  std::cout << std::endl;
  s->reorderList(a);
  std::cout << "(" << a->val << ")->" << a->next->val << "\n";
  std::cout << "(" << b->val << ")->" << b->next->val << "\n";
  std::cout << "(" << c->val << ")->" << c->next->val << "\n";
  std::cout << "(" << d->val << ")->" << d->next->val << "\n";
  cur = a;
  // while (cur) {
  //   std::cout << cur->val;
  //   cur = cur->next;
  // }
  return 0;
}
