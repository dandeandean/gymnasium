pub fn trap(height: Vec<i32>) -> i32 {
    //let (mut i, mut j): (usize, usize) = (0, height.len() - 1);
    let mut from_left: Vec<i32> = vec![];
    let mut from_right: Vec<i32> = vec![];
    let (mut left_max, mut right_max): (i32, i32) = (0, 0);
    for (i, &h) in height.iter().enumerate() {
        left_max = std::cmp::max(h, left_max);
        right_max = std::cmp::max(height[height.len() - (i + 1)], right_max);
        from_left.push(left_max);
        from_right.insert(0, right_max);
    }
    let mut ret = 0;
    for (i, &h) in height.iter().enumerate() {
        ret += std::cmp::min(from_left[i], from_right[i]) - h;
    }
    ret
}
fn main() {
    let height = vec![0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1];
    dbg!(trap(height));
}
