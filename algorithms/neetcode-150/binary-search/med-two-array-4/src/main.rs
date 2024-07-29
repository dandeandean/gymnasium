pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
    let (m, n) = (nums1.len(), nums2.len());
    if m > n {
        return find_median_sorted_arrays(nums2, nums1);
    }
    let (mut high, mut low, total) = (m, 0, m + n);
    let mut result = 0.0;
    while low <= high {
        let i = low + (high - low) / 2;
        let j = (total + 1) / 2 - i;
        let left1 = match i {
            0 => i32::min_value(),
            _ => nums1[i - 1],
        };
        let left2 = match j {
            0 => i32::min_value(),
            _ => nums2[j - 1],
        };
        let right1 = nums1.get(i).unwrap_or(&i32::max_value()).clone();
        let right2 = nums2.get(j).unwrap_or(&i32::max_value()).clone();
        if left1 <= right2 && left2 <= right1 {
            match total % 2 {
                0 => {
                    result =
                        (std::cmp::max(left1, left2) + std::cmp::min(right1, right2)) as f64 / 2.0;
                }
                _ => result = std::cmp::max(left1, left2) as f64,
            }
            break;
        } else if left1 > right2 {
            high = i.checked_sub(1).unwrap_or(0);
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
