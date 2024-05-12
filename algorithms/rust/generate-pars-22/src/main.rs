pub fn generate_parenthesis(n: i32) -> Vec<String> {
    let out: Vec<String> = vec![];
    let stack: Vec<String> = vec![];
    fn generate(left: i32, right: i32, n: i32, &mut stack: Vec<String>, &mut out: Vec<String>) {
        if left == right && right == n {
            out.push(stack.join(""));
            return;
        }
        if left < n {
            stack.push("(".to_string());
            generate(left + 1, right, n, stack.clone(), out.clone());
            stack.pop();
        }
        if right < n {
            stack.push(")".to_string());
            generate(left + 1, right, n, stack.clone(), out.clone());
            stack.pop();
        }
    };
    generate(0, 0, n, stack, out);
    out
}
fn main() {}
