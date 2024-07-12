pub fn min_window(s: String, t: String) -> String {
    if t == "".to_string() {
        return t;
    }
    let mut t_hash: std::collections::HashMap<char, i32> =
        std::collections::HashMap::with_capacity(t.len());
    dbg!(t_hash.len());
    for c in t.chars() {
        *t_hash.entry(c).or_insert(0) += 1;
    }
    "abc".to_string()
}
fn main() {
    dbg!(
        min_window("ADOBECODEBANC".to_string(), "ABC".to_string()),
        min_window("A".to_string(), "A".to_string()),
        min_window("AA".to_string(), "A".to_string()),
    );
}
