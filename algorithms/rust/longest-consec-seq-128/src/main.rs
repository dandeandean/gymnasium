use std::collections::HashSet;

pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
    let set: HashSet<&i32> = HashSet::from_iter(nums.iter());
    let mut starts: Vec<&i32> = vec![];
    for &num in set.iter() {
        if !set.contains(&(num - 1)) {
            starts.push(num);
        }
    }
    let mut best_count: i32 = 0;
    for &start in starts {
        let mut count = 0;
        while set.contains(&(start + count)) {
            count += 1;
        }
        if count > best_count {
            best_count = count;
        }
    }
    return best_count;
}
fn main() {
    let nums = vec![0, 3, 7, 2, 5, 8, 4, 6, 0, 1];
    let ans = longest_consecutive(nums);
    dbg!(ans);
}
