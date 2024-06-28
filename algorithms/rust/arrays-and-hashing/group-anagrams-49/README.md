---
tags:
  - neetcode
---

See it on the [website](https://leetcode.com/problems/group-anagrams/description/). 

Given an array of strings `strs`, group **the anagrams** together. You can return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Example 1:

```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

## Example 2:

```
Input: strs = [""]
Output: [[""]]
```

## Example 3:

```
Input: strs = ["a"]
Output: [["a"]]
```
## Constraints:

- `1 <= strs.length <= 104`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

# Solution:
The key here is to use the sorted string as a key to group the anagrams together. After that we can add them all to a [[Hash Map]] that takes those sorted strings and adds to the vector pointed to by that key. Collect and submit that hash map for the final answer. 
```rust 
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

```