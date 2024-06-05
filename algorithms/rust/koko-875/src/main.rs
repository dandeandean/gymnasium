pub fn min_eating_speed(piles: Vec<i32>, h: i32) -> i32 {
    let (mut l, &(mut r)) = (1, piles.iter().max().unwrap());
    let mut ret = r;
    while l < r {
        let k = (l + r) / 2;
        let mut tt: f32 = 0.0;
        for &p in &piles {
            tt += (p as f32) / (k as f32);
        }
        if tt < h as f32 {
            ret = k;
            r = k - 1;
        } else {
            l = k + 1;
        }
    }
    ret
}
fn main() {
    dbg!(min_eating_speed(vec![3, 6, 7, 11], 8));
}
