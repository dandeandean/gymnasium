pub fn character_replacement(s: String, k: i32) -> i32 {
    let (mut l, mut r, mut best): (usize, usize, i32) = (0, 1, 0);
    let mut mk = k;
    while r < s.len() {
        if s.chars().nth(l) != s.chars().nth(r) {
            mk -= 1;
        }
        if mk < 0 {
            l += 1;
        }
        best = std::cmp::max(best, 1 + r as i32 - l as i32);
        dbg!(r, l);
        r += 1;
    }
    std::cmp::max(best, 1 + r as i32 - l as i32)
}
fn main() {
    dbg!(
        character_replacement("abab".to_string(), 2),
        character_replacement("AABABBA".to_string(), 1),
        character_replacement("abcabcde".to_string(), 2),
        character_replacement("ABAA".to_string(), 0)
    );
}
