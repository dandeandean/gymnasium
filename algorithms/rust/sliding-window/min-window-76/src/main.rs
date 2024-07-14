use std::usize;

pub fn min_window(s: String, t: String) -> String {
    if t == "".to_string() {
        return t;
    }
    let mut t_hash: std::collections::HashMap<char, i32> =
        std::collections::HashMap::with_capacity(t.len());
    let mut window: std::collections::HashMap<char, i32> =
        std::collections::HashMap::with_capacity(t_hash.len());
    for c in t.chars() {
        *t_hash.entry(c).or_insert(0) += 1;
        window.entry(c).or_insert(0);
    }
    // we can make this one balance
    let (mut have, need): (i32, i32) = (0, t_hash.len() as i32);
    let (mut l, mut r): (usize, usize) = (0, 0);
    let s_chars = s.as_bytes();
    let mut res: (usize, usize) = (0, 0);
    let mut res_len = usize::MAX;
    while r < s.len() {
        let c: char = s_chars[r] as char;
        if t_hash.contains_key(&c) {
            // we already know it's in there
            *window.get_mut(&c).unwrap() += 1;
            if window.get(&c) == Some(&t_hash[&c]) {
                have += 1;
            }
            while have == need {
                if (r - l) < res_len {
                    res = (l, r);
                    res_len = r + 1 - l;
                }
                let cl = s_chars[l] as char;
                // we are not sure what the left is pointing at so we need this
                if let Some(v) = window.get_mut(&cl) {
                    *v -= 1;
                    if *v < t_hash[&cl] {
                        have -= 1;
                    }
                }
                l += 1;
            }
        }
        r += 1;
    }
    match res_len {
        usize::MAX => "".to_string(),
        _ => s[res.0..res.1 + 1].to_string(),
    }
}
fn main() {
    dbg!(
        min_window("ADOBECODEBANC".to_string(), "ABC".to_string()),
        min_window("A".to_string(), "A".to_string()),
        min_window("A".to_string(), "AA".to_string()),
    );
}
