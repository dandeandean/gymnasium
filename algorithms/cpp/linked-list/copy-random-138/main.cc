// Definition for a Node.
#include <cstddef>
#include <iostream>
#include <unordered_map>
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
    std::vector<Node *> nodes;
    std::unordered_map<Node *, Node *> node_map;
    while (cur) {
      Node *new_node = new Node(cur->val);
      node_map[cur] = new_node;
      cur = cur->next;
    }
    cur = head;
    // assign random & next to new nodes
    while (cur) {
      Node *new_node = node_map[cur];
      new_node->next = node_map[cur->next];
      new_node->random = node_map[cur->random];
      cur = cur->next;
    }
    std::cout << "(out)" << std::endl;
    return node_map[head];
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
