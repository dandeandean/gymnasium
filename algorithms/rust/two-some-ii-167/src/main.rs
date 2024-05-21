pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
    let mut left: usize = 0;
    let mut right: usize = numbers.len() - 1;
    // potential problem with over/underflow of usize
    while left < right {
        if numbers[left] + numbers[right] == target {
            return vec![left as i32 + 1, right as i32 + 1];
        }
        if numbers[left] + numbers[right] > target {
            right -= 1;
        }
        if numbers[left] + numbers[right] < target {
            left += 1;
        }
    }
    numbers
}

fn main() {
    let numbers = vec![2, 7, 11, 15];
    let target = 9;
    dbg!(two_sum(numbers, target));
}
