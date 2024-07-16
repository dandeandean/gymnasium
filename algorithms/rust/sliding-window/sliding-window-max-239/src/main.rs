pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
    let mut out: Vec<i32> = vec![];
    let (mut l, mut r) = (0, 0);
    let mut q: std::collections::VecDeque<usize> = std::collections::VecDeque::new();
    while r < nums.len() {
        // there's gotta be a better way of doing this ......
        while !q.is_empty() && nums[*q.iter().last().unwrap() as usize] < nums[r] {
            q.pop_back();
        }
        q.push_back(r);
        if &l > q.get(0).unwrap() {
            q.pop_front();
        }
        if (r + 1) as i32 >= k {
            out.push(nums[*q.get(0).unwrap()]);
            l += 1;
        }
        r += 1;
    }
    out
}

pub fn dumb_max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
    // too  slow to pass leetcode
    let mut out: Vec<i32> = vec![];
    let (mut l, mut r) = (0, (k) as usize);
    // LEQ || l..=r
    while r <= nums.len() {
        out.push(*nums[l..r].iter().max().unwrap());
        l += 1;
        r += 1;
    }
    out
}
fn main() {
    dbg!(
        dumb_max_sliding_window(vec![1, 3, -1, -3, 5, 3, 6, 7], 3),
        dumb_max_sliding_window(vec![1], 1)
    );
}
