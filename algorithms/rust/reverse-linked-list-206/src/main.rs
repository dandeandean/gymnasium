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
pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    // from: https://leetcode.com/problems/reverse-linked-list/solutions/4904416/beats-100-full-solution-explained-with-dry-run-java-c-python-rust-javascript-go/
    let mut dummy = None;
    let mut current = head;
    while let Some(mut node) = current {
        let next = node.next.take();
        node.next = dummy.take();
        dummy = Some(node);
        current = next;
    }
    dummy
}
fn main() {}
