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
    let mut start: Option<Box<ListNode>> = None;
    let mut cur: Option<Box<ListNode>> = head;
    let mut mem: Vec<Option<Box<ListNode>>> = vec![];
    while cur.is_some() {
        mem.push(cur.clone());
        cur = cur
            .unwrap() // we must be able to unwrap it as we know it's Some from above ^^
            .next;
    }
    while !mem.is_empty() {
        dbg!(mem.pop());
    }
    start
}
fn main() {}
