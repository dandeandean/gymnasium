---
tags:
  - neetcode
---
# Problem 
Given an integer array `nums` and an integer k, return the k most frequent elements. You may return the answer in any order.

## Example 1:
```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```
## Example 2:
```
Input: nums = [1], k = 1
Output: [1]
```
 
## Constraints:

`1 <= nums.length <= 105`
`-104 <= nums[i] <= 104`
`k is in the range [1, the number of unique elements in the array].`

It is guaranteed that the answer is unique.
 

## Follow up:
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

# Solution
So my first thought is to take the array, and hash it so that the key is the unique number in `nums` and the value is how many times we see it in `nums`.

We just add all of the `(value, key)` from the hash map into a binary heap and pop elements for the range 0..k.

```Rust
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
```
