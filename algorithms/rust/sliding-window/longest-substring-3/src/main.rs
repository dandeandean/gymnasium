pub fn length_of_longest_substring(s: String) -> i32 {
    let mut used: std::collections::HashSet<char> = std::collections::HashSet::new();
    let (mut ret, mut con): (i32, i32) = (0, 0);
    for c in s.chars() {
        if !used.contains(&c) {
            con += 1;
            ret = std::cmp::max(ret, con);
        } else {
            used.clear();
            con = 1;
        }
        used.insert(c);
    }
    ret
}
fn main() {
    dbg!(
        length_of_longest_substring("abcabcbb".to_string()),
        length_of_longest_substring("bbbbb".to_string()),
        length_of_longest_substring("pwwkew".to_string()),
        length_of_longest_substring("".to_string()),
    );
}
