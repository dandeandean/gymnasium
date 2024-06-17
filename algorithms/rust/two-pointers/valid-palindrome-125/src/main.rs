pub fn is_palindrome(s: String) -> bool {
    let ss: String = s
        .chars()
        .filter(|c| c.is_alphanumeric())
        .map(|c| c.to_ascii_lowercase())
        .collect();
    let mut left = 0;
    let mut right = ss.len().checked_sub(1).unwrap_or(0);
    dbg!(&ss, left, right);
    while left < right {
        if ss.chars().nth(right).unwrap() != ss.chars().nth(left).unwrap() {
            return false;
        }
        if right == 0 {
            break;
        }
        left += 1;
        right -= 1
    }
    true
}
fn main() {
    dbg!(is_palindrome("aa".to_string()));
}
