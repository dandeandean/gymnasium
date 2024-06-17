use std::iter::zip;
pub fn car_fleet(target: i32, position: Vec<i32>, speed: Vec<i32>) -> i32 {
    // Go through all of the position, speed pairs and calculate their arrival times
    let mut pairs: Vec<(i32, i32)> = zip(position, speed).collect();
    let mut times_left: Vec<f32> = vec![];
    pairs.sort();
    pairs.reverse();
    for (start, speed) in pairs {
        let speed_new: f32 = (target as f32 - start as f32) / speed as f32;
        match times_left.last() {
            None => times_left.push(speed_new),
            Some(last) => {
                if last < &speed_new {
                    times_left.push(speed_new);
                }
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
