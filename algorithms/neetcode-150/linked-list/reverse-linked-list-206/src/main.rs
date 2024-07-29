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
    let (mut previous, mut cur) = (None, head);
    while let Some(mut node) = cur {
        let next = node.next;
        node.next = previous;
        previous = Some(node);
        cur = next;
    }
    previous
}
fn main() {
    let mut head: Option<Box<ListNode>> = Some(Box::new(ListNode::new(1)));
    let mut mid: Option<Box<ListNode>> = Some(Box::new(ListNode::new(2)));
    let tail: Option<Box<ListNode>> = Some(Box::new(ListNode::new(3)));
    mid.as_mut().unwrap().next = tail;
    head.as_mut().unwrap().next = mid.clone();
    dbg!(&head);
    dbg!(reverse_list(head));
}
