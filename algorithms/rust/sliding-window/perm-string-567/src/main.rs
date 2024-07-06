pub fn check_inclusion(s1: String, s2: String) -> bool {
    if s1.len() > s2.len() {
        return false;
    }
    let mut s1count: std::collections::HashMap<char, i32> = std::collections::HashMap::new();
    let mut s2count: std::collections::HashMap<char, i32> = std::collections::HashMap::new();
    let (mut l, mut r): (usize, usize) = (0, 0);
    for c in s1.chars() {
        *s1count.entry(c).or_insert(0) += 1;
        *s2count.entry(s2.chars().nth(r).unwrap()).or_insert(0) += 1;
        r += 1;
    }
    if s2count == s1count {
        return true;
    }
    while r < s2.len() {
        *s2count.entry(s2.chars().nth(r).unwrap()).or_insert(0) += 1;
        if r - l + 1 > s1.len() {
            let key = s2.chars().nth(l).unwrap();
            *s2count.entry(key).or_insert(0) -= 1;
            if s2count[&key] == 0 {
                s2count.remove_entry(&key);
            }
        }
        if s2count == s1count {
            return true;
        }
        l += 1;
        r += 1;
    }
    s2count == s1count
}
fn main() {
    dbg!(
        check_inclusion("ab".to_string(), "eidbaooo".to_string()),
        check_inclusion("ab".to_string(), "eidboaoo".to_string()),
        check_inclusion("adc".to_string(), "dcda".to_string()),
        check_inclusion("a".to_string(), "ab".to_string()),
    );
}
