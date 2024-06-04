pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
    let (mut left, mut right) = (0, matrix.len() - 1);
    let mut row_ptr = 0;
    while left <= right {
        let mid = (left + right) / 2;
        if matrix[mid][0] >= target {
            right = mid;
            if right > 0 {
                right -= 1;
            } else {
                return matrix[mid][0] == target;
            }
        } else {
            left = mid + 1;
        }
        row_ptr = mid;
    }
    let (mut left, mut right) = (0, matrix[row_ptr].len() - 1);
    while left <= right {
        let mid = (left + right) / 2;
        if matrix[row_ptr][mid] == target {
            return true;
        }
        if matrix[row_ptr][mid] > target {
            right = mid;
            if right > 0 {
                right -= 1;
            } else {
                return false;
            }
        } else {
            left = mid + 1;
        }
    }
    false
}
fn main() {
    let matrix = vec![vec![1, 3, 5, 7], vec![10, 11, 16, 20], vec![23, 30, 34, 60]];
    dbg!(search_matrix(matrix, 3));
    let matrix = vec![vec![1]];
    dbg!(search_matrix(matrix, 1));
    let matrix = vec![vec![1], vec![3], vec![5]];
    dbg!(search_matrix(matrix, 3));
}
