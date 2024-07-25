// Definition for a Node.
#include <cstddef>
#include <iostream>
#include <vector>
class Node {
public:
  int val;
  Node *next;
  Node *random;

  Node(int _val) {
    val = _val;
    next = NULL;
    random = NULL;
  }
};

class Solution {
public:
  Node *copyRandomList(Node *head) {
    Node *cur = head;
    std::vector<Node *> v;
    while (cur) {
      v.push_back(new Node(cur->val));
      cur = cur->next;
    }
    return v.front();
  }
};

int main(void) {
  Node *a = new Node(1);
  Node *b = new Node(2);
  Node *c = new Node(3);
  Node *d = new Node(4);
  a->next = b;
  b->next = c;
  c->next = d;

  a->random = c;
  b->random = NULL;
  c->random = d;
  d->random = d;

  Solution *s = new Solution;
  s->copyRandomList(a);
  Node *cur = a;
  while (cur) {
    std::cout << "(" << cur->val << ")" << std::endl;
    cur = cur->next;
  }

  return 0;
}
