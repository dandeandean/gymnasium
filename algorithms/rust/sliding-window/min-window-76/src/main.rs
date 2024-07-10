pub fn min_window(s: String, t: String) -> String {
    "abc".to_string()
}
fn main() {
    dbg!(
        min_window("ADOBECODEBANC".to_string(), "ABC".to_string()),
        min_window("A".to_string(), "A".to_string()),
        min_window("AA".to_string(), "A".to_string()),
    );
}
