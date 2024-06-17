fn calc_hours(piles: &Vec<i32>, k: i32) -> i32 {
    let mut ret: i32 = 0;
    for &p in piles {
        if p < k {
            ret += 1;
        } else {
            let mut res: i32 = p / k;
            // check if there's some left over
            if (p % k) != 0 {
                res += 1;
            }
            ret = ret.checked_add(res).unwrap_or(i32::max_value());
        }
    }
    ret
}
pub fn min_eating_speed(piles: Vec<i32>, h: i32) -> i32 {
    let (mut l, mut r): (i32, i32) = (1, piles.iter().max().unwrap().to_owned());
    let mut ret = r;
    while l <= r {
        let k = (l + r) / 2;
        // minimize calc'd hours
        if calc_hours(&piles, k) <= h {
            ret = std::cmp::min(ret, k);
            r = k - 1;
        } else {
            l = k + 1;
        }
    }
    ret
}

fn main() {
    dbg!(min_eating_speed(vec![3, 6, 7, 11], 8));
    dbg!(min_eating_speed(vec![30, 11, 23, 4, 20], 5));
    dbg!(min_eating_speed(vec![30, 11, 23, 4, 20], 6));
    let big = vec![
        332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077,
        337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285,
        629455728, 941802184,
    ];
    let large = 823855818;
    dbg!(min_eating_speed(big, large));
    dbg!(min_eating_speed(
        vec![805306368, 805306368, 805306368],
        1000000000
    ));
}
