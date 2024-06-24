pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
    let (m, n) = (nums1.len(), nums2.len());
    if m > n {
        return find_median_sorted_arrays(nums2, nums1);
    }
    let (mut high, mut low, total) = (0, m, m + n);
    while low <= high {
        let i = low + (high - low) / 2;
        let j = (total + 1) / 2 - i;
        let left1 = nums1[i - 1];
        let right1 = nums1[i];
        let left2 = nums2[j - 1];
        let right2 = nums2[j];
        if left1 <= right2 && left2 <= right1 {
            match total % 2 {
                0 => {
                    return (std::cmp::max(left1, left2) + std::cmp::min(right1, right2)) as f64
                        / 2.0;
                }
                _ => return std::cmp::max(left1, left2) as f64,
            }
        } else if left2 > right2 {
            high = i - 1;
        } else {
            low = i + 1;
        }
    }
    0.0
}
fn main() {
    dbg!(
        find_median_sorted_arrays(vec![1, 3], vec![2]),
        find_median_sorted_arrays(vec![2], vec![1, 3]),
        find_median_sorted_arrays(vec![1, 2], vec![3, 4])
    );
}
