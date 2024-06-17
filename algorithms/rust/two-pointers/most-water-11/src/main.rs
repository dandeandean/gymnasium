fn calc_area(height: &Vec<i32>, i: i32, j: i32) -> i32 {
    let mut cap = height[j as usize];
    if height[i as usize] < cap {
        cap = height[i as usize];
    }
    let dist = (j as i32) - (i as i32);
    let res = cap * dist;
    res
}
pub fn max_area(height: Vec<i32>) -> i32 {
    let (mut left, mut right) = (0, height.len() - 1);
    let mut bestsofar = calc_area(&height, left as i32, right as i32);
    while left < right {
        let current_area = calc_area(&height, left as i32, right as i32);
        if current_area > bestsofar {
            bestsofar = current_area;
        }
        if height[left] > height[right] {
            right -= 1;
        } else {
            left += 1;
        }
    }
    bestsofar
}
fn main() {
    dbg!(max_area(vec![1, 8, 6, 2, 5, 4, 8, 3, 7])); // 49
    dbg!(max_area(vec![1, 2])); // 1
}
