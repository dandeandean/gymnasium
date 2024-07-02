pub fn character_replacement(s: String, k: i32) -> i32 {
    let (mut l, mut best): (usize, i32) = (1, 0);
    let mut counts: std::collections::HashMap<char, i32> = std::collections::HashMap::new();
    let mut maxf: i32;
    //while r < s.len() {
    for (r, c) in s.chars().enumerate() {
        *counts.entry(c).or_insert(0) += 1;
        maxf = std::cmp::max(*counts.values().into_iter().max().unwrap(), counts[&c]);
        if r as i32 - l as i32 - maxf + 1 > k {
            *counts.entry(c).or_insert(0) -= 1;
            l += 1;
        }
        best = std::cmp::max(best, r as i32 - l as i32 + 1);
    }
    best
}
fn main() {
    dbg!(
        character_replacement("abab".to_string(), 2),
        character_replacement("AABABBA".to_string(), 1),
        character_replacement("abcabcde".to_string(), 2),
        character_replacement("ABAA".to_string(), 0)
    );
}
