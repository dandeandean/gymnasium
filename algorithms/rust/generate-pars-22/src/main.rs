fn generate(left: i32, right: i32, n: i32, mut stack: Vec<String>) -> Vec<String> {
    let mut out: Vec<String> = vec![];
    if left == right && right == n {
        out.push(stack.join(""));
        return out;
    }
    if left < n {
        let mut sl = stack.clone();
        sl.push("(".to_string());
        generate(left + 1, right, n, sl);
        stack.pop();
    }
    if right < n {
        let mut sr = stack.clone();
        sr.push("(".to_string());
        generate(left, right + 1, n, sr);
        stack.pop();
    }
    out
}
pub fn generate_parenthesis(n: i32) -> Vec<String> {
    let out = generate(0, 0, n, vec![]);
    out
}
fn main() {
    let res = generate_parenthesis(3);
    dbg!(res);
}
