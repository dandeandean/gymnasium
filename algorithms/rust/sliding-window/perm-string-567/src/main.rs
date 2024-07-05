pub fn check_inclusion(s1: String, s2: String) -> bool {
    if s1.len() > s2.len() {
        return false;
    }
    let (mut l, mut r): (usize, usize) = (0, s1.len());
    while r < s2.len() {
        s2.chars().nth(l);
        l += 1;
        r += 1;
    }
    true
}
fn main() {
    dbg!(
        check_inclusion("ab".to_string(), "eidbaooo".to_string()),
        check_inclusion("ab".to_string(), "eidboaoo".to_string()),
    );
}
