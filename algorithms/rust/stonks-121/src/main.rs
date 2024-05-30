pub fn max_profit(prices: Vec<i32>) -> i32 {
    let (mut best, mut i, mut j) = (0, 0, 1);
    while j < prices.len() {
        if prices[i] < prices[j] {
            best = std::cmp::max(best, prices[j] - prices[i]);
        } else {
            i = j;
        }
        j += 1;
    }
    best
}
fn main() {
    dbg!(max_profit(vec![7, 1, 5, 3, 6, 4]));
    dbg!(max_profit(vec![7, 6, 4, 3, 1]));
}
