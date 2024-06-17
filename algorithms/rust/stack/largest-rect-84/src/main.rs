use std::cmp;
pub fn largest_rectangle_area(heights: Vec<i32>) -> i32 {
    let mut stack: Vec<(usize, i32)> = vec![];
    let mut max_so_far: i32 = 0;
    for (i, &h) in heights.iter().enumerate() {
        let mut start_of_block = i;
        while !stack.is_empty() && stack.last().unwrap().1 > h {
            let (index, height) = stack.pop().unwrap();
            max_so_far = cmp::max(max_so_far, height * (i - index) as i32);
            start_of_block = index;
        }
        stack.push((start_of_block, h))
    }
    let len_height = heights.len();
    for (i, h) in stack {
        let thickness = len_height as i32 - i as i32;
        max_so_far = cmp::max(max_so_far, h * thickness);
    }
    max_so_far
}
fn main() {
    let heights = vec![2, 1, 5, 6, 2, 3];
    dbg!(largest_rectangle_area(heights));
}
