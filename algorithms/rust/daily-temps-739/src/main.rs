pub fn daily_temperatures(temperatures: Vec<i32>) -> Vec<i32> {
    let mut temps_out: Vec<i32> = vec![0; temperatures.iter().len()];
    let mut stack: Vec<(usize, &i32)> = vec![];
    for (i, t) in temperatures.iter().enumerate() {
        if !stack.is_empty() {
            while t > stack.last().unwrap().1 {
                let popped = stack.pop().unwrap();
                temps_out[popped.0] = (i - popped.0) as i32;
                if stack.is_empty() {
                    break;
                }
            }
        }
        stack.push((i, t))
    }
    temps_out
}
fn main() {
    dbg!(daily_temperatures(vec![73, 74, 75, 71, 69, 72, 76, 73]));
    dbg!(daily_temperatures(vec![30, 40, 50, 60]));
}
