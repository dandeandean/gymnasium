use std::{
    collections::{BinaryHeap, HashMap},
    ops::Add,
};

pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
    let mut counts: HashMap<i32, i32> = HashMap::new();
    for n in nums {
        let count_at_n = counts.entry(n).or_insert(0);
        *count_at_n += 1;
    }
    let mut heap: BinaryHeap<(&i32, &i32)> = BinaryHeap::new();
    for k in counts.keys() {
        // gotta be in this order so it sort them correctly
        heap.push((counts.get(k).unwrap(), k))
    }
    let mut out: Vec<i32> = vec![];
    for _ in 0..k {
        out.push(heap.pop().unwrap().1.to_owned())
    }
    out
}
fn main() {}
