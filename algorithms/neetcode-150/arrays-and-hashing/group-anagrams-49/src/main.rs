use std::collections::HashMap;

fn sort_string(w: &String) -> String {
    let mut alphabet: Vec<usize> = vec![0; 26];
    for c in w.chars() {
        let ci = c as usize - 'a' as usize;
        alphabet[ci] += 1;
    }
    let mut out = String::from("");
    for count in alphabet {
        let letter = (count as u8 + 97u8) as char;
        out.push(letter);
    }
    return out;
}

pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
    let mut m: HashMap<String, Vec<String>> = HashMap::new();
    for word in strs {
        let key = sort_string(&word);
        let e = m.entry(key).or_insert(vec![]);
        e.push(word);
    }
    m.into_values().collect()
}

fn main() {}
