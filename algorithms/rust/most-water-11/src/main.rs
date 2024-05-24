pub fn max_area(height: Vec<i32>) -> i32 {
    let mut bestsofar = 0;
    for (left, &h) in height.iter().enumerate() {
        let right = height.len() - left - 1;
        let cap = core::cmp::min(h, height[right]) as usize;
        let dist = height.len() + 1 - left;
        bestsofar = core::cmp::max(cap * dist, bestsofar);
        dbg!(&cap, &bestsofar);
    }
    bestsofar as i32
}
fn main() {
    dbg!(max_area(vec![1, 8, 6, 2, 5, 4, 8, 3, 7])); // 49
    dbg!(max_area(vec![1, 1])); // 1
}
