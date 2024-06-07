#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}
impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}
// from: https://leetcode.com/problems/reverse-linked-list/solutions/4904416/beats-100-full-solution-explained-with-dry-run-java-c-python-rust-javascript-go/
pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut curr: Option<Box<ListNode>> = head;
    let mut prev: Option<Box<ListNode>> = None;
    while let Some(mut node) = curr {
        let holder = node.next;
        node.next = prev;
        prev = Some(node);
        curr = holder;
    }
    curr
}
fn main() {}
