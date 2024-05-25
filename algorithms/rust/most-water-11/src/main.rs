fn calc_area(height: &Vec<i32>, i: usize, j: usize) -> i32 {
    let mut cap = height[i];
    if height[i] < cap {
        cap = height[i];
    }
    let dist = (j as i32) - (i as i32);
    cap * dist
}
pub fn max_area(height: Vec<i32>) -> i32 {
    let (mut left, mut right) = (0, height.len() - 1);
    let mut bestsofar = calc_area(&height, left, right);
    while left < right {
        if height[left] > height[right] {
            right -= 1;
        }
        if height[right] > height[left] {
            left += 1;
        }
        if height[right] == height[left] {
            right -= 1;
            left += 1;
        }
        if right == left || left == height.len() {
            break;
        }
        bestsofar = core::cmp::max(calc_area(&height, left, right), bestsofar);
    }
    bestsofar
}
fn main() {
    //dbg!(max_area(vec![1, 8, 6, 2, 5, 4, 8, 3, 7])); // 49
    dbg!(max_area(vec![1, 2])); // 1
}
