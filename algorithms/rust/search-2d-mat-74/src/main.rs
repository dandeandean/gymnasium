pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
    matrix
        .iter()
        .flat_map(|array| array.iter())
        .cloned()
        .collect::<Vec<i32>>()
        .contains(&target)
}
fn main() {
    let matrix = vec![vec![1, 3, 5, 7], vec![10, 11, 16, 20], vec![23, 30, 34, 60]];
    dbg!(search_matrix(matrix, 3));
    let matrix = vec![vec![1]];
    dbg!(search_matrix(matrix, 1));
    let matrix = vec![vec![1], vec![3], vec![5]];
    dbg!(search_matrix(matrix, 3));
}
