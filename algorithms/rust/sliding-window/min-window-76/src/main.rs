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
    let (resi, resj): (usize, usize) = (0, 0);
    let mut res: (usize, usize);
    let mut res_len = usize::max_value();
    while r < s.len() {
        let c: char = s_chars[r] as char;
        if t_hash.contains_key(&c) {
            // we already know it's in there
            if let Some(v) = window.get_mut(&c) {
                *v += 1;
                if *v == t_hash[&c] {
                    have += 1;
                }
            }
            while have == need {
                dbg!(r, l);
                if (r - l) < res_len {
                    res = (r, l);
                    res_len = res.0 + 1 - res.1;
                }
                let cl = s_chars[l] as char;
                if let Some(v) = window.get_mut(&cl) {
                    *v -= 1;
                    if *v < t_hash[&cl] {
                        have += 1;
                    }
                }
                l += 1;
            }
        }
        r += 1;
    }
    s[resi..resj].to_string()
}
fn main() {
    dbg!(
        min_window("ADOBECODEBANC".to_string(), "ABC".to_string()),
        min_window("A".to_string(), "A".to_string()),
        min_window("AA".to_string(), "A".to_string()),
    );
}
