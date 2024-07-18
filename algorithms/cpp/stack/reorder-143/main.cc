// Definition for singly-linked list.
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
    while (cur) {
      cur = cur->next;
    }
  }
};
int main(void) { return 0; }
