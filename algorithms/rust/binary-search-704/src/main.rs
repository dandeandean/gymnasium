pub fn search(nums: Vec<i32>, target: i32) -> i32 {
    let (mut left, mut right) = (0, nums.len() - 1);
    while left <= right {
        let i = (left + right) / 2;
        if nums[i] == target {
            return i as i32;
        }
        if nums[i] > target {
            right = i - 1;
        } else {
            left = i + 1;
        }
    }
    return -1;
}
fn main() {
    let nums = vec![5]; //[-1, 0, 3, 5, 9, 12];
    dbg!(search(nums, -5));
}
