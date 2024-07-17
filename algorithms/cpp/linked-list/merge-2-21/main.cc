// Definition for singly-linked list.
struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};
ListNode *mergeTwoLists(ListNode *list1, ListNode *list2) {
  ListNode *current, *anchor = new ListNode;
  // ListNode *anchor = current;
  while (list1 && list2) {
    if (list1->val <= list2->val) {
      current->next = list1;
      list1 = list1->next;
    } else {
      current->next = list2;
      list2 = list2->next;
    }
    current = current->next;
  }
  if (list1) {
    current->next = list1;
  } else {
    current->next = list2;
  }
  return anchor->next;
}

int main(void) {

  ListNode *list1 = new ListNode(1);
  list1->next = new ListNode(2);
  ListNode *list2 = new ListNode(2);
  list2->next = new ListNode(5);
  ListNode *out = new ListNode;
  out = mergeTwoLists(list1, list2);
  return 0;
}
