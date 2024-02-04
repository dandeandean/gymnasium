fn can_place_flowers(flowerbed: Vec<i32>, n: i32) -> bool {
    let mut count: i32 = 1;
    let mut m: i32 = 0;
    for b in flowerbed {
        if b == 1 {
            m += (count - 1)/2;
            if m > n {
                return true;
            }
            count = 0;
        }
        else{
            count += 1;
        }
    }
    m += count/2;

    return m >= n
}
fn main() {

    let test_vec = vec![0,1,0,0,1];
    let go = can_place_flowers(test_vec, 2);
    println!("bool: {}",go);
}
