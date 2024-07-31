/**
 * Definition for singly-linked list.
 **/
struct ListNode {
  int val;
  ListNode *next;
  ListNode()
      : val(0), next(nullptr){} * ListNode(int x)
      : val(x),
        next(nullptr){} * ListNode(int x, ListNode *next)
      : val(x),
        next(next) {}
};

class Solution {
public:
  ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
    ListNode *anchor = new ListNode;
    ListNode *cur = anchor;
    int carry = 0;
    while (l1 || l2 || carry) {
      int add_res = carry;
      if (l1) {
        add_res += l1->val;
        l1 = l1->next;
      }
      if (l2) {
        add_res += l2->val;
        l2 = l2->next;
      }
      carry = add_res / 10;
      cur->next = new ListNode(add_res % 10);
      cur = cur->next;
    }
    return anchor->next;
  }
};
int main(void) { return 0; }
