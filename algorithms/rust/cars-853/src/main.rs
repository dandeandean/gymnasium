use std::iter::zip;

pub fn car_fleet(target: i32, position: Vec<i32>, speed: Vec<i32>) -> i32 {
    // Go through all of the position, speed pairs and calculate their arrival times
    let mut pairs: Vec<(i32, i32)> = zip(position, speed).collect();
    pairs.sort();
    // dbg!(&pairs);
    let mut times_left: Vec<i32> = vec![];
    for (start, speed) in pairs {
        // calculate the 'ticks' it will take to get to the target
        // i32 should be good enough
        let time: i32 = (target - start) / speed;
        times_left.push(time);
    }
    let mut out: i32 = 1;
    for i in (0..times_left.len()).rev() {
        if times_left.get(i) > times_left.get(i + 1) {
            out += 1
        }
    }
    dbg!(&times_left);
    out
}
fn main() {
    let target = 12;
    let position = vec![6, 8];
    let speed = vec![3, 2];
    dbg!(car_fleet(target, position, speed));
}
