pub fn search(nums: Vec<i32>, target: i32) -> i32 {
    let (mut l, mut r) = (0, nums.len() - 1);
    while l < r {
        let m = (l + r) / 2;
        l = m + 1;
        r = m.checked_sub(1).unwrap_or(0);
    }
    -1
}
fn main() {
    dbg!(search(vec![4, 5, 6, 7, 0, 1, 2], 0));
    dbg!(search(vec![4, 5, 6, 7, 0, 1, 2], 3));
    dbg!(search(vec![1], 0));
    dbg!(search(vec![1], 1));
}
