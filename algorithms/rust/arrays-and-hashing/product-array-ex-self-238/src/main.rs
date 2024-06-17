use std::collections::HashMap;

pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
    let mut prefix_at_i: HashMap<usize, i32> = HashMap::new();
    let mut postfix_at_i: HashMap<usize, i32> = HashMap::new();
    // calc prefixes & postfixes
    let mut preduct: i32 = 1;
    let mut postduct: i32 = 1;
    let cap = nums.len() - 1;
    for i in 0..cap {
        preduct *= nums[i];
        prefix_at_i.insert(i + 1, preduct);

        postfix_at_i.insert(cap - i, postduct);
        postduct *= nums[cap - i];
    }
    prefix_at_i.insert(0, 1);
    postfix_at_i.insert(0, postduct);
    let mut out: Vec<i32> = vec![];
    for i in 0..nums.len() {
        out.push(prefix_at_i[&i] * postfix_at_i[&i]);
    }
    out
}

fn main() {
    let nums = vec![1, 2, 3, 4];
    let out = product_except_self(nums);
    dbg!(out);
}
