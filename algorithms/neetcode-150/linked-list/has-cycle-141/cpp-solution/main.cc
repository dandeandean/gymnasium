// Definition for singly-linked list.
#import <unordered_set>
struct ListNode {
  int val;
  ListNode *next;
  ListNode(int x) : val(x), next(NULL) {}
};
class Solution {
public:
  bool hasCycle(ListNode *head) {
    ListNode *cur = head;
    std::unordered_set<ListNode *> been_to;
    while (cur) {
      if (been_to.count(cur->next)) {
        return true;
      }
      been_to.insert(cur);
      cur = cur->next;
    }
    return false;
  }
};
int main() { return 0; }
