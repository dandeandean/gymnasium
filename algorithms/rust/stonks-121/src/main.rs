pub fn max_profit(prices: Vec<i32>) -> i32 {
    let mut left: usize = 0;
    let mut right: usize = 1;
    let mut best: i32 = 0;
    while right < prices.len() {
        if prices[right] < prices[left] {
            left = right;
        } else {
            //calc POSSIBLE new best
            best = std::cmp::max(best, prices[right] - prices[left]);
        }
        // window fattens
        right += 1;
    }
    best
}
fn main() {
    dbg!(max_profit(vec![7, 1, 5, 3, 6, 4]));
    dbg!(max_profit(vec![7, 6, 4, 3, 1]));
}
