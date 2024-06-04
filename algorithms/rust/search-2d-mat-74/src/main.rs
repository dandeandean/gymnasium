pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
    let nums: Vec<i32> = matrix
        .iter()
        .flat_map(|array| array.iter())
        .cloned()
        .collect();
    let (mut left, mut right) = (0, nums.len() - 1);
    while left <= right {
        let i = (left + right) / 2;
        if nums[i] == target {
            return true;
        }
        if nums[i] > target {
            if i == 0 {
                return false;
            }
            right = i - 1;
        } else {
            left = i + 1;
        }
    }
    return false;
}
fn main() {
    let matrix = vec![vec![1, 3, 5, 7], vec![10, 11, 16, 20], vec![23, 30, 34, 60]];
    dbg!(search_matrix(matrix, 3));
    let matrix = vec![vec![1]];
    dbg!(search_matrix(matrix, 1));
    let matrix = vec![vec![1], vec![3], vec![5]];
    dbg!(search_matrix(matrix, 3));
}
