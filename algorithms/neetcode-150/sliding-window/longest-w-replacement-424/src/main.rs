pub fn character_replacement(s: String, k: i32) -> i32 {
    let mut counts: std::collections::HashMap<char, i32> = std::collections::HashMap::new();
    let (mut l, mut r, mut maxf): (usize, usize, i32) = (0, 0, 0);
    while r < s.len() {
        let c = s.chars().nth(r).unwrap();
        *counts.entry(c).or_insert(0) += 1;
        maxf = std::cmp::max(maxf, counts[&c]);
        if (r as i32 - l as i32 + 1) - maxf > k {
            *counts.entry(s.chars().nth(l).unwrap()).or_insert(1) -= 1;
            l += 1;
        }
        r += 1;
    }
    r as i32 - l as i32
}
fn main() {
    dbg!(
        character_replacement("abab".to_string(), 2),
        character_replacement("AABABBA".to_string(), 1),
        character_replacement("abcabcde".to_string(), 2),
        character_replacement("ABAA".to_string(), 0)
    );
}
