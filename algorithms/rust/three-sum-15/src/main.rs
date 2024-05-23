pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut out: Vec<Vec<i32>> = vec![];
    let mut nums_sorted = nums.clone();
    nums_sorted.sort();
    for (i, &num) in nums_sorted.iter().enumerate() {
        if num > 0 {
            break;
        }
        if i > 0 && num == nums_sorted[i - 1] {
            continue;
        }
        let (mut left, mut right): (usize, usize) = (i + 1, nums_sorted.len() - 1);
        while left < right {
            let sumsofar = num + nums_sorted[left] + nums_sorted[right];
            if sumsofar > 0 {
                right -= 1;
            } else if sumsofar < 0 {
                left += 1;
            } else {
                out.push(vec![num, nums_sorted[left], nums_sorted[right]]);
                left += 1;
                right -= 1;
                while nums_sorted[left] == nums_sorted[left - 1] && left < right {
                    left += 1;
                }
            }
        }
    }
    out
}

fn main() {
    let nums = vec![-1, 0, 1, 2, -1, -4];
    let resp = three_sum(nums);
    dbg!(resp);
}
