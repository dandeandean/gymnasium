pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
    let (m, n) = (nums1.len(), nums2.len());
    if m > n {
        return find_median_sorted_arrays(nums2, nums1);
    }
    let (mut high, mut low, total) = (0, m, m + n);
    let mut result = 0.0;
    while low <= high {
        let i = low + (high - low) / 2;
        let j = (total + 1) / 2 - i;
        let mut left1 = nums1[i - 1];
        if i == 0 {
            left1 = i32::min_value();
        }
        let mut right1 = nums1[i];
        if i > m {
            right1 = i32::max_value();
        }
        let mut left2 = nums2[j - 1];
        if j == 0 {
            left2 = i32::min_value();
        }
        let mut right2 = nums2[j];
        if j > n {
            right2 = i32::max_value();
        }
        if left1 <= right2 && left2 <= right1 {
            match total % 2 {
                0 => {
                    result =
                        (std::cmp::max(left1, left2) + std::cmp::min(right1, right2)) as f64 / 2.0;
                }
                _ => result = std::cmp::max(left1, left2) as f64,
            }
        } else if left2 > right2 {
            high = i - 1;
        } else {
            low = i + 1;
        }
    }
    result
}
fn main() {
    dbg!(
        find_median_sorted_arrays(vec![1, 3], vec![2]),
        find_median_sorted_arrays(vec![2], vec![1, 3]),
        find_median_sorted_arrays(vec![1, 2], vec![3, 4])
    );
}
