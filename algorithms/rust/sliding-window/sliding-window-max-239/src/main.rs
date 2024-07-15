pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
    let mut out: Vec<i32> = vec![];
    let (mut l, mut r) = (0, (k - 1) as usize);
    let mut q: std::collections::VecDeque<usize> = std::collections::VecDeque::new();
    while r < nums.len() {
        // there's gotta be a better way of doing this ......
        while !q.is_empty() && nums[*q.iter().last().unwrap() as usize] < nums[r] {
            //
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
fn main() {
    dbg!(
        max_sliding_window(vec![1, 3, -1, -3, 5, 3, 6, 7], 3),
        max_sliding_window(vec![1], 1)
    );
}
