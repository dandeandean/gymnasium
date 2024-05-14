// based off of:  https://leetcode.com/problems/generate-parentheses/solutions/4570907/performance-optimized-backtrack-solution-in-rust-0ms-2-28-mb-64-memory
fn generate(n: i32, s: &mut String, out: &mut Vec<String>) {
    let num_open = s.chars().filter(|c| *c == '(').count() as i32;
    let num_clos = s.chars().filter(|c| *c == ')').count() as i32;
    if n == num_clos && num_open == n {
        out.push(s.clone());
        return;
    };

    if num_open < n {
        s.push('(');
        generate(n, s, out);
        s.pop();
    }
    if num_clos < num_open {
        s.push(')');
        generate(n, s, out);
        s.pop();
    }
}
pub fn generate_parenthesis(n: i32) -> Vec<String> {
    let mut out: Vec<String> = Vec::new();
    generate(n, &mut "".to_string(), &mut out);
    out
}
fn main() {
    let res = generate_parenthesis(3);
    dbg!(res);
}
