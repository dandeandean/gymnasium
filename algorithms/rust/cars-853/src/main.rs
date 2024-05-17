use std::iter::zip;
pub fn car_fleet(target: i32, position: Vec<i32>, speed: Vec<i32>) -> i32 {
    // Go through all of the position, speed pairs and calculate their arrival times
    let mut pairs: Vec<(i32, i32)> = zip(position, speed).collect();
    let mut times_left: Vec<f32> = vec![];
    pairs.sort();
    pairs.reverse();
    for (start, speed) in pairs {
        times_left.push((target as f32 - start as f32) / speed as f32);
        if times_left.len() >= 2 {
            let last = times_left.get(times_left.len() - 1).unwrap();
            let second = times_left.get(times_left.len() - 2).unwrap();
            if last <= second {
                times_left.pop();
            }
        }
    }
    times_left.len() as i32
}
fn main() {
    let target = 10;
    let position = vec![0, 4, 2];
    let speed = vec![2, 1, 3];
    dbg!(car_fleet(target, position, speed));
}
