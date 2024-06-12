pub fn find_min(nums: Vec<i32>) -> i32 {
    let (mut l, mut r) = (0, nums.len().checked_sub(1).unwrap_or(0));
    let mut ret = i32::max_value();
    while l <= r {
        let m = (l + r) / 2;
        ret = std::cmp::min(ret, nums[m]);
        if nums[l] < nums[r] || l == r {
            return std::cmp::min(nums[l], ret);
        }
        if nums[m] > nums[r] {
            // the edge is somewhere between mid & right
            l = m + 1;
        } else {
            r = m.checked_sub(1).unwrap_or(0);
        }
    }
    //nums.iter().min().unwrap().to_owned()
    std::cmp::min(ret, nums[0])
}
fn main() {
    dbg!(find_min(vec![3, 4, 5, 1, 2]));
    dbg!(find_min(vec![4, 5, 6, 7, 0, 1, 2]));
    dbg!(find_min(vec![11, 13, 15, 17]));
    dbg!(find_min(vec![2, 1]));
    dbg!(find_min(vec![1]));
}
